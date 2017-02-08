package process

import (
	"crawler/downloader/request"
	"net/http"
	"crawler/schedule"
	"bytes"
)

type PageProcess interface {
	OnProcess(schedule schedule.Schedule, request *request.Request, response *Result)
}

type Result struct {
	*http.Response
}

func NewResult(response *http.Response) *Result {
	return &Result{response}
}

func (this *Result) BodyString() string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(this.Body)
	return buf.String()
}