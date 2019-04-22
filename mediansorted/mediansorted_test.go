package mediansorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArray(t *testing.T) {
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func TestFindMedianSortedArray2(t *testing.T) {
	assert.Equal(t, 1.0, findMedianSortedArrays([]int{1, 1}, []int{1, 2}))
}

func TestFindMedianSortedArrayEven(t *testing.T) {
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func TestFindMedianSortedArrayNegative(t *testing.T) {
	assert.Equal(t, -1.0, findMedianSortedArrays([]int{3}, []int{-2, -1}))
}

func TestFindMedianSortedArrayNegative4(t *testing.T) {
	assert.Equal(t, -1.0, findMedianSortedArrays([]int{3}, []int{-2, -1}))
}
