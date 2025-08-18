package leetcode

// Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order
// of the non-zero elements. Note that you must do this in-place without making a copy of the array.
func moveZeroes(nums []int) {
	writeIndex := 0

	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != 0 {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	for writeIndex < len(nums) {
		nums[writeIndex] = 0
		writeIndex++
	}
}
