package leetcode

import (
	"sort"
)

// Group Anagrams
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Example 2:
// Input: strs = [""]
// Output: [[""]]
//
// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]
//
// Constraints:
// - 1 <= strs.length <= 10^4
// - 0 <= strs[i].length <= 100
// - strs[i] consists of lowercase English letters only.
func groupAnagrams(strs []string) [][]string {
	var keyMap = make(map[string][]string)
	for _, str := range strs {
		key := alphabetizeString(str)
		keyMap[key] = append(keyMap[key], str)
	}

	// assemble the
	var ans [][]string
	for _, keys := range keyMap {
		ans = append(ans, keys)
	}
	return ans
}

func alphabetizeString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}
