// =============================================================================
// Project: tinyfmt
// File: sprint.go
// Description: Functions for printing formatted strings to io.Writer and os.Stdout.
// Datasheet/Docs:
//
// Author: Jason Duffy
// Created on: 07/07/2024
//
// Copyright: (C) 2024, Jason Duffy
// License: See LICENSE file in the project root for full license information.
// Disclaimer: See DISCLAIMER file in the project root for full disclaimer.
// =============================================================================

// -------------------------------------------------------------------------- //
//                               Import Statement                             //
// -------------------------------------------------------------------------- //

package tinyfmt

import (
	"errors"
	"reflect"

	"github.com/Jason-Duffy/tinystrconv"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// Sprint concatenates the string representations of the provided arguments.
func Sprint(arguments ...interface{}) string {
	var result string
	for _, argument := range arguments {
		switch value := argument.(type) {
		case string:
			result += value
		case int:
			str, _ := tinystrconv.IntToString(value, 10)
			result += str
		case bool:
			result += tinystrconv.BoolToString(value)
		case float64:
			str, _ := tinystrconv.FloatToString(value, -1) // Use -1 for full precision
			result += str
		default:
			result += formatUnsupported(value)
		}
	}
	return result
}

// Sprintf formats the provided arguments according to the format specifier.
func Sprintf(format string, arguments ...interface{}) (string, error) {
	var result []byte
	argIndex := 0

	for i := 0; i < len(format); i++ {
		if format[i] == '%' {
			if i+1 < len(format) {
				i++

				// Handle precision for floats (e.g., "%.2f")
				precision := -1
				if format[i] == '.' {
					i++
					start := i
					for i < len(format) && format[i] >= '0' && format[i] <= '9' {
						i++
					}
					if start < i {
						precision = 0
						for j := start; j < i; j++ {
							precision = precision*10 + int(format[j]-'0')
						}
					}
				}

				// Handle different format specifiers
				switch format[i] {
				case 'd':
					if argIndex >= len(arguments) {
						return "", errors.New("missing argument for %d")
					}
					intVal, ok := arguments[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %d is not an int")
					}
					str, err := tinystrconv.IntToString(intVal, 10)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'f':
					if argIndex >= len(arguments) {
						return "", errors.New("missing argument for %f")
					}
					floatVal, ok := arguments[argIndex].(float64)
					if !ok {
						return "", errors.New("argument for %f is not a float64")
					}
					str, err := tinystrconv.FloatToString(floatVal, precision)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'b':
					if argIndex >= len(arguments) {
						return "", errors.New("missing argument for %b")
					}
					intVal, ok := arguments[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %b is not an int")
					}
					str, err := tinystrconv.IntToString(intVal, 2)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'x':
					if argIndex >= len(arguments) {
						return "", errors.New("missing argument for %x")
					}
					intVal, ok := arguments[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %x is not an int")
					}
					str, err := tinystrconv.IntToString(intVal, 16)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'o':
					if argIndex >= len(arguments) {
						return "", errors.New("missing argument for %o")
					}
					intVal, ok := arguments[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %o is not an int")
					}
					str, err := tinystrconv.IntToString(intVal, 8)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'v':
					if argIndex >= len(arguments) {
						return "", errors.New("missing argument for %v")
					}
					str, err := formatValue(arguments[argIndex])
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 's':
					if argIndex >= len(arguments) {
						return "", errors.New("missing argument for %s")
					}
					strVal, ok := arguments[argIndex].(string)
					if !ok {
						return "", errors.New("argument for %s is not a string")
					}
					result = append(result, []byte(strVal)...)
					argIndex++
				case '%':
					result = append(result, '%')
				default:
					return "", errors.New("unsupported format specifier")
				}
			} else {
				return "", errors.New("incomplete format specifier at end of string")
			}
		} else {
			result = append(result, format[i])
		}
	}

	return string(result), nil
}

// -------------------------------------------------------------------------- //
//                             Private Functions                              //
// -------------------------------------------------------------------------- //

// formatUnsupported handles the formatting of unsupported types.
func formatUnsupported(value interface{}) string {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Struct:
		return formatStruct(v)
	case reflect.Slice, reflect.Array:
		return formatSlice(v)
	case reflect.Map:
		return formatMap(v)
	default:
		return "<unsupported>"
	}
}

// formatValue handles the formatting of %v format specifier.
func formatValue(value interface{}) (string, error) {
	switch value := value.(type) {
	case bool:
		return tinystrconv.BoolToString(value), nil
	case string:
		return value, nil
	case int:
		return tinystrconv.IntToString(value, 10)
	case float64:
		return tinystrconv.FloatToString(value, -1)
	default:
		return formatUnsupported(value), nil
	}
}

// formatStruct formats a struct as a string.
func formatStruct(v reflect.Value) string {
	result := "{"
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			result += " "
		}
		field := v.Type().Field(i).Name
		value := v.Field(i).Interface()
		result += field + ":" + Sprint(value)
	}
	result += "}"
	return result
}

// formatSlice formats a slice or array as a string.
func formatSlice(v reflect.Value) string {
	result := "["
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			result += " "
		}
		result += Sprint(v.Index(i).Interface())
	}
	result += "]"
	return result
}

// formatMap formats a map as a string.
func formatMap(v reflect.Value) string {
	result := "{"
	keys := v.MapKeys()
	for i, key := range keys {
		if i > 0 {
			result += " "
		}
		value := v.MapIndex(key).Interface()
		result += Sprint(key.Interface()) + ":" + Sprint(value)
	}
	result += "}"
	return result
}
