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

func duplicateZeros(arr []int) {
	zeros := 0
	lastIndex := len(arr) - 1
	for left := 0; left <= (lastIndex - zeros); left++ {
		if arr[left] == 0 {
			if left == lastIndex-zeros {
				arr[lastIndex] = 0
				lastIndex--
				break
			}
			zeros++
		}
	}
	for i := lastIndex - zeros; i >= 0; i-- {
		arr[i+zeros] = arr[i]
		if arr[i] == 0 {
			zeros--
			arr[i+zeros] = 0
		}
	}
}

func duplicateZerosV1(arr []int) {
	zeros := 0
	for _, v := range arr {
		if v == 0 {
			zeros += 1
		}
	}
	length := len(arr)
	if zeros == 0 || zeros == length {
		return
	}

	oldEnd := length - 1
	newEnd := oldEnd + zeros

	for newEnd >= 0 {
		if newEnd < length {
			arr[newEnd] = arr[oldEnd]
		}
		if arr[oldEnd] == 0 {
			newEnd--
			if newEnd < length {
				arr[newEnd] = 0
			}
		}
		newEnd--
		oldEnd--
	}
}
func countZeros(arr []int) int {
	zeros := 0
	for i := 0; i < len(arr) && (i+zeros) < len(arr); i++ {
		if arr[i] == 0 {
			zeros++
		}
	}
	return zeros
}
