package twosum

import "testing"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 5, 8, 10, 3, 121, 328, 2, 7, 34, 78}
	target := 18
	t.Log(twoSum(nums, target))
}
