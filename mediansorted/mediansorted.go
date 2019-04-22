package mediansorted

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// run the count sort first
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	return combineSort(nums1, nums2)
}

func findHigh(nums1 []int, nums2 []int) int {
	// find highest value
	var high int

	// this might look messy, but it's faster than iterating the array
	if len(nums1) > 0 && len(nums2) > 0 {
		if nums1[len(nums1)-1] > nums2[len(nums2)-1] {
			high = nums1[len(nums1)-1]
		} else {
			high = nums2[len(nums2)-1]
		}
	}

	if len(nums1) > 0 && len(nums2) == 0 {
		high = nums1[len(nums1)-1]
	}

	if len(nums2) > 0 && len(nums1) == 0 {
		high = nums2[len(nums2)-1]
	}

	return high
}

func combineSort(nums1, nums2 []int) float64 {
	high := findHigh(nums1, nums2)

	indexArray := make([]int, ((high * 2) + 1))

	for _, num := range nums1 {
		if num < 0 {
			indexArray[int(math.Abs(float64(num)))] = indexArray[int(math.Abs(float64(num)))] + 1
			continue
		}

		indexArray[high+num] = indexArray[high+num] + 1
	}

	for _, num := range nums2 {
		if num < 0 {
			indexArray[int(math.Abs(float64(num)))] = indexArray[int(math.Abs(float64(num)))] + 1
			continue
		}

		indexArray[high+num] = indexArray[high+num] + 1
	}

	var results []int
	// fill negative numbers first,
	for i := high - 1; i > 0; i-- {
		if indexArray[i] <= 0 {
			continue
		}

		results = append(results, i*-1)
	}

	for i := high; i < high*2+1; i++ {
		if indexArray[i] <= 0 {
			continue
		}

		for j := 0; j < indexArray[i]; j++ {
			results = append(results, int(i-high))
		}

	}

	if len(results)%2 != 0 {
		return float64(results[int(len(results)/2)])
	}

	return float64((results[int(len(results)/2)] + results[int((len(results)/2)-1)])) / 2.0
}
