package heaps

/*
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	a, b := h[i], h[j]
	x1,y1 := a[0], a[1]
	x2,y2 := b[0], b[1]
	resI := math.Sqrt(float64(x1 * x1 + y1 * y1))
	resJ := math.Sqrt(float64(x2 * x2 + y2 * y2))
	return resI < resJ
}
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {*h = append(*h, x.([]int))}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
    temp := IntHeap(points)
		heap.Init(&temp)

		for len(temp) > k {
			temp.Pop()
		}
		return temp
}
*/