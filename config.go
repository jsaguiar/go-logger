package go_logger

// Config defines logger config
type Config struct {
	environment   string
	level         string
	projectFields Fields
}

// NewConfig returns initialized config
func NewConfig(environment, level string) Config {
	return Config{
		environment: environment,
		level:       level,
	}
}

func NewConfigWithProjectFields(environment, level string, projectFields Fields) Config {
	return Config{
		environment:   environment,
		level:         level,
		projectFields: projectFields,
	}
}
