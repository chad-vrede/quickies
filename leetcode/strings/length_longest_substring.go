package leetcode

// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	left := 0
	maxLength := 0

	for right := 0; right < len(s); right++ {
		if lastIndex, found := charMap[s[right]]; found && lastIndex >= left {
			left = lastIndex + 1
		}

		charMap[s[right]] = right
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
