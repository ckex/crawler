package downloader

import (
	"crawler/downloader/request"
	"net/http"
)

type Downloader interface {
	Download(*request.Request) (resp *http.Response, err error)
}