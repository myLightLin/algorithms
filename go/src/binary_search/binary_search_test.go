package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var (
		target   = 3
		nums     = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected = 2
	)
	actual := binarySearch(nums, target)
	if actual != expected {
		t.Errorf("目标元素的索引是 %d, 期望是 %d", actual, expected)
	}

	actual2 := binarySearch1(nums, target)
	if actual2 != expected {
		t.Errorf("目标元素的索引是 %d, 期望是 %d", actual2, expected)
	}
}
