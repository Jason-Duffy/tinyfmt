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

- Go 1.20 or later

## Installation

To install `tinyfmt`, use `go get`:

```sh
go get github.com/Jason-Duffy/tinyfmt
```

## Usage

Here are some examples of how to use `tinyfmt`:

### Error Handling

All functions return an error as the second return value. You can discard this error if you don't need it by using `_`:

```go
result, err := tinyfmt.Sprintf("Hello, %s!", "world")
if err != nil {
    println("Error:", err.Error())
} else {
    println(result)
}

// Discarding the error
result, _ = tinyfmt.Sprintf("Hello, %s!", "world")
println(result)
```

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
	result, _ := tinyfmt.Sprintf("Hello, %s!", "world")
	println(result)

	result, _ = tinyfmt.Sprintf("Value: %d", 42)
	println(result)

	result, _ = tinyfmt.Sprintf("Hex: %x", 255)
	println(result)

	result, _ = tinyfmt.Sprintf("Binary: %b", 7)
	println(result)

	result, _ = tinyfmt.Sprintf("Octal: %o", 64)
	println(result)

	result, _ = tinyfmt.Sprintf("Float: %.2f", 3.14159)
	println(result)

	result, _ = tinyfmt.Sprintf("Bool: %v", true)
	println(result)
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
	tinyfmt.Printf("Hello, %s!", "world")
	tinyfmt.Printf("Value: %d", 42)
	tinyfmt.Printf("Hex: %x", 255)
	tinyfmt.Printf("Binary: %b", 7)
	tinyfmt.Printf("Octal: %o", 64)
	tinyfmt.Printf("Float: %.2f", 3.14159)
	tinyfmt.Printf("Bool: %v", true)
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
	tinyfmt.PrintToIO(&buf, "Hello, %s!", "world")
	println(buf.String())

	// Printing to standard output
	tinyfmt.PrintToIO(os.Stdout, "Value: %d\n", 42)
}
```

## Code Size

Using `tinyfmt` results in significantly smaller code size compared to the standard library. When built with TinyGo for a Pico target, the code size increase when using `tinyfmt` was approximately **1.5kB**, compared to **40kB** when using the Go `fmt` package.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
