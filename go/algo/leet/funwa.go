package leet

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

func heightChecker(heights []int) int {
	return 0
}
