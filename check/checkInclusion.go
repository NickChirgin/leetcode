package check

import "fmt"

func CheckInclusion(s1, s2 string) bool {
	sMap := map[rune]int{}
	for _, val := range s1 {
		sMap[val]++
	}
	substring := ""
	left := 0
	for i:= len(s1); i<= len(s2); i++ {
		substring = s2[left:i]	
		copyMap := map[rune]int{}
		for k, v := range sMap {
			copyMap[k] = v
		}
		for _, val := range substring {
			if _, ok := copyMap[val]; !ok {
				break
			}
			fmt.Println(substring, copyMap)
			copyMap[val]--
			if copyMap[val] == 0 {
				delete(copyMap, val)
			}
		}
	if len(copyMap) == 0 {
		return true
	}
	left++
	}
	return false
}