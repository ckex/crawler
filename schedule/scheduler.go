package schedule

import "robot/downloader/request"

type Schedule interface {
	Pull() *request.Request
	Push(*request.Request)
}
