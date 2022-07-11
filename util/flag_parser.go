package util

import (
	"fmt"
	"reflect"
)

type Options = map[string]interface{}

func GetOptions[T any](model T) Options {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	m := make(map[string]interface{})

	// Iterating over struct fields (aka keys)
	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i)
		field := t.Field(i)
		tag := field.Tag

		// Skip if there are no tags.
		if len(tag) == 0 {
			continue
		}

		flag := tag.Get("flag")

		if _, ok := m[flag]; !ok {
			m[flag] = value
		}

	}

	return m
}

func StringifyOptions(options Options) []string {
	var values []string

	for flag, value := range options {
		flagPrefix := "--"

		if len(flag) == 1 {
			flagPrefix = "-"
		}

		stringValue := fmt.Sprint(value)
		extendedFlag := fmt.Sprintf("%s%s", flagPrefix, flag)

		if isBoolean(stringValue) {
			if stringValue == "true" {
				values = append(values, extendedFlag)
			}

			continue
		}

		if stringValue != "" {
			values = append(values, extendedFlag, stringValue)
		}
	}

	return values
}

func isBoolean(str string) bool {
	return str == "true" || str == "false"
}
