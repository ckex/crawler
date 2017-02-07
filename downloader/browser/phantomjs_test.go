package browser

import (
	"fmt"
	"testing"
	"crawler/downloader/request"
	"bytes"
	"crawler/conf"
	"net/http"
)

//  go test -run="Test_Phantomjs_Download"
func Test_Phantomjs_Download(t *testing.T) {
	var cleint = NewPhantom("/Users/ckex/work/golang/workspace/src/crawler/crawler_bin/phantomjs", conf.CACHE_DIR)
	//var cleint = DefaultPhantom
	fmt.Printf("%v\n", cleint)
	response, err := cleint.Download(&request.Request{
		Url:"http://www.tianyancha.com/search?key=%E6%B7%B1%E5%9C%B3%E5%A4%A9%E9%81%93%E8%AE%A1%E7%84%B6%E9%87%91%E8%9E%8D%E6%9C%8D%E5%8A%A1%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8&checkFrom=searchBox",
		// Url:"http://www.baidu.com",
	})
	if err != nil {
		t.Error(err)
		return
	}
	print(response)
}

func Test_Phantomjs_China_Unicom(t *testing.T) {

	url := "https://uac.10010.com/portal/Service/MallLogin?callback=jQuery17206799110582967826_1484807327918&req_time=1484807337290&redirectURL=http%3A%2F%2Fwww.10010.com&userName=ckex868%40vip.qq.com&password=ckex13534117937&pwdType=01&productType=05&redirectType=01&rememberMe=1&_=1484807337292"
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.StatusCode)
	var cleint = NewPhantom("/Users/ckex/work/golang/workspace/src/crawler/crawler_bin/phantomjs", conf.CACHE_DIR)
	response, err := cleint.Download(&request.Request{
		Url:"http://iservice.10010.com/e4/index_server.html",
		Site:request.Site{
			//Header:resp.Header,
		},
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	print(response)
}

func print(response *http.Response) {
	buf := new(bytes.Buffer)
	cookies := response.Cookies()
	buf.ReadFrom(response.Body)
	fmt.Println("Body------>", buf.String())
	for index, value := range cookies {
		fmt.Printf("Cookie -- %d -- %v\n", index, value)
	}
}