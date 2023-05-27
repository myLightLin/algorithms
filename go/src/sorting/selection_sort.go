package sorting

/**
 * 选择排序的原理如下：
 * 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
 * 然后，再从剩余未排序元素中继续寻找最小（大）元素，
 * 然后放到已排序序列的末尾。
 * 以此类推，直到所有元素均排序完毕。
 * 它是一个原地，非稳定排序。
 */
func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				swap(nums, i, minIndex)
			}
		}
	}
}

/**
 * @example
 * nums := []int{4, 1, 3, 1, 5, 2}
 * selectionSort(nums)
 * fmt.Println(nums)
 * @output
 * [1 1 2 3 4 5]
 */
