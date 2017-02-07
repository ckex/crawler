package  browser

import (
	"crawler/downloader/request"
	"io"
	"net/http"
	"time"
	"strings"
	"math/rand"
	"net"
	"crypto/tls"
	"net/http/cookiejar"
	"compress/gzip"
	"compress/flate"
	"compress/zlib"
	"crawler/downloader"
)

type HttpDownloader struct {
	cookieJar *cookiejar.Jar
}

func NewHttpDownloader() downloader.Downloader {
	clinet := new(HttpDownloader)
	cjar, _ := cookiejar.New(nil)
	clinet.cookieJar = cjar
	return clinet
}

func (client *HttpDownloader) Download(reqeust *request.Request) (resp *http.Response, err error) {
	param, err := NewParam(reqeust)
	if err != nil {
		return nil, err
	}
	param.client = client.buildClient(param)
	resp, err = httpRequest(param)
	if err == nil {
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			var gzipReader *gzip.Reader
			gzipReader, err = gzip.NewReader(resp.Body)
			if err == nil {
				resp.Body = gzipReader
			}

		case "deflate":
			resp.Body = flate.NewReader(resp.Body)

		case "zlib":
			var readCloser io.ReadCloser
			readCloser, err = zlib.NewReader(resp.Body)
			if err == nil {
				resp.Body = readCloser
			}
		}
	}
	resp = param.writeback(resp)
	return
}

// send uses the given *http.Request to make an HTTP request.
func httpRequest(param *Param) (resp *http.Response, err error) {
	req, err := http.NewRequest(param.method, param.url.String(), param.body)
	if err != nil {
		return nil, err
	}

	req.Header = param.header

	if param.tryTimes <= 0 {
		for {
			resp, err = param.client.Do(req)
			if err != nil {
				if !param.enableCookie {
					l := len(UserAgents["common"])
					r := rand.New(rand.NewSource(time.Now().UnixNano()))
					req.Header.Set("User-Agent", UserAgents["common"][r.Intn(l)])
				}
				time.Sleep(param.retryPause)
				continue
			}
			break
		}
	} else {
		for i := 0; i < param.tryTimes; i++ {
			resp, err = param.client.Do(req)
			if err != nil {
				if !param.enableCookie {
					l := len(UserAgents["common"])
					r := rand.New(rand.NewSource(time.Now().UnixNano()))
					req.Header.Set("User-Agent", UserAgents["common"][r.Intn(l)])
				}
				time.Sleep(param.retryPause)
				continue
			}
			break
		}
	}

	return resp, err
}

// buildClient creates, configures, and returns a *http.Client type.
func (this *HttpDownloader) buildClient(param *Param) *http.Client {
	client := &http.Client{
		CheckRedirect: param.checkRedirect,
	}

	if param.enableCookie {
		client.Jar = this.cookieJar
	}

	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, param.dialTimeout)
			if err != nil {
				return nil, err
			}
			if param.connTimeout > 0 {
				c.SetDeadline(time.Now().Add(param.connTimeout))
			}
			return c, nil
		},
	}

	if param.proxy != nil {
		transport.Proxy = http.ProxyURL(param.proxy)
	}

	if strings.ToLower(param.url.Scheme) == "https" {
		transport.TLSClientConfig = &tls.Config{RootCAs: nil, InsecureSkipVerify: true}
		transport.DisableCompression = true
	}
	client.Transport = transport
	return client
}

