package dwlog

import (
	"fmt"
	"io"
	"log"
	"runtime"
)

func resetPrefix() {
	dwlog.Logger.SetPrefix("")
	dwlog.Logger.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func setColor(colorCode string, text string) string {
	return fmt.Sprintf("%s%s%s", colorCode, text, ColorReset)
}

func getCaller() (string, int) {
	if dwlog.Logger.Writer() == io.Discard {
		return "", 0
	}

	_, callerfile, lineno, ok := runtime.Caller(2)
	if !ok {
		return "", 0
	}

	return callerfile, lineno
}
