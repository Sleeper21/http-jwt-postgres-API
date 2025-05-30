package domain

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	WithError(err error, message string)
	WithErrorf(err error, format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
}
