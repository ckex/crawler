package browser

import (
	"net/url"
	"io"
	"net/http"
	"time"
	"fmt"
	"strings"
	"robot/downloader/request"
	"math/rand"
)

func NewParam(request *request.Request) (param *Param, err error) {
	param = new(Param)
	// URL
	param.url, err = UrlEncode(request.GetUrl())
	if err != nil {
		return nil, err
	}
	// Proxy
	if request.GetProxy() != "" {
		if param.proxy, err = url.Parse(request.GetProxy()); err != nil {
			return nil, err
		}
	}
	// Header
	param.header = request.GetHeader()
	if param.header == nil {
		param.header = make(http.Header)
	}
	// Methed
	switch method := strings.ToUpper(request.GetMethod()); method {
	case "GET", "HEAD":
		param.method = method
	case "POST":
		param.method = method
		param.body = strings.NewReader(request.GetPostParams())
	default:
		param.method = "GET"
	}
	//Cookie
	param.enableCookie = request.GetEnableCookie()
	//UserAgent
	if len(param.header.Get("User-Agent")) == 0 {
		if param.enableCookie {
			param.header.Add("User-Agent", UserAgents["common"][0])
		} else {
			l := len(UserAgents["common"])
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			param.header.Add("User-Agent", UserAgents["common"][r.Intn(l)])
		}
	}
	//Dial Timeout
	param.dialTimeout = request.GetDialTimeout()
	if param.dialTimeout < 0 {
		param.dialTimeout = 0
	}
	//Conn Timeout
	param.connTimeout = request.GetConnTimeout()
	param.tryTimes = request.GetTryTimes()
	param.retryPause = request.GetRetryPause()
	param.redirectTimes = request.GetRedirectTimes()
	return
}

type Param struct {
	method        string
	url           *url.URL
	proxy         *url.URL
	body          io.Reader
	header        http.Header
	enableCookie  bool
	dialTimeout   time.Duration
	connTimeout   time.Duration
	tryTimes      int
	retryPause    time.Duration
	redirectTimes int
	client        *http.Client
}

// 回写Request内容
func (self *Param) writeback(resp *http.Response) *http.Response {
	if resp == nil {
		resp = new(http.Response)
		resp.Request = new(http.Request)
	} else if resp.Request == nil {
		resp.Request = new(http.Request)
	}

	resp.Request.Method = self.method
	resp.Request.Header = self.header
	resp.Request.Host = self.url.Host

	return resp
}

// checkRedirect is used as the value to http.Client.CheckRedirect
// when redirectTimes equal 0, redirect times is ∞
// when redirectTimes less than 0, not allow redirects
func (self *Param) checkRedirect(req *http.Request, via []*http.Request) error {
	if self.redirectTimes == 0 {
		return nil
	}
	if len(via) >= self.redirectTimes {
		if self.redirectTimes < 0 {
			return fmt.Errorf("not allow redirects.")
		}
		return fmt.Errorf("stopped after %v redirects.", self.redirectTimes)
	}
	return nil
}


