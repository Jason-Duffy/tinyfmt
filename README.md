# tinyfmt

`tinyfmt` is a lightweight Go package that provides basic string formatting and printing capabilities. It avoids importing large standard libraries like `fmt` and `strconv`, making it ideal for small and embedded applications where binary size is a concern.

## Features

- **Sprint**: Concatenate strings and convert different types to string.
- **Sprintf**: Format strings with various format specifiers.
- **Printf**: Print formatted strings to the standard output.
- **PrintToIO**: Print formatted strings to a specified `io.Writer`.

## Goals

The main goals of `tinyfmt` are:

1. **Lightweight**: Minimize binary size by avoiding large standard libraries.
2. **Flexibility**: Provide basic formatting and printing functions.
3. **Simplicity**: Keep the API simple and easy to use.

## Requirements

- Go 1.15 or later

## Installation

To install `tinyfmt`, use `go get`:

```sh
go get github.com/Jason-Duffy/tinyfmt
```

## Usage

Here are some examples of how to use `tinyfmt`:

### Sprint

`Sprint` concatenates strings and converts different types to string.

```go
package main

import (
	"github.com/Jason-Duffy/tinyfmt"
)

func main() {
	result := tinyfmt.Sprint("Hello, ", "world!")
	println(result)

	result = tinyfmt.Sprint("Value: ", 42)
	println(result)

	result = tinyfmt.Sprint("Bool: ", true)
	println(result)

	result = tinyfmt.Sprint("Float: ", 3.14159)
	println(result)

	result = tinyfmt.Sprint("Mixed: ", "string", ", ", 123, ", ", false)
	println(result)
}
```

### Sprintf

`Sprintf` formats strings with various format specifiers.

```go
package main

import (
	"github.com/Jason-Duffy/tinyfmt"
)

func main() {
	result, err := tinyfmt.Sprintf("Hello, %s!", "world")
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}

	result, err = tinyfmt.Sprintf("Value: %d", 42)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}

	result, err = tinyfmt.Sprintf("Hex: %x", 255)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}

	result, err = tinyfmt.Sprintf("Binary: %b", 7)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}

	result, err = tinyfmt.Sprintf("Octal: %o", 64)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}

	result, err = tinyfmt.Sprintf("Float: %.2f", 3.14159)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}

	result, err = tinyfmt.Sprintf("Bool: %v", true)
	if err != nil) {
		println("Error:", err.Error())
	} else {
		println(result)
	}
}
```

### Printf

`Printf` prints formatted strings to the standard output.

```go
package main

import (
	"github.com/Jason-Duffy/tinyfmt"
)

func main() {
	err := tinyfmt.Printf("Hello, %s!", "world")
	if err != nil {
		println("Error:", err.Error())
	}

	err = tinyfmt.Printf("Value: %d", 42)
	if err != nil {
		println("Error:", err.Error())
	}

	err = tinyfmt.Printf("Hex: %x", 255)
	if err != nil {
		println("Error:", err.Error())
	}

	err = tinyfmt.Printf("Binary: %b", 7)
	if err != nil {
		println("Error:", err.Error())
	}

	err = tinyfmt.Printf("Octal: %o", 64)
	if err != nil {
		println("Error:", err.Error())
	}

	err = tinyfmt.Printf("Float: %.2f", 3.14159)
	if err != nil {
		println("Error:", err.Error())
	}

	err = tinyfmt.Printf("Bool: %v", true)
	if err != nil {
		println("Error:", err.Error())
	}
}
```

### PrintToIO

`PrintToIO` prints formatted strings to a specified `io.Writer`.

```go
package main

import (
	"bytes"
	"github.com/Jason-Duffy/tinyfmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	err := tinyfmt.PrintToIO(&buf, "Hello, %s!", "world")
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(buf.String())
	}

	// Printing to standard output
	err = tinyfmt.PrintToIO(os.Stdout, "Value: %d\n", 42)
	if err != nil {
		println("Error:", err.Error())
	}
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
