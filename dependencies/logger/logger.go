package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LmptLogger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

func (l LmptLogger) Debug(args ...interface{}) {
	l.logger.WithFields(l.fields).Debug(args...)
}

func (l LmptLogger) Debugf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Debugf(format, args...)
}

func (l LmptLogger) Info(args ...interface{}) {
	l.logger.WithFields(l.fields).Info(args...)
}

func (l LmptLogger) Infof(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Infof(format, args...)
}

func (l LmptLogger) Error(args ...interface{}) {
	l.logger.WithFields(l.fields).Error(args...)
}

func (l LmptLogger) Errorf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Errorf(format, args...)
}

func (l LmptLogger) Warning(args ...interface{}) {
	l.logger.WithFields(l.fields).Warning(args...)
}

func (l LmptLogger) Warningf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Warningf(format, args...)
}

func (l LmptLogger) WithError(err error, message string) {
	l.logger.WithFields(l.fields).WithError(err).Error(message)
}

func (l LmptLogger) WithErrorf(err error, format string, args ...interface{}) {
	l.logger.WithFields(l.fields).WithError(err).Errorf(format, args...)
}

func NewTextLogger() LmptLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	setLevel(logger)

	return LmptLogger{
		logger: logger,
		fields: logrus.Fields{},
	}
}

func NewJsonLogger() LmptLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	setLevel(logger)

	return LmptLogger{
		logger: logger,
		fields: getEnvFields(),
	}
}

func setLevel(logger *logrus.Logger) {
	logLevel := os.Getenv("LOG_LEVEL")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
}

func getEnvFields() logrus.Fields {
	env := os.Getenv("DD_ENV")
	if env == "" {
		env = "dev"
	}

	fields := logrus.Fields{
		// 	"bu_code":            os.Getenv("BU_CODE"),
		// 	"bu_ownership":       os.Getenv("BU_OWNERSHIP"),
		"env": env,
		// 	"project_tangram":    os.Getenv("PROJECT_TANGRAM"),
		// 	"project_tangram_id": os.Getenv("PROJECT_TANGRAM_ID"),
		"data_privacy": 0,
	}

	return fields
}
