package utils

import (
	"reflect"
	"strings"
)

func ExtractJSONFields[T any]() []string {
	var columns []string
	typ := reflect.TypeOf((*T)(nil)).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// Ambil nama dari tag `json:"name"`
		jsonTag := field.Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0] // handle `json:"name,omitempty"`

		if jsonName != "" && jsonName != "-" {
			columns = append(columns, jsonName)
		}
	}

	return columns
}
