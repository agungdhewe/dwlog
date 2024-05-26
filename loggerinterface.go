package dwlog

type LoggerInterface interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Log(msg string, args ...any)
}
