package stringutils

import "strings"

// Reverse reverses a string
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	return s == Reverse(s)
}
