package leetcode

// First Unique Character in a String
// Given a string s, find the first non-repeating character in it and return its index.
// If it does not exist, return -1.
func firstUniqChar(s string) int {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}

	return -1
}
