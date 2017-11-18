package go_logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Parallel()

	expectedConfig := Config{environment: "test", level: "warning"}
	config := NewConfig("test", "warning")

	assert.Equal(t, config, expectedConfig)
}

func TestNewConfigWithProjectFields(t *testing.T) {
	t.Parallel()

	expectedConfig := Config{environment: "test", level: "warning", projectFields: Fields{}.AddField("key", "value")}
	config := NewConfigWithProjectFields("test", "warning", Fields{}.AddField("key", "value"))

	assert.Equal(t, config, expectedConfig)
}
