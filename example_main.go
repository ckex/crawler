package main

import (
	_ "crawler/init"
	logger "github.com/alecthomas/log4go"
	"os"
	"os/signal"
	"crawler/downloader/request"
	"crawler/spider"
	"crawler/schedule"
	"fmt"
	"crawler/process"
)

func main() {
	defer logger.Close()
	c := make(chan os.Signal, 1)
	logger.Info("Start crawler ...")
	startSpider()
	signal.Notify(c)
	s := <-c
	logger.Info("Google Bye, ", s)
}

func startSpider() {
	spider := spider.New(&TestPageProcess{

	})
	spider.Push(&request.Request{
		Url:"http://www.baidu.com",
	})
	spider.Start()
}

type TestPageProcess struct {

}

func (this *TestPageProcess) OnProcess(schedule schedule.Schedule, request *request.Request, response *process.Result) () {
	fmt.Printf("%T %+v \n", schedule, schedule)
	fmt.Printf("%T %+v \n", request, request)
	fmt.Printf("%T %+v \n", response, response)
	fmt.Printf("%s \n", response.Status)
	fmt.Printf("%s \n", response.BodyString())
}