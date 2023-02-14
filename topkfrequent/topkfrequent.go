package topkfrequent

import (
	"sort"
)

type Pair struct {
	Key int
	Value int
}
func TopKFrequent(nums []int, k int) []int {
	res := map[int]int{}
	for _, v := range nums {
		res[v]++
	}
	keys := make([]Pair, len(res))
	i := 0
	for k, v:= range res {
		keys[i] = Pair{k, v}
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Value > keys[j].Value
	})
	i = 0
	result := make([]int, k)
	for i <= k-1 {
		result[i] = keys[i].Key
		i++
	}
	return result
}