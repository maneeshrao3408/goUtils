# go-utils

`go-utils` is a general-purpose Go utility library providing helpful functions for strings, math, time, and friendly greetings.  
It is designed to be lightweight, easy to use, and modular.

## Features

- **String utilities** (`stringutils`)
  - `Reverse(string) string` → reverses a string
  - `IsPalindrome(string) bool` → checks if a string is a palindrome
  - `ToSnakeCase(string) string` → converts camelCase to snake_case

- **Math utilities** (`mathutils`)
  - `Sum(numbers ...int) int` → sum of integers
  - `Average(numbers ...float64) float64` → average of numbers
  - `RandomInt(min, max int) int` → random integer in range

- **Time utilities** (`timeutils`)
  - `DaysBetween(start, end time.Time) int` → number of days between two dates
  - `IsWeekend(time.Time) bool` → checks if a date is on the weekend

- **Greet utilities** (`greet`)
  - `Hello(name string)` → prints a friendly greeting
  - `Goodbye(name string)` → prints a farewell message

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
	fmt.Println("Average:", mathutils.Average(2.0, 4.0, 6.0))
	fmt.Println("Reverse:", stringutils.Reverse("hello"))
	fmt.Println("Palindrome:", stringutils.IsPalindrome("madam"))
}
```

## Output
```bash
Sum: 10
Average: 4
Reverse: olleh
Palindrome: true

