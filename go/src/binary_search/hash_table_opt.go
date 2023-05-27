package binarysearch

// 给定一个整数数组 nums 和一个目标元素 target ，
// 请在数组中搜索“和”为 target 的两个元素，并返回它们的数组索引。
// 返回任意一个解即可。
func twoSumBruteForce(nums []int, target int) []int {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumHashTable(nums []int, target int) []int {
	hashTable := map[int]int{}
	for idx, val := range nums {
		if preIdx, ok := hashTable[target-val]; ok {
			return []int{preIdx, idx}
		}
		hashTable[val] = idx
	}
	return nil
}
