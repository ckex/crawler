package init

import (
 	logger "github.com/alecthomas/log4go"
)

func init()  {
	logger.LoadConfiguration("./conf/log4go.xml")
}
