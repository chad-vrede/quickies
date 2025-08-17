package leetcode

// Valid Anagram
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
// typically using all the original letters exactly once.
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    charCount := make(map[rune]int)
    
    for _, char := range s {
        charCount[char]++
    }
    
    for _, char := range t {
        charCount[char]--
        if charCount[char] < 0 {
            return false
        }
    }
    
    return true
}