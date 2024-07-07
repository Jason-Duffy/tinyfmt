// =============================================================================
// Project: tinyfmt
// File: print.go
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
	"io"
	"os"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// PrintToIo formats according to a format specifier and writes to the provided io.Writer.
func PrintToIo(w io.Writer, format string, arguments ...interface{}) error {
	result, err := Sprintf(format, arguments...)
	if err != nil {
		return errors.New("failed to format the string")
	}
	_, err = w.Write([]byte(result))
	return err
}

// Printf formats according to a format specifier and writes to os.Stdout.
func Printf(format string, arguments ...interface{}) error {
	return PrintToIo(os.Stdout, format, arguments...)
}
