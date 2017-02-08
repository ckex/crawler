package spider

import (
	"crawler/process"
	"sync"
	"crawler/downloader"
	"crawler/downloader/browser"
	"crawler/schedule"
	"time"
)

type Spider struct {
	process.PageProcess
	schedule.Schedule
	downloader.Downloader
	Goroutines, status int
	once               sync.Once
	Listener           []Listener
	lock               *sync.Mutex
}

type initFun func(*Spider) error

const (
	wait = iota
	run
	stop
)

var (
	inits = make([]initFun, 0)
)

func New(pageProcess process.PageProcess) *Spider {
	return &Spider{
		PageProcess:pageProcess,
		Downloader:browser.NewHttpDownloader(),
		Schedule:schedule.NewQueueSchedule(),
		Listener:make([]Listener, 0),
	}
}

func (spider *Spider) Start() {
	spider.once.Do(spider.run)
}

func (spider *Spider) AddListener(listener Listener) {
	spider.Listener = append(spider.Listener, listener)
}

func (spider *Spider) SetSchedule(sche *schedule.QueueSchedule) {
	spider.lock.Lock()
	defer spider.lock.Unlock()
	for _, value := range spider.Schedule.Close() {
		sche.Push(value)
	}
}

func (spider *Spider) run() {
	spider.initComponent()
	for i := 0; i < spider.Goroutines; i++ {
		go func(this *Spider) {
			this.execute()
		}(spider)
	}
}
func (this *Spider) Stop() {
	this.status = stop
}

func (this *Spider) execute() {
	for {
		if this.status == stop {
			break
		}
		if request := this.Schedule.Pull(); request != nil {
			resp, err := this.Download(request)
			if err != nil {
				// Listener
				for _, fn := range this.Listener {
					fn.OnError(this, request, err)
				}
			} else {
				this.OnProcess(this.Schedule, request, process.NewResult(resp))
				// Listener
				for _, fn := range this.Listener {
					fn.OnSuccess(this, request, resp)
				}
			}
			continue
		}
		time.Sleep(time.Duration(time.Nanosecond * 1000 * 1000 * 500))
	}
}

func addInitFunc(fn initFun) {
	inits = append(inits, fn)
}

func initDownloader(this *Spider) error {
	if this.Downloader == nil {
		this.Downloader = browser.NewHttpDownloader()
	}
	return nil
}

func initGoroutines(this *Spider) error {
	if this.Goroutines < 1 {
		this.Goroutines = 1
	}
	return nil
}

func (spider *Spider ) initComponent() {
	addInitFunc(initDownloader)
	addInitFunc(initGoroutines)
	for _, val := range inits {
		if err := val(spider); err != nil {
			panic(err)
		}
	}
	spider.status = run

}

