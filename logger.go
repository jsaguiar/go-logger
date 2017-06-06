package go_logger

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/uniplaces/go-logger/internal"
	"github.com/uniplaces/go-logger/logger"
)

var instance logger.Logger
var once sync.Once

// Init initializes logger instance
func Init(config Config) error {
	if instance != nil {
		return errors.New("logger cannot be initialized twice")
	}

	once.Do(func() {
		// todo use implementation according to the env
		instance = internal.NewLogrusLogger(config.level, os.Stdout)
	})

	return nil
}

// InitWithInstance sets logger to an instance
func InitWithInstance(newInstance logger.Logger) {
	instance = newInstance
}

func ErrorWithFields(message string, fields logger.Fields) {
	if instance == nil {
		return
	}

	instance.ErrorWithFields(message, fields)
}

func Errorf(message string, parameters ...interface{}) {
	Error(fmt.Sprintf(message, parameters...))
}

func Error(message string) {
	if instance == nil {
		return
	}

	instance.Error(message)
}

func WarningWithFields(message string, fields logger.Fields) {
	if instance == nil {
		return
	}

	instance.WarningWithFields(message, fields)
}

func Warningf(message string, parameters ...interface{}) {
	Warning(fmt.Sprintf(message, parameters...))
}

func Warning(message string) {
	if instance == nil {
		return
	}

	instance.Warning(message)
}

func InfoWithFields(message string, fields logger.Fields) {
	if instance == nil {
		return
	}

	instance.InfoWithFields(message, fields)
}

func Infof(message string, parameters ...interface{}) {
	Info(fmt.Sprintf(message, parameters...))
}

func Info(message string) {
	if instance == nil {
		return
	}

	instance.Info(message)
}

func DebugWithFields(message string, fields logger.Fields) {
	if instance == nil {
		return
	}

	instance.DebugWithFields(message, fields)
}

func Debugf(message string, parameters ...interface{}) {
	Debug(fmt.Sprintf(message, parameters...))
}

func Debug(message string) {
	if instance == nil {
		return
	}

	instance.Debug(message)
}
