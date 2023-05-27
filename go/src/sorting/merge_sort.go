package sorting

/**
 * 归并排序是一个非原地，稳定排序。它的原理如下：
 * 1. 把长度为n的输入序列分成两个长度为n/2的子序列；
 * 2. 对这两个子序列分别采用归并排序；
 * 3. 将两个排序好的子序列合并成一个最终的排序序列。
 * 归并排序 merge 函数的逻辑：
 * 1. 申请两个临时数组，分别存放左右两个子序列；
 * 2. 从左右两个子序列的起始位置开始比较，将较小的元素放入原数组中；
 * 3. 如果左子序列已经全部放入原数组中，则将右子序列剩余元素放入原数组中；
 * 4. 如果右子序列已经全部放入原数组中，则将左子序列剩余元素放入原数组中；
 * 5. 重复步骤2~4。
 */
func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i := left
	j := mid + 1
	k := 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	copy(nums[left:right+1], tmp)
}

/**
 * @example
 * nums := []int{7, 3, 2, 6, 0, 1, 5, 4}
 * mergeSort(nums, 0, len(nums) - 1)
 * fmt.Println(nums)
 * @output
 * [1 1 2 3 4 5]
 */
