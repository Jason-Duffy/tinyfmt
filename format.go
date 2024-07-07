// =============================================================================
// Project: tinyfmt
// File: format.go
// Description: Implementation of the Format function in the tinyfmt package.
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

	"github.com/Jason-Duffy/tinystrconv"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// Format formats a string with the provided arguments, similar to fmt.Sprintf.
func Format(format string, args ...interface{}) (string, error) {
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
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %d")
					}
					intVal, ok := args[argIndex].(int)
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
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %f")
					}
					floatVal, ok := args[argIndex].(float64)
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
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %b")
					}
					intVal, ok := args[argIndex].(int)
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
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %x")
					}
					intVal, ok := args[argIndex].(int)
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
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %o")
					}
					intVal, ok := args[argIndex].(int)
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
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %v")
					}
					boolVal, ok := args[argIndex].(bool)
					if !ok {
						return "", errors.New("argument for %v is not a bool")
					}
					str := tinystrconv.BoolToString(boolVal)
					result = append(result, []byte(str)...)
					argIndex++
				case 's':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %s")
					}
					strVal, ok := args[argIndex].(string)
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
