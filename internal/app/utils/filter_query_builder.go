package utils

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Operator string

// Will be added in the future throughout development
const (
	OpEqual      Operator = "="
	OpNotEqual   Operator = "!="
	OpInArray    Operator = "IN"
	OpNotInArray Operator = "NOT IN"
	OpLike       Operator = "LIKE"
	OpILike      Operator = "ILIKE"
)

type EloquentQuery struct {
	Operator Operator
	Value    interface{}
}

func ApplyFilter(q *gorm.DB, filters map[string]EloquentQuery) *gorm.DB {
	if filters == nil {
		return q
	}

	for column, each := range filters {
		if each.Value == nil || reflect.ValueOf(each.Value).IsZero() {
			continue
		}

		q = q.Where(fmt.Sprintf("%s %s ?", column, each.Operator), each.Value)
	}

	return q
}

func GetExactMatchFilter(value interface{}) EloquentQuery {
	if value == "" || value == nil {
		return EloquentQuery{}
	}

	return EloquentQuery{
		Operator: "=",
		Value:    value,
	}
}
