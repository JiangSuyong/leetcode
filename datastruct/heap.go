package datastruct

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Less(i, j int) bool {
	if h.Len() < i || h.Len() < j {
		return false
	}
	return h[i] < h[j]
}
