package temperature

func DailyTemperatures(temperatures []int) []int {
   stack := [][]int{}
	 answer := make([]int, len(temperatures))
	 for i, temp := range temperatures {
		if len(stack) == 0 || temp <= stack[len(stack)-1][0] {
			stack = append(stack, []int{temp, i})
		}else {
			count := 0
			for _, v := range stack {
				if temp > v[0] {
					answer[v[1]] =i- v[1]
					count++	
				}
			}
			stack = stack[:len(stack)-count]
			stack = append(stack, []int{temp, i})
		}
	 }
	return answer
}