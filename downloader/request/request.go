package request

import (
	"net/http"
	"sync"
	"time"
	"strings"
)

const (
	DefaultMethod = "GET"           // 默认请求方法
	DefaultDialTimeout = 2 * time.Minute // 默认请求服务器超时
	DefaultConnTimeout = 2 * time.Minute // 默认下载超时
	DefaultTryTimes = 3               // 默认最大下载次数
	DefaultRetryPause = 5 * time.Second // 默认重新下载前停顿时长
)

type Site struct {
	UserAgnet, Proxy, Domain             string
	Header                               http.Header
	EnableCookie                         bool
	DialTimeout, ConnTimeout, RetryPause time.Duration
	TryTimes, RedirectTimes              int
}

type Request struct {
	Site
	Url        string //目标URL，必须设置
	Method     string //GET POST
	PostParams string //post params
	once       sync.Once
}

// prepare
func (this *Request) prepare() {
	if this.Method == "" {
		this.Method = DefaultMethod
	}
	this.Method = strings.ToUpper(this.Method)

	if this.Header == nil {
		this.Header = make(http.Header)
	}

	if this.DialTimeout < 0 {
		this.DialTimeout = 0
	} else if this.DialTimeout == 0 {
		this.DialTimeout = DefaultDialTimeout
	}

	if this.ConnTimeout < 0 {
		this.ConnTimeout = 0
	} else if this.ConnTimeout == 0 {
		this.ConnTimeout = DefaultConnTimeout
	}

	if this.TryTimes == 0 {
		this.TryTimes = DefaultTryTimes
	}

	if this.RetryPause <= 0 {
		this.RetryPause = DefaultRetryPause
	}

}

func (this *Request)  GetUrl() string {
	this.once.Do(this.prepare)
	return this.Url
}
func (this *Request)  GetMethod() string {
	this.once.Do(this.prepare)
	return this.Method
}
func (this *Request)  GetPostParams() string {
	this.once.Do(this.prepare)
	return this.PostParams
}
func (this *Request)  GetUserAgent() string {
	this.once.Do(this.prepare)
	return this.UserAgnet
}
func (this *Request)  GetProxy() string {
	this.once.Do(this.prepare)
	return this.Proxy
}
func (this *Request)  GetDomain() string {
	this.once.Do(this.prepare)
	return this.Domain
}
func (this *Request)  GetHeader() http.Header {
	this.once.Do(this.prepare)
	return this.Header
}
func (this *Request)  GetEnableCookie() bool {
	this.once.Do(this.prepare)
	return this.EnableCookie
}
func (this *Request)  GetConnTimeout() time.Duration {
	this.once.Do(this.prepare)
	return this.ConnTimeout
}
func (this *Request)  GetDialTimeout() time.Duration {
	this.once.Do(this.prepare)
	return this.DialTimeout
}
func (this *Request)  GetTryTimes() int {
	this.once.Do(this.prepare)
	return this.TryTimes
}
func (this *Request)  GetRetryPause() time.Duration {
	this.once.Do(this.prepare)
	return this.RetryPause
}
func (this *Request)  GetRedirectTimes() int {
	this.once.Do(this.prepare)
	return this.RedirectTimes
}