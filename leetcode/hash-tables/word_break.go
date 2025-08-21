package leetcode

// Word Break
// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into
// a space-separated sequence of one or more dictionary words.
//
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
//
// Example 1:
// Input: s = "leetcode", wordDict = ["leet","code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".
//
// Example 2:
// Input: s = "applepenapple", wordDict = ["apple","pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
// Note that you are allowed to reuse a dictionary word.
//
// Example 3:
// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// Output: false
// Explanation: Cannot be segmented into dictionary words.
//
// Constraints:
// - 1 <= s.length <= 300
// - 1 <= wordDict.length <= 1000
// - 1 <= wordDict[i].length <= 20
// - s and wordDict[i] consist of only lowercase English letters.
// - All the strings of wordDict are unique.
func wordBreak(s string, wordDict []string) bool {
	cache := make(map[int]bool)
	return canBreak(s, 0, wordDict, cache)
}

func canBreak(s string, index int, wordDict []string, cache map[int]bool) bool {
	// base case
	if index == len(s) {
		return true
	}

	// check cache
	if val, ok := cache[index]; ok {
		return val
	}

	// try each word in dictionary
	for _, word := range wordDict {
		// check if current position matches this word
		if index+len(word) <= len(s) && s[index:index+len(word)] == word {
			// recursively check if rest can be broken
			if canBreak(s, index+len(word), wordDict, cache) {
				cache[index] = true
				return true
			}
		}
	}

	// no word worked, cache false result
	cache[index] = false
	return false
}
