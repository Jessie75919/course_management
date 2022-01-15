package logger

import (
	"course_management/config"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

const logFile = "var/log/debug.log"

func Setup() {
	errorHandler := os.Stderr
	infoHandler := os.Stdout
	debugHandler := ioutil.Discard

	if config.EnvConfig.DEBUG {
		debugHandler = os.Stdout
	}

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	Debug = createLog(debugHandler, "Debug", f)
	Info = createLog(infoHandler, "Info", f)
	Error = createLog(errorHandler, "Error", f)
}

func createLog(handler io.Writer, prefix string, f io.Writer) *log.Logger {
	logger := log.New(handler, prefix+": ", log.Ldate|log.Lmicroseconds|log.LUTC)
	logger.SetOutput(f)
	return logger
}
