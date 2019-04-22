package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	results := twoSum([]int{2, 7, 11, 15}, 9)

	assert.Equal(t, results, []int{0, 1})
}

func TestTwoSum2(t *testing.T) {
	results := twoSum([]int{3, 2, 3}, 6)

	assert.Equal(t, []int{0, 2}, results)
}

func TestTwoSum3(t *testing.T) {
	results := twoSum([]int{3, 2, 4}, 6)

	assert.Equal(t, []int{1, 2}, results)
}
