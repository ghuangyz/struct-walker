package swalker

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GetValueOf(data interface{}, key string) reflect.Value {
	key = sanitizeKey(key)
	value := flatten(reflect.ValueOf(data))
	return getValueOf(value, key)
}

func sanitizeKey(key string) string {
	buffer := make([]byte, len(key))
	lhs, rhs := 0, 0
	for rhs < len(key) {
		if key[rhs] != '.' || (lhs > 0 && buffer[lhs-1] != '.') {
			buffer[lhs] = key[rhs]
			lhs++
		}
		rhs++
	}
	if lhs > 0 && buffer[lhs-1] == '.' {
		lhs--
	}
	return string(buffer[:lhs])
}

// Recursion helper
func getValueOf(value reflect.Value, key string) reflect.Value {
	if len(key) == 0 {
		return value
	}

	var currentValue reflect.Value
	index := strings.Index(key, ".")
	if index == -1 {
		return resolveValue(value, key)
	}

	current := key[:index]
	next := key[(index + 1):]
	currentValue = resolveValue(value, current)
	return getValueOf(currentValue, next)
}

func resolveValue(value reflect.Value, key string) reflect.Value {
	var child reflect.Value
	if value.Kind() == reflect.Map {
		child = value.MapIndex(reflect.ValueOf(key))
	} else {
		if arrayKey, index, ok := deriveKeyAndIndex(key); ok {
			if arrayKey == "" {
				if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
					panic(fmt.Sprintf(
						"Asking for Array on key %s, but value is %v",
						key,
						value.Kind(),
					))
				}
				child = value.Index(index)
			} else {
				child = value.FieldByName(arrayKey).Index(index)
			}
		} else {
			child = value.FieldByName(key)
		}
	}
	return flatten(child)
}

func flatten(value reflect.Value) reflect.Value {
	// If the child is a pointer type, de-reference to actual value
	for value.Kind() == reflect.Ptr {
		value = reflect.Indirect(value)
	}
	return value
}

func deriveKeyAndIndex(key string) (string, int, bool) {
	left := strings.Index(key, "[")
	right := strings.LastIndex(key, "]")
	if left == -1 && right == -1 {
		return "", -1, false
	}

	// Panic if the [] are malformed
	if left == -1 || right == -1 || right < left {
		panic(fmt.Sprintf(
			"Malformed key(%s): contains [ or ] or both, but not well formed",
			key,
		))
	}

	realKey := key[:left]
	index, err := strconv.Atoi(key[(left + 1):right])
	// Panic if the index is not an integer
	if err != nil {
		panic(fmt.Sprintf(
			"Malformed key(%s): string between [] cannot be intepreted as integer",
			key,
		))
	}
	return realKey, index, true
}
