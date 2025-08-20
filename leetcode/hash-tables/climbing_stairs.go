package leetcode

// Climbing Stairs
// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Example 1:
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
//
// Example 2:
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
//
// Example 3:
// Input: n = 4
// Output: 5
// Explanation: There are five ways to climb to the top.
// 1. 1+1+1+1
// 2. 2+1+1
// 3. 1+2+1
// 4. 1+1+2
// 5. 2+2
//
// Constraints:
// - 1 <= n <= 45
func climbStairs(n int) int {
	cache := make(map[int]int)
	return climbHelper(n, cache)
}

func climbHelper(n int, cache map[int]int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if v, ok := cache[n]; ok {
		return v
	}

	cache[n] = climbHelper(n-1, cache) + climbHelper(n-2, cache)
	return cache[n]
}
