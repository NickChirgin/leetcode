package heaps

import "sort"

func lastStoneWeight(stones []int) int {
	for len(stones) != 1 {
			last := len(stones) - 1
			sort.Ints(stones)
			sub := stones[last] - stones[last-1]
			stones[last-1] = sub
			stones = stones[:last]
	}
	return stones[0]
}