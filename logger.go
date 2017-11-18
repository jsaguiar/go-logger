package go_logger

import (
	"errors"
	"os"
	"sync"

	"github.com/jsaguiar/go-logger/internal"
)

var instance Logger
var defaultFields Fields
var once sync.Once

// Init initializes logger instance
func Init(config Config) error {
	if instance != nil {
		return errors.New("logger cannot be initialized more than once")
	}

	InitWithInstance(internal.NewLogrusLogger(config.level, os.Stdout), config.projectFields)

	return nil
}

// InitWithInstance sets logger to an instance (for testing purposes)
func InitWithInstance(newInstance Logger, newDefaultFields Fields) error {
	if instance != nil {
		return errors.New("logger cannot be initialized more than once")
	}

	once.Do(func() {
		instance = newInstance
		defaultFields = newDefaultFields
	})

	return nil
}

// Error logs a error message
func Error(err error) {
	Builder().Error(err)
}

// Error logs a error message with Fields
func (builder builder) Error(err error) {
	if instance == nil {
		return
	}

	instance.ErrorWithFields(err.Error(), builder.getFieldsWithDefaultValues(defaultFields))
}

// Warning logs a warning message
func Warning(message string) {
	Builder().Warning(message)
}

// Warning logs a warning message with Fields
func (builder builder) Warning(message string) {
	if instance == nil {
		return
	}

	instance.WarningWithFields(message, builder.getFieldsWithDefaultValues(defaultFields))
}

// Info logs a info message
func Info(message string) {
	Builder().Info(message)
}

// Info logs a info message with Fields
func (builder builder) Info(message string) {
	if instance == nil {
		return
	}

	instance.InfoWithFields(message, builder.getFieldsWithDefaultValues(defaultFields))
}

// Debug logs a debug message
func Debug(message string) {
	Builder().Debug(message)
}

// Debug logs a debug message with Fields
func (builder builder) Debug(message string) {
	if instance == nil {
		return
	}

	instance.DebugWithFields(message, builder.getFieldsWithDefaultValues(defaultFields))
}
