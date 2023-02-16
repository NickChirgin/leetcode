package maxArea

import (
	"fmt"
)

	func MaxArea(height []int) int {
		first := 0
		area := 0
		last := len(height) - 1
		for first < last {
			fmt.Println(height[first], height[last])
			if (last - first) * min(height[first], height[last]) > area {
				area = (last - first) * min(height[first], height[last]) 
			}
			if height[first] > height[last]{
				last--
			} else {
				first++
			}
		}
		return area
	}

	func min(x,y int) int {
		if x < y {
			return x
		}
		return y
	}