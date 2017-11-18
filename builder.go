package go_logger

const (
	contextFieldsKey = "context"
	logType          = "app"
)

type builder struct {
	fields        Fields
	contextFields Fields
}

// Builder initializes logger Fields builder
func Builder() builder {
	return builder{
		fields:        Fields{},
		contextFields: Fields{},
	}
}

// AddField adds a new field
func (builder builder) AddField(key string, value interface{}) builder {
	builder.fields.AddField(key, value)

	return builder
}

func (builder builder) AddFields(fields map[string]interface{}) builder {
	builder.fields.AddFields(fields)

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

func (builder builder) getFieldsWithDefaultValues(defaultFields Fields) map[string]interface{} {
	return builder.AddFields(defaultFields).getFields()
}

func (builder builder) getFields() map[string]interface{} {
	fields := builder.fields.getFieldsAsMap()
	if len(builder.contextFields) > 0 {
		fields[contextFieldsKey] = builder.contextFields.getFieldsAsMap()
	}

	if len(fields) == 0 {
		return nil
	}

	return fields
}
