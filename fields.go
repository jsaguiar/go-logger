package go_logger

type fields map[string]interface{}


// AddField adds a new field
func (fields fields) AddField(key string, value interface{}) fields {
	fields[key] = value

	return fields
}

// AddFields adds a map of fields
func (fields fields) AddFields(fieldsMap map[string]interface{}) fields {
	for k, v := range fieldsMap {
		fields.AddField(k, v)
	}

	return fields
}
