// =============================================================================
// Project: tinyfmt
// File: error.go
// Description: Functions for formatting error messages.
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
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// Errorf formats according to a format specifier and returns the string as a value that satisfies error.
func Errorf(format string, arguments ...interface{}) error {
	result, err := Sprintf(format, arguments...)
	if err != nil {
		return err
	}
	return errors.New(result)
}
