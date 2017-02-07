package spider

import (
	"crawler/process"
	"sync"
	"crawler/downloader"
	"crawler/downloader/browser"
)

type Spider struct {
	process.PageProcess
	once       sync.Once
	downloader.Downloader
	Goroutines int
}

type initFun func(*Spider) error

var (
	inits = make([]initFun, 0)
)

func New(pageProcess *process.PageProcess) *Spider {
	return &Spider{
		PageProcess:pageProcess,
		Downloader:browser.NewHttpDownloader(),
	}
}

func (spider *Spider) Start() {
	spider.once.Do(spider.run)
}

func (spider *Spider) run() {
	spider.initComponent()
}

func addInitFunc(fn initFun) {
	inits = append(inits, fn)
}

func initDownloader(this *Spider) error {
	if this.Downloader == nil {
		this.Downloader = browser.NewHttpDownloader()
	}
}

func initGoroutines(this *Spider) error {
	if this.Goroutines < 1 {
		this.Goroutines = 1
	}
}

func (spider *Spider ) initComponent() {
	addInitFunc(initDownloader)
	addInitFunc(initGoroutines)
	for _, val := range inits {
		if err := val(spider); err != nil {
			panic(err)
		}
	}

}

