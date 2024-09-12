package garritsen

type ILogger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Warn(format string, args ...interface{})
}
