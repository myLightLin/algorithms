package sorting

/**
 * 插入排序的原理如下：
 * 1. 从第一个元素开始，该元素可以认为已经被排序；
 * 2. 取出下一个元素，在已经排序的元素序列中从后向前扫描；
 * 3. 如果该元素（已排序）大于新元素，将该元素移到下一位置；
 * 4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
 * 5. 将新元素插入到该位置后；
 * 6. 重复步骤2~5。
 * 它是一个原地，稳定排序。
 */

func insertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}

/**
 * @example
 * nums := []int{4, 1, 3, 1, 5, 2}
 * insertionSort(nums)
 * fmt.Println(nums)
 * @output
 * [1 1 2 3 4 5]
 */
