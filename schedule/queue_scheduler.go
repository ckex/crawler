package schedule

import (
	"crawler/downloader/request"
	"sync"
)

type QueueSchedule struct {
	queue chan *request.Request
	mutex sync.Mutex
}

func NewQueueSchedule() Schedule {
	schedule := new(QueueSchedule)
	schedule.queue = make(chan *request.Request, 10)
	return schedule
}

func (this *QueueSchedule) Pull() *request.Request {
	return <-this.queue
}

func (this *QueueSchedule)  Push(request *request.Request) error {
	this.queue <- request
	return nil
}

func (this *QueueSchedule) Close() []*request.Request {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	close(this.queue)
	result := make([]*request.Request, 0)
	for {
		if req, ok := <-this.queue; ok {
			result = append(result, req)
			continue
		}
		break
	}
	return result
}

