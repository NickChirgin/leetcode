package rpn

import "strconv"
func EvalRPN(tokens []string) int {
	operations := map[string]func(x, y int)int{
		"+": func(x, y int) int {return x+y},
		"-": func(x, y int) int {return x-y},
		"*": func(x, y int) int {return x*y},
		"/": func(x, y int) int {return x/y}}
	stack := []int{}
	for _, val := range tokens {
		if _, ok := operations[val]; ok {
			last := len(stack)-1
			result := operations[val](stack[last-1], stack[last])
			stack = stack[:last]
			stack[last-1]= result
		}else {
			intval, _ := strconv.Atoi(val)
			stack = append(stack, intval)
		}
	}
	return stack[0]
}