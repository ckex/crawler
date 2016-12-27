package main

import (
	_ "robot/init"
	logger "github.com/alecthomas/log4go"
	"os"
	"os/signal"
)

func main()  {
	defer  logger.Close()
	c := make(chan os.Signal,1)
	logger.Info("start robot ...")
	signal.Notify(c)
	s := <-c
	logger.Info("Google Bye, ",s)
}
