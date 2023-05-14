package minstack

import "math"

func AsteroidCollision(asteroids []int) []int {
 stack := []int{}
 for _, asteroid := range asteroids {
	last := len(stack) - 1
	if len(stack) == 0  ||  !isOpposite(stack[last], asteroid){
		stack = append(stack, asteroid)
		continue
	}
	if stack[last] == asteroid * -1 {
		stack = stack[:last]
		continue
	}
	temp := asteroid
	for last >= 0 && isOpposite(stack[last], temp) {
		temp = max(stack[last], temp)
		stack[last] = temp
		last--
	}
	stack = stack[:last+1]
	stack = append(stack, temp)
 }
 return stack
}



func isOpposite(a, b int) bool {
	return (a > 0 && b < 0) || (b > 0 && a < 0)
}

func max(a, b int) int {
	if math.Abs(float64(a)) > math.Abs(float64(b)) {
		return a
	}
	return b
}