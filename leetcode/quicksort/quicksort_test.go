package quicksort

import "testing"

func TestQuickSort(t *testing.T) {
	nums := []int{1, 23, 4, 56, 23, 4, 6, 46, 13, 4, 5, 20, 21, 24, 45, 8, 17}
	sortNums := QuickSort(nums)
	t.Log(sortNums)
}
