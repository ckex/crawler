package downloader

import (
	"robot/downloader/request"
	"net/http"
)

type Downloader interface {
	Download(*request.Request) (resp *http.Response, err error)
}