package go_logger

// Fields is used to keep track of the values to log
type Fields map[string]interface{}

// AddField adds a new field
func (fields Fields) AddField(key string, value interface{}) Fields {
	fields[key] = value

	return fields
}

// AddFields adds a map of Fields
func (fields Fields) AddFields(fieldsMap map[string]interface{}) Fields {
	for k, v := range fieldsMap {
		fields.AddField(k, v)
	}

	return fields
}

func (fields Fields) getFieldsAsMap() map[string]interface{} {
	return fields
}
