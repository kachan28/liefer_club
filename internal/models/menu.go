package models

import (
	"reflect"
	"strings"
)

type Menu struct {
	Name       string      `json:"menu name"`
	Id         int64       `json:"menu id"`
	Db         string      `json:"-"`
	DishGroups []DishGroup `json:"dish groups"`
}

func (m *Menu) ScanField(column string, value string) {
	const jsonKey = "json"
	fieldIndex := -1
	rt := reflect.TypeOf(*m)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tagValue := strings.Split(field.Tag.Get(jsonKey), ",")[0]
		if tagValue == column {
			fieldIndex = i
		}
	}
	if fieldIndex != -1 {
		reflect.ValueOf(m).Elem().Field(fieldIndex).SetString(value)
	}
}
