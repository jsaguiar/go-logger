package go_logger

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestNewConfig(t *testing.T) {
	expectedConfig := Config{environment: "test", level: "warning"}
	config := NewConfig("test", "warning")

	assert.Equal(t, config, expectedConfig)
}