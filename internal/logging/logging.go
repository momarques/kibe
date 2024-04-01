package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()
var LogFile = "debug.log"

func init() {
	file, err := os.OpenFile(LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Log.Out = file
	if os.Getenv("DEBUG") == "1" {
		Log.Level = logrus.DebugLevel
	}
}
