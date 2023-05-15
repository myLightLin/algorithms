package binarysearch

func binarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m - 1
		} else {
			return m
		}
	}
	return -1
}

func binarySearch1(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m
		} else {
			return m
		}
	}
	return -1
}
