package go_logger

import "os"

const (
	contextFieldsKey = "context"
	logType          = "app"
)

type builder struct {
	fields        fields
	contextFields fields
}

// Builder initializes logger fields builder
func Builder() builder {
	return builder{
		fields:        fields{},
		contextFields: fields{},
	}
}

// AddField adds a new field
func (builder builder) AddField(key string, value interface{}) builder {
	builder.fields.AddField(key, value)

	return builder
}

func (builder builder) AddFields(fields map[string]interface{}) builder {
	builder.AddFields(fields)

	return builder
}

// AddContextField adds a new context field
func (builder builder) AddContextField(key string, value interface{}) builder {
	builder.contextFields.AddField(key, value)

	return builder
}

// AddContextField adds a new context field
func (builder builder) AddContextFields(fields map[string]interface{}) builder {
	builder.contextFields.AddFields(fields)

	return builder
}

func (builder builder) getFieldsWithMandatoryKeys() map[string]interface{} {
	builder.
		AddField("type", logType).
		AddField("app-id", os.Getenv("APPID")).
		AddField("env", os.Getenv("GOENV")).
		AddField("git-hash", os.Getenv("GITHASH"))

	return builder.getFields()
}

func (builder builder) getFields() map[string]interface{} {
	fields := builder.fields
	if len(builder.contextFields) > 0 {
		fields[contextFieldsKey] = builder.contextFields
	}

	if len(fields) == 0 {
		return nil
	}

	return fields
}
