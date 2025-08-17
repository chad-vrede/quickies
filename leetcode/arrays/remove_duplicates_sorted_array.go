package leetcode

// Remove Duplicates from Sorted Array
// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place 
// such that each unique element appears only once. The relative order of the elements should be kept the same.
// Return k after placing the final result in the first k slots of nums.
func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    writeIndex := 1
    
    for readIndex := 1; readIndex < len(nums); readIndex++ {
        if nums[readIndex] != nums[readIndex-1] {
            nums[writeIndex] = nums[readIndex]
            writeIndex++
        }
    }
    
    return writeIndex
}