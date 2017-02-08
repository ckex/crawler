package schedule

import "crawler/downloader/request"

type Schedule interface {
	Pull() *request.Request

	Push(req *request.Request) error

	Close() []*request.Request
}
