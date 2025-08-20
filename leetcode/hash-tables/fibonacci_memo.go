package leetcode

// Fibonacci Number with Memoization
// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence,
// such that each number is the sum of the two preceding ones, starting from 0 and 1.
//
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
//
// Given n, calculate F(n) using memoization to avoid recalculating the same values.
//
// Example 1:
// Input: n = 2
// Output: 1
// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
//
// Example 2:
// Input: n = 3
// Output: 2
// Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
//
// Example 3:
// Input: n = 4
// Output: 3
// Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
//
// Constraints:
// - 0 <= n <= 30
func fib(n int) int {
	fibCache := make(map[int]int)
	return fibHelper(n, fibCache)
}

func fibHelper(n int, cache map[int]int) int {
	if n == 1 {
		return 1
	}
	if n <= 0 {
		return 0
	}

	if v, ok := cache[n]; ok {
		return v
	}

	cache[n] = fibHelper(n-1, cache) + fibHelper(n-2, cache)
	return cache[n]
}
