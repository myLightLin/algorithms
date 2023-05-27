package sorting

import "math"

/*
 * 基数排序的原理如下：
 * 1. 取得数组中的最大数，并取得位数；
 * 2. arr为原始数组，从最低位开始取每个位组成radix数组；
 * 3. 对radix进行计数排序（利用计数排序适用于小范围数的特点）；
 * 它是一个非原地，稳定排序。
 */
func radixSort(nums []int) {
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	for exp := 1; max >= exp; exp *= 10 {
		countingSortDigit(nums, exp)
	}
}

func countingSortDigit(nums []int, exp int) {
	counter := make([]int, 10)
	n := len(nums)
	for i := 0; i < n; i++ {
		d := digit(nums[i], exp)
		counter[d]++
	}
	for i := 1; i < 10; i++ {
		counter[i] += counter[i-1]
	}
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := digit(nums[i], exp)
		j := counter[d] - 1
		res[j] = nums[i]
		counter[d]--
	}
	for i := 0; i < n; i++ {
		nums[i] = res[i]
	}
}

func digit(num, exp int) int {
	return (num / exp) % 10
}
