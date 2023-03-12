package heaps

import (
	"container/heap"
	"fmt"
	"sort"
)
type ArrHeap [][]int

func (h ArrHeap) Len() int           { return len(h) }
func (h ArrHeap) Less(i, j int) bool {
	return	h[i][0] < h[j][0]
}
func (h ArrHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ArrHeap) Push(x interface{}) {*h = append(*h, x.([]int))}
func (h *ArrHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func GetOrder(tasks [][]int) []int {
	for i := range tasks {
		tasks[i] = append(tasks[i], i)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})
	var minHeap ArrHeap
	res := []int{} 
	i, time := 0, tasks[0][0]
	heap.Init(&minHeap)
	for minHeap.Len() > 0 || i < len(tasks) {
		for i < len(tasks) && time >= tasks[i][0] {
			heap.Push(&minHeap, tasks[i][1:])
			i++
		}
		if minHeap.Len() == 0 && i == len(tasks) {
			break
		}
		if minHeap.Len() > 0 {
			time = tasks[i][0]
		} else {
			pop := heap.Pop(&minHeap).([]int)
			procTime, index := pop[0], pop[1]
			time += procTime
			res = append(res, index)
		}
	}
	fmt.Println(tasks)
	return res 
}