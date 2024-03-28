package leet

import (
	"sort"
)

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	for _, v := range nums {
		if v == 1 {
			count += 1
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}
	if count > maxCount {
		maxCount = count
	}
	return maxCount
}

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		scale := 0
		for num > 0 {
			num /= 10
			scale++
		}
		if scale%2 == 0 {
			count++
		}
	}
	return count
}
