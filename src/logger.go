package dwlog

import (
	"fmt"
	"io"
	"log"
)

var logger *log.Logger

func New() *DwLogger {
	logger = log.New(log.Writer(), "", log.Lmicroseconds|log.Lshortfile)
	return &DwLogger{
		Logger: logger,
	}

}

func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

func (dwl *DwLogger) Info(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)

	label := setColor(ColorBgWhite, setColor(ColorFgBlack, " INFO "))
	text := setColor(ColorNormalText, msg)

	dwl.Logger.SetFlags(0)
	dwl.Logger.SetPrefix(label)
	dwl.Logger.Printf(" %s", text)

	resetPrefix()
}

func (dwl *DwLogger) Log(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)

	label := setColor(ColorBgCyan, setColor(ColorFgBlack, " LOG  "))
	text := setColor(ColorNormalText, msg)

	callerfile, lineno := getCaller()
	ref := setColor(ColorFgCyan, fmt.Sprintf("%s:%d", callerfile, lineno))

	dwl.Logger.SetFlags(0)
	dwl.Logger.SetPrefix(label)

	if callerfile != "" {
		dwl.Logger.Printf(" %s %s", text, ref)
	} else {
		dwl.Logger.Printf(" %s", text)
	}

	resetPrefix()
}

func (dwl *DwLogger) Warning(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)

	label := setColor(ColorBgYellow, setColor(ColorFgBlack, " WARN "))
	text := setColor(ColorFgYellow, msg)

	callerfile, lineno := getCaller()
	ref := setColor(ColorNormalText, fmt.Sprintf("%s:%d", callerfile, lineno))

	dwl.Logger.SetFlags(0)
	dwl.Logger.SetPrefix(label)

	if callerfile != "" {
		dwl.Logger.Printf(" %s %s", text, ref)
	} else {
		dwl.Logger.Printf(" %s", text)
	}

	resetPrefix()
}

func (dwl *DwLogger) Error(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)

	label := setColor(ColorBgRed, setColor(ColorFgYellow, " ERR  "))
	text := setColor(ColorFgRed, msg)

	callerfile, lineno := getCaller()
	ref := setColor(ColorNormalText, fmt.Sprintf("%s:%d", callerfile, lineno))

	dwl.Logger.SetFlags(0)
	dwl.Logger.SetPrefix(label)

	if callerfile != "" {
		dwl.Logger.Printf(" %s %s", text, ref)
	} else {
		dwl.Logger.Printf(" %s", text)
	}

	resetPrefix()
}
