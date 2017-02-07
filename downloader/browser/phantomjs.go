package browser

import (
	"path/filepath"
	"os"
	"log"
	"crawler/utils"
	"net/http"
	"crawler/downloader/request"
	"mime"
	"strings"
	"time"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"crawler/conf"
	"crawler/downloader"
	"fmt"
)

type (
	Phantom struct {
		PhantomjsFile string            //Phantomjs完整文件名
		TempJsDir     string            //临时js存放目录
		jsFileMap     map[string]string //已存在的js文件
	}
	Response struct {
		Cookies []string
		Body    string
	}
)

var DefaultPhantom = NewPhantom(conf.PHANTOMJS, conf.CACHE_DIR)

func NewPhantom(phantomjsFile, tempJsDir string) downloader.Downloader {

	if exist := utils.FileExist(phantomjsFile); !exist {
		fmt.Printf("%s is not exists\n", phantomjsFile)
		return nil
	}

	phantom := &Phantom{
		PhantomjsFile:phantomjsFile,
		TempJsDir:tempJsDir,
		jsFileMap:make(map[string]string),
	}
	if !filepath.IsAbs(phantom.PhantomjsFile) {
		phantom.PhantomjsFile, _ = filepath.Abs(phantom.PhantomjsFile)
	}
	if !filepath.IsAbs(phantom.TempJsDir) {
		phantom.TempJsDir, _ = filepath.Abs(phantom.TempJsDir)
	}
	// 创建/打开目录
	err := os.MkdirAll(phantom.TempJsDir, 0777)
	if err != nil {
		log.Printf("[E] Surfer: %v\n", err)
		return phantom
	}
	phantom.createJsFile("js", js)
	return phantom
}


// 实现 Downloader 下载器接口
func (client *Phantom) Download(reqeust *request.Request) (response *http.Response, err error) {
	var encoding = "utf-8"
	if _, params, err := mime.ParseMediaType(reqeust.GetHeader().Get("Content-Type")); err == nil {
		if cs, ok := params["charset"]; ok {
			encoding = strings.ToLower(strings.TrimSpace(cs))
		}
	}

	reqeust.GetHeader().Del("Content-Type")

	param, err := NewParam(reqeust)
	if err != nil {
		return nil, err
	}
	response = param.writeback(response)

	var args = []string{
		client.jsFileMap["js"],
		reqeust.GetUrl(),
		param.header.Get("Cookie"),
		encoding,
		param.header.Get("User-Agent"),
		reqeust.GetPostParams(),
		strings.ToLower(param.method),
	}

	for i := 0; i < param.tryTimes; i++ {
		fmt.Println(client.PhantomjsFile, args)
		cmd := exec.Command(client.PhantomjsFile, args...)
		if response.Body, err = cmd.StdoutPipe(); err != nil {
			fmt.Errorf(err.Error())
			time.Sleep(param.retryPause)
			continue
		}
		if cmd.Start() != nil || response.Body == nil {
			fmt.Errorf(err.Error())
			time.Sleep(param.retryPause)
			continue
		}
		var b []byte
		b, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Errorf(err.Error())
			time.Sleep(param.retryPause)
			continue
		}
		resultResp := Response{}
		err = json.Unmarshal(b, &resultResp)
		if err != nil {
			fmt.Errorf("Unmarshal error. ", err.Error())
			time.Sleep(param.retryPause)
			continue
		}
		response.Header = param.header
		for _, cookie := range resultResp.Cookies {
			response.Header.Add("Set-Cookie", cookie)
		}
		response.Body = ioutil.NopCloser(strings.NewReader(resultResp.Body))
		break
	}

	if err == nil {
		response.StatusCode = http.StatusOK
		response.Status = http.StatusText(http.StatusOK)
	} else {
		response.StatusCode = http.StatusBadGateway
		response.Status = http.StatusText(http.StatusBadGateway)
	}
	return
}

//销毁js临时文件
func (self *Phantom) DestroyJsFiles() {
	p, _ := filepath.Split(self.TempJsDir)
	if p == "" {
		return
	}
	for _, filename := range self.jsFileMap {
		os.Remove(filename)
	}
	if len(utils.WalkDir(p)) == 1 {
		os.Remove(p)
	}
}

func (self *Phantom) createJsFile(fileName, jsCode string) {
	fullFileName := filepath.Join(self.TempJsDir, fileName)
	// 创建并写入文件
	f, _ := os.Create(fullFileName)
	f.Write([]byte(jsCode))
	f.Close()
	self.jsFileMap[fileName] = fullFileName
}

/*
* system.args[0] == post.js
* system.args[1] == url
* system.args[2] == cookie
* system.args[3] == pageEncode
* system.args[4] == userAgent
* system.args[5] == postdata
* system.args[6] == method
 */
const js string = `
var system = require('system');
var page = require('webpage').create();
var url = system.args[1];
var cookie = system.args[2];
var pageEncode = system.args[3];
var userAgent = system.args[4];
var postdata = system.args[5];
var method = system.args[6];
page.onResourceRequested = function(requestData, request) {
    request.setHeader('Cookie', cookie)
};
phantom.outputEncoding = pageEncode;
page.settings.userAgent = userAgent;
page.open(url, method, postdata, function(status) {
   if (status !== 'success') {
        console.log('Unable to access network');
    } else {
        var cookies = new Array();
        for(var i in page.cookies) {
        	var cookie = page.cookies[i];
        	var c = cookie["name"] + "=" + cookie["value"];
        	for (var obj in cookie){
        		if(obj == 'name' || obj == 'value'){
        			continue;
        		}
				c +=  "; " +　obj + "=" +  cookie[obj];
    		}
			cookies[i] = c;
		}
        var resp = {
            "Cookies": cookies,
            "Body": page.content
        };
        console.log(JSON.stringify(resp));
    }
    phantom.exit();
});
`
