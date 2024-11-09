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

import (
	"fmt"

	"github.com/huhu415/gorange"
)

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
"xxx3-1xxx, xjlkjfd13slsv-sdf" => [1,2,3,13]

// Handle extra separators and Handle spaces
"1,,2,,,   3  " => [1,2,3]

// Handle extra dashes, and Support Chinese comma
"   1----   3  ，5 " => [1,2,3]

// Handle empty input
"" => []
```
