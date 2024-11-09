# gorange

A Go library for parsing number ranges from strings with flexible format support.

## Features

- Parse both single numbers and ranges (e.g., "1-3")
- Ignore non-digit noise characters
- Handle extra separators and spaces
- Support Chinese comma (，)
- Auto-sort range numbers
- Flexible format support

## Usage

```go
package main

import "github.com/huhu415/gorange"

func main() {
    numbers, err := gorange.ExtractRange("1-3,5,7-9")
    if err != nil {
        panic(err)
    }
    fmt.Println(numbers) // Output: [1 2 3 5 7 8 9]
}
```

## Examples

```go
// Basic usage
"1-3" => [1,2,3]
"1,2,3" => [1,2,3]

// Ignore noise
"xxx1-3xxx" => [1,2,3]

// Handle extra separators
"1,,2,,,3" => [1,2,3]

// Handle extra dashes
"1----3" => [1,2,3]

// Handle spaces
"  1  -  3  " => [1,2,3]

// Handle empty input
"" => []

// Auto-sort range
"3-1" => [1,2,3]

// Support Chinese comma
"1，3-5，7" => [1,3,4,5,7]
```

## API

### func ExtractRange(input string) ([]int, error)

ExtractRange parses number ranges from a string and returns a slice of integers.

Parameters:
- `input`: A string containing numbers and ranges

Returns:
- `[]int`: Slice containing all extracted numbers
- `error`: Parsing errors (currently always returns nil)
