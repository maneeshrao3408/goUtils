# go-utils

`go-utils` is a general-purpose Go utility library providing helpful functions for strings, math, time, and friendly greetings.  
It is designed to be lightweight, easy to use, and modular.

## Features

- **String utilities** (`stringutils`)
  - `Reverse(string) string` → reverses a string
  - `IsPalindrome(string) bool` → checks if a string is a palindrome

- **Math utilities** (`mathutils`)
  - `Sum(numbers ...int) int` → sum of integers
  - `RandomInt(min, max int) int` → random integer in range

---

## Installation

Use `go get` to install the package:

```bash
go get github.com/maneeshrao3408/go-utils@v1.0.0

```

## Usage

```bash
package main

import (
	"fmt"
	"github.com/maneeshrao3408/go-utils/mathutils"
	"github.com/maneeshrao3408/go-utils/stringutils"
)

func main() {
	fmt.Println("Sum:", mathutils.Sum(1, 2, 3, 4))
	fmt.Println("Random Number:", mathutils.RandomInt(1, 100))
	fmt.Println("Reverse:", stringutils.Reverse("hello"))
	fmt.Println("Palindrome:", stringutils.IsPalindrome("madam"))
}
```

## Output
```bash
Sum: 10
Random Number: 73
Reverse: olleh
Palindrome: true

