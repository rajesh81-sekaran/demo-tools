package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"demo-tools/sumup"

	Log "github.com/sirupsen/logrus"
)

var (
	//LogFile
	LogFile = flag.String("logFile", "", "log file location")
	//DebugLevel
	DebugLevel = flag.String("d", "info", "debug level, any one of panic, fatal, error, warn, warning, info, debug, trace")
)

func main() {
	flag.Parse()
	defLevel, _ := Log.ParseLevel("info")
	Log.SetLevel(defLevel)
	Log.SetReportCaller(true)
	if *LogFile != "" {
		logFile, err := os.OpenFile(*LogFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
		if err != nil {
			Log.Errorf("Error [%s] opening log file [%s]", err, *LogFile)
		} else {
			Log.SetOutput(logFile)
		}
	}
	Log.Infof("In main(testing commit)")
	t := "SpewTest"
	fmt.Printf("t:")
	spew.Dump(t)
	fmt.Println(sumup.Sum(40, 2))
}
