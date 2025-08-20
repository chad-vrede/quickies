package leetcode

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "base case n=0",
			input:    0,
			expected: 0,
		},
		{
			name:     "base case n=1",
			input:    1,
			expected: 1,
		},
		{
			name:     "small case n=2",
			input:    2,
			expected: 1,
		},
		{
			name:     "small case n=3",
			input:    3,
			expected: 2,
		},
		{
			name:     "small case n=4",
			input:    4,
			expected: 3,
		},
		{
			name:     "medium case n=5",
			input:    5,
			expected: 5,
		},
		{
			name:     "medium case n=6",
			input:    6,
			expected: 8,
		},
		{
			name:     "larger case n=10",
			input:    10,
			expected: 55,
		},
		{
			name:     "constraint boundary n=30",
			input:    30,
			expected: 832040,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fib(tt.input)
			if result != tt.expected {
				t.Errorf("fib(%d) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkFib(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib(30)
	}
}
