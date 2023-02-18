package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/carFleet"
)
func main() {
	nums := []int{10,8,0,5,3}
	speed := []int{2,4,1,1,3}
	fmt.Println(carFleet.CarFleet(12, nums, speed))
}