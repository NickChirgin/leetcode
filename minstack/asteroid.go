package minstack

func AsteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, asteroid := range asteroids {
	 for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
		 last := len(stack) - 1
		 diff := asteroid + stack[last]
		 if diff < 0 {
			 stack = stack[:last]
		 } else if diff > 0 {
			 asteroid = 0
		 } else {
			 asteroid = 0
			 stack = stack[:last]
		 }
	 }
	 if asteroid != 0 {
		 stack = append(stack, asteroid)
	 }
	}
	return stack
 }
