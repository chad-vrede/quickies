package leetcode

import (
	"strings"
	"unicode"
)

// Valid Palindrome
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
// and removing all non-alphanumeric characters, it reads the same forward and backward.
// Given a string s, return true if it is a palindrome, or false otherwise.
func isPalindrome(s string) bool {
	cleaned := ""
	for _, char := range strings.ToLower(s) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned += string(char)
		}
	}

	left, right := 0, len(cleaned)-1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}
