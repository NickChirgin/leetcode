package substring

import "fmt"

func MinWindow(s string, t string) string {
  tMap := map[rune]int{}
  left := 0
  substring := ""
	second := 0	
	for _, val := range t {
		tMap[val]++
	}
	copy := map[rune]int{}
	for k, v := range tMap {
		copy[k] = v
	}
	result := ""
	count := 0
	for r, v := range s {
		substring = s[left: r+1]
		fmt.Println(string(v), substring, len(copy))
		if _, ok := copy[v]; ok {
			copy[v]--
			count++
			if count == 1 {
				left = r 
			}
			if count == 2 {
				second = r
			}
			if copy[v] == 0 {
				delete(copy, v)
			}
		}
		if len(copy) == 0 {
			if result == "" {
				result = substring
			}
			for k, v := range tMap {
				copy[k] = v
			}
			copy[rune(s[second])]--
			if len(result) > len(s[left:r+1]) {
				result = substring
			}
			substring = ""
			left = second
			count=0
		}
	}
	return result
}