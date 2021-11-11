package logger

type nopLogger struct{}

func NewNopLogger() Logger {
	return &nopLogger{}
}

func (logger nopLogger) With(args ...interface{}) Logger {
	return logger
}

func (_ nopLogger) Debug(message string, args ...interface{}) {
}
func (_ nopLogger) Info(message string, args ...interface{}) {
}
func (_ nopLogger) Error(message string, args ...interface{}) {
}
func (_ nopLogger) Panic(message string, args ...interface{}) {
}
func (_ nopLogger) Fatal(message string, args ...interface{}) {
}
func (_ nopLogger) Latency() func(message string, args ...interface{}) {
	return func(message string, args ...interface{}) {}
}
