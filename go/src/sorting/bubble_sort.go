package sorting

/**
 * 冒泡排序的原理如下：
 * 1. 比较相邻的元素。如果第一个比第二个大，就交换它们两个；
 * 2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。
 *    这步做完后，最后的元素会是最大的数；
 * 3. 针对所有的元素重复以上的步骤，除了最后一个；
 * 4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
 * 它是一个原地，稳定排序。
 */

func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}

/**
 * @example
 * nums := []int{4, 1, 3, 1, 5, 2}
 * bubbleSort(nums)
 * fmt.Println(nums)
 * @output
 * [1 1 2 3 4 5]
 */
