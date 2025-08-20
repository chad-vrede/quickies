package leetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "base case n=1",
			input:    1,
			expected: 1,
		},
		{
			name:     "base case n=2",
			input:    2,
			expected: 2,
		},
		{
			name:     "small case n=3",
			input:    3,
			expected: 3,
		},
		{
			name:     "small case n=4",
			input:    4,
			expected: 5,
		},
		{
			name:     "medium case n=5",
			input:    5,
			expected: 8,
		},
		{
			name:     "medium case n=6",
			input:    6,
			expected: 13,
		},
		{
			name:     "larger case n=10",
			input:    10,
			expected: 89,
		},
		{
			name:     "constraint boundary n=45",
			input:    45,
			expected: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := climbStairs(tt.input)
			if result != tt.expected {
				t.Errorf("climbStairs(%d) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkClimbStairs(t *testing.B) {
	for i := 0; i < t.N; i++ {
		climbStairs(45)
	}
}
