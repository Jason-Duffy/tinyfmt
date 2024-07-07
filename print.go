package tinyfmt

import (
	"errors"
	"io"
	"os"
)

// PrintToIo formats according to a format specifier and writes to the provided io.Writer.
func PrintToIo(w io.Writer, format string, arguments ...interface{}) error {
	result, err := Format(format, arguments...)
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
