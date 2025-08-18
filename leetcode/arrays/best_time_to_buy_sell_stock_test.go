package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "basic profit case",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "no profit possible",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "single price",
			prices:   []int{1},
			expected: 0,
		},
		{
			name:     "two prices - profit",
			prices:   []int{1, 5},
			expected: 4,
		},
		{
			name:     "two prices - no profit",
			prices:   []int{5, 1},
			expected: 0,
		},
		{
			name:     "all same prices",
			prices:   []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "increasing prices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "buy low sell high with fluctuation",
			prices:   []int{2, 4, 1, 7, 5, 1, 3},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.prices)
			if result != tt.expected {
				t.Errorf("maxProfit(%v) = %v, expected %v", tt.prices, result, tt.expected)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4, 2, 8, 1, 9, 3, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(prices)
	}
}
