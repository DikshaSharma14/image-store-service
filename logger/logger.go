package logger

import (
	"fmt"
	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func init() {
	if Logger == nil {
		logFile := "seelog.xml"
		logger, err := seelog.LoggerFromConfigAsFile(logFile)
		if err != nil {
			fmt.Println(err.Error(), err)
		}

		Logger = logger
		defer Logger.Flush()
	}
}
