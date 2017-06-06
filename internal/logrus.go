package internal

import (
	"fmt"
	"io"
	"log"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/uniplaces/go-logger/logger"
)

// defines which levels log stacktrace
var stackTraceLevels = map[logrus.Level]bool{
	logrus.DebugLevel: false,
	logrus.InfoLevel:  false,
	logrus.WarnLevel:  false,
	logrus.ErrorLevel: true,
	logrus.FatalLevel: true,
	logrus.PanicLevel: true,
}

type logrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger(level string, writer io.Writer) logger.Logger {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(fmt.Sprintf("invalid log level (%s)", level))
	}

	instance := &logrusLogger{
		Logger: logrus.New(),
	}

	instance.Formatter = &logrus.JSONFormatter{}
	instance.Level = logrusLevel
	instance.Out = writer

	log.SetOutput(instance.Writer())

	return instance
}

func (logger *logrusLogger) ErrorWithFields(message string, fields logger.Fields) {
	entry := logger.entry(fields, stackTraceLevels[logrus.ErrorLevel])
	entry.Error(message)
}

func (logger *logrusLogger) Error(message string) {
	logger.ErrorWithFields(message, nil)
}

func (logger *logrusLogger) WarningWithFields(message string, fields logger.Fields) {
	entry := logger.entry(fields, stackTraceLevels[logrus.WarnLevel])
	entry.Warning(message)
}

func (logger *logrusLogger) Warning(message string) {
	logger.WarningWithFields(message, nil)
}

func (logger *logrusLogger) InfoWithFields(message string, fields logger.Fields) {
	entry := logger.entry(fields, stackTraceLevels[logrus.InfoLevel])
	entry.Info(message)
}

func (logger *logrusLogger) Info(message string) {
	logger.InfoWithFields(message, nil)
}

func (logger *logrusLogger) DebugWithFields(message string, fields logger.Fields) {
	entry := logger.entry(fields, stackTraceLevels[logrus.DebugLevel])
	entry.Debug(message)
}

func (logger *logrusLogger) Debug(message string) {
	logger.DebugWithFields(message, nil)
}

func (logger *logrusLogger) entry(fields logger.Fields, includeStackTrace bool) *logrus.Entry {
	stackTrace := buildStackTrace()

	if len(fields) == 0 {
		logFields := logrus.Fields{}
		if includeStackTrace {
			logFields["stacktrace"] = stackTrace
		}

		return logger.WithFields(logFields)
	}

	logFields := logrus.Fields(fields)
	if includeStackTrace {
		logFields["stacktrace"] = stackTrace
	}

	return logger.WithFields(logFields)
}

func buildStackTrace() []string {
	var stacktrace []string

	skip := 0
	for {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}

		skip++

		stacktrace = append(stacktrace, fmt.Sprintf("%s on line %d", file, line))
	}

	return stacktrace
}
