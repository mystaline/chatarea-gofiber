package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func ExtractSelectColumns[T any]() []string {
	return extractColumnsRecursive(reflect.TypeOf((*T)(nil)).Elem(), "")
}

func extractColumnsRecursive(t reflect.Type, prefix string) []string {
	var columns []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct && !field.Anonymous {
			columns = append(columns, extractColumnsRecursive(field.Type, "")...)
			continue
		}

		// Ambil nama dari tag `json:"name"`
		gormTag := field.Tag.Get("gorm")
		alias := getAliasTag(gormTag)
		column := getColumnFromTag(gormTag)

		if column == "" {
			continue
		}

		if alias != "" {
			columns = append(columns, fmt.Sprintf("%s as %s", alias, column))
		} else {
			columns = append(columns, column)
		}
	}
	fmt.Println("columns", columns)
	return columns
}

// func ExtractSelectColumns[T any]() []string {
// 	var columns []string
// 	typ := reflect.TypeOf((*T)(nil)).Elem()

// 	for i := 0; i < typ.NumField(); i++ {
// 		field := typ.Field(i)

// 		gormTag := field.Tag.Get("gorm")
// 		alias := getAliasTag(gormTag)
// 		column := getColumnFromTag(gormTag)

// 		if column == "" {
// 			continue
// 		}

// 		if alias != "" {
// 			columns = append(columns, fmt.Sprintf("%s as %s", alias, column))
// 		} else {
// 			columns = append(columns, column)
// 		}
// 	}
// 	fmt.Println("columns", columns)
// 	return columns
// }

func getAliasTag(tag string) string {
	splitted := strings.Split(tag, ";")
	for _, each := range splitted {
		if strings.HasPrefix(each, "alias:") {
			return strings.TrimPrefix(each, "alias:")
		}
	}

	return ""
}

func getColumnFromTag(tag string) string {
	splitted := strings.Split(tag, ";")
	for _, each := range splitted {
		if strings.HasPrefix(each, "column:") {
			return strings.TrimPrefix(each, "column:")
		}
	}

	return ""
}
