package validPalyndrome

import (
	"fmt"
	"regexp"
	"strings"
)
func ValidPalidrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
	last := len(s)-1
	first := 0
	fmt.Println(s)
	for first <= last {
			if s[first] != s[last]{

				return false
			}
			first++
			last--
	}
	return true
}