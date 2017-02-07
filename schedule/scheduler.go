package schedule

import "crawler/downloader/request"

type Schedule interface {
	Pull() *request.Request
	Push(*request.Request)
}
