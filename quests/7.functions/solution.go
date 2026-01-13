package functions

import (
	"fmt"
	"math"
	"strings"
)

// Divide returns a / b or an error if b == 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("b cannot be zero")
	}

	return a / b, nil
}

// SumAll returns the sum of all provided integers
func SumAll(nums ...int) int {
	res := 0
	for _, number := range nums {
		res += number
	}

	return res
}

// MaxMin returns the max and min of all provided integers
// Returns an error if no numbers are provided
func MaxMin(nums ...int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, fmt.Errorf("nums is empty")
	}

	min := math.MaxInt64
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min, nil
}

// ConcatAll joins all strings using the provided separator
func ConcatAll(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}
