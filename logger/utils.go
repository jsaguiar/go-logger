package logger

func CreateFields(fields map[string]interface{}, extraFields ...extraField) Fields {
	if len(extraFields) > 0 && fields != nil {
		tmpExtraFields := make(map[string]interface{}, len(extraFields))
		for _, extraField := range extraFields {
			tmpExtraFields[extraField.key] = extraField.value
		}

		if len(tmpExtraFields) > 0 {
			fields["extra_info"] = tmpExtraFields
		}
	}

	return Fields(fields)
}

func CreateExtraField(key string, value interface{}) extraField {
	return extraField{key: key, value: value}
}
