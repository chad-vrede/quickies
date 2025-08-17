package leetcode

// Contains Duplicate
// Given an integer array nums, return true if any value appears at least twice in the array, 
// and return false if every element is distinct.
func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    
    return false
}