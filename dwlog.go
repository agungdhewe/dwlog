package dwlog

import (
	"io"
	"log"
)

var dwlog *DwLogger

func New() *DwLogger {
	dwlog = &DwLogger{
		Logger: log.New(log.Writer(), "", log.Lmicroseconds|log.Lshortfile),
	}
	return dwlog
}

func SetOutput(w io.Writer) {
	dwlog.Logger.SetOutput(w)
}

func Info(msg string, args ...any) {
	dwlog.Info(msg, args...)
}

func Log(msg string, args ...any) {
	dwlog.Log(msg, args...)
}

func Warning(msg string, args ...any) {
	dwlog.Warning(msg, args...)
}

func Error(msg string, args ...any) {
	dwlog.Error(msg, args...)
}
