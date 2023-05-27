package sorting

/**
 * 计数排序的原理如下：
 * 找到数组中最大的数字 m，然后新建一个长度为 m + 1 的数组
 * 接着遍历数组，以数字为索引，统计每个数字出现的次数作为数组值
 */
func countingSortNaive(nums []int) {
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}

	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}

	for i, num := 0, 0; num < m+1; num++ {
		for j := 0; j < counter[num]; j++ {
			nums[i] = num
			i++
		}
	}
}

func countingSort(nums []int) {
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}

	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 求 counter 的前缀和
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}

	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		res[counter[num]-1] = num
		counter[num]--
	}
	copy(nums, res)
}
