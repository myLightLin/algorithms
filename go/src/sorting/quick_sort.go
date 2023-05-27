package sorting

/**
 * 快速排序的原理如下：
 * 1. 从数列中挑出一个元素，称为 “基准”（pivot）；
 * 2. 重新排序数列，所有元素比基准值小的摆放在基准前面，
 *    所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
 *    在这个分割之后，该基准是它的最后位置。这个称为分割（partition）操作；
 * 3. 递归地把小于基准值元素的子数列和大于基准值元素的子数列排序。
 * 它是一个原地，非稳定排序。
 */
func quickSort(nums []int) {
	quickSortC(nums, 0, len(nums)-1)
}

func quickSortC(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(nums, start, end)
	quickSortC(nums, start, pivot-1)
	quickSortC(nums, pivot+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

/**
 * @example
 * nums := []int{4, 1, 3, 1, 5, 2}
 * quickSort(nums)
 * fmt.Println(nums)
 * @output
 * [1 1 2 3 4 5]
 */
