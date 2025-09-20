package leetcode

// Top K Frequent Elements
// Given an integer array nums and an integer k, return the k most frequent elements.
// You may return the answer in any order.
//
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
//
// Example 2:
// Input: nums = [1], k = 1
// Output: [1]
//
// Constraints:
// - 1 <= nums.length <= 10^5
// - k is in the range [1, the number of unique elements in the array].
// - It is guaranteed that the answer is unique.
//
// Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
func topKFrequent(nums []int, k int) []int {
	var res []int
	frequencies := make(map[int]int, 0)
	buckets := make([][]int, len(nums)+1)

	for _, num := range nums {
		frequencies[num]++
	}

	for element, frequency := range frequencies {
		buckets[frequency] = append(buckets[frequency], element)
	}

	for frequency := len(buckets) - 1; frequency >= 0 && k > 0; frequency-- {
		for _, element := range buckets[frequency] {
			if k > 0 {
				res = append(res, element)
				k--
			}
		}
	}

	return res
}
