package carFleet

import "sort"
func CarFleet(target int, position, speed []int) int {
	if len(speed) == 0 {
		return 1
	}
	posMap := map[int]int{}
	for i, val := range position {
		posMap[val] = speed[i]
	}
	sort.Ints(position)
	stack := []float32{}
	for i:= len(position)-1; i >=0; i-- {
		stack = append(stack, float32(target-position[i])/float32(posMap[position[i]]))
		if len(stack)>=2 && stack[len(stack)-1]	<= stack[len(stack)-2]{
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}