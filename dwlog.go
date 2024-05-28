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
	if dwlog == nil {
		New()
	}
	dwlog.Logger.SetOutput(w)
}

func GetCallerStack() int {
	return dwlog.callerstack
}

func SetCallerStack(i int) {
	dwlog.callerstack = i
}

func Info(msg string, args ...any) {
	if dwlog == nil {
		New()
	}
	prevCallerStack := GetCallerStack()
	SetCallerStack(3)
	dwlog.Info(msg, args...)
	SetCallerStack(prevCallerStack)
}

func Log(msg string, args ...any) {
	if dwlog == nil {
		New()
	}
	prevCallerStack := GetCallerStack()
	SetCallerStack(3)
	dwlog.Log(msg, args...)
	SetCallerStack(prevCallerStack)
}

func Warning(msg string, args ...any) {
	if dwlog == nil {
		New()
	}
	prevCallerStack := GetCallerStack()
	SetCallerStack(3)
	dwlog.Warning(msg, args...)
	SetCallerStack(prevCallerStack)
}

func Error(msg string, args ...any) {
	if dwlog == nil {
		New()
	}
	prevCallerStack := GetCallerStack()
	SetCallerStack(3)
	dwlog.Error(msg, args...)
	SetCallerStack(prevCallerStack)
}
