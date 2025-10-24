package mathutils

import "math/rand"

// Sum returns the sum of integers
func Sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// RandomInt returns a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
