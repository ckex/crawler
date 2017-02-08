package spider

import (
	"net/http"
	"crawler/downloader/request"
)

type Listener interface {
	OnError(*Spider, *request.Request, error)
	OnSuccess(*Spider, *request.Request, *http.Response)
}
