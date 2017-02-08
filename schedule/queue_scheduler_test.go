package schedule

import (
	"testing"
	"fmt"
	"crawler/downloader/request"
)

func Test_queue_scheduler(t *testing.T) {
	q := NewQueueSchedule()
	q.Push(&request.Request{

	})
	for index, value := range q.Close() {
		fmt.Printf("%d,%v \n", index, value)
	}
	for {
		e := q.Pull()
		if e == nil {
			break
		} else {
			fmt.Printf("%v \n", e)
		}
	}

}