package utils

import (
	"reflect"
	"strings"
)

func FilterAllowedPayloadFields(data any, allowed []string) map[string]interface{} {
	result := map[string]interface{}{}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		for _, col := range allowed {
			if col == jsonTag {
				result[col] = val.Field(i).Interface()
			}
		}
	}
	return result
}

func FilterEmptyFields(data map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	for key, value := range data {
		if value == nil || reflect.ValueOf(value).IsZero() {
			continue
		}

		result[key] = value
	}

	return result
}
