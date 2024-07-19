package utils

import "reflect"

func StructToMapWithJSONKeys(input interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeField := t.Field(i)
		jsonTag := typeField.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			jsonTag = typeField.Name
		}
		result[jsonTag] = field.Interface()
	}

	return result, nil
}
