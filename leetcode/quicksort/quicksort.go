package quicksort

func QuickSort(nums []int) []int {

	quicksort(nums, 0, len(nums)-1)

	return nums
}

func quicksort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quicksort(nums, 0, pivot-1)
		quicksort(nums, pivot+1, end)
	}

}

func partition(nums []int, start int, end int) int {
	//选取最后一个元素作为基准pivot
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap1(nums, i, j)
			i++
		}
	}
	swap1(nums, i, end)
	return i
}

func swap1(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
