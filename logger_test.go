package go_logger

import (
	"bytes"
	"errors"
	"testing"

	"sync"

	"github.com/jsaguiar/go-logger/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	assert.Nil(t, instance)

	config := NewConfig("env", "warning")
	Init(config)

	assert.NotNil(t, instance)

	resetInstance()
}

func TestLogWithFields(t *testing.T) {
	var buffer bytes.Buffer
	err := InitWithInstance(internal.NewLogrusLogger("error", &buffer), Fields{}.AddField("project", "default"))
	require.Nil(t, err)

	Builder().
		AddField("key", "value").
		AddContextField("foo", "bar").
		Error(errors.New("error test"))

	assert.Contains(t, buffer.String(), "\"context\":{\"foo\":\"bar\"}")
	assert.Contains(t, buffer.String(), "\"key\":\"value\"")
	assert.Contains(t, buffer.String(), "\"stack_trace\"")
	assert.Contains(t, buffer.String(), "\"project\":\"default\"")

	resetInstance()
}

func TestLog(t *testing.T) {
	var buffer bytes.Buffer
	err := InitWithInstance(internal.NewLogrusLogger("error", &buffer), Fields{})
	require.Nil(t, err)

	Error(errors.New("test error"))

	assert.NotContains(t, buffer.String(), "\"context\"")
	assert.NotContains(t, buffer.String(), "\"key\":\"value\"")
	assert.Contains(t, buffer.String(), "\"stack_trace\"")

	resetInstance()
}

func resetInstance() {
	instance = nil
	defaultFields = Fields{}
	once = sync.Once{}
}
