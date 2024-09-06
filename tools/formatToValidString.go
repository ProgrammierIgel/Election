package tools

import (
	"strings"
)

func FormatToValidString(str string) string {
	newStr := ""
	for iteration, letter := range string(str) {
		if iteration == 0 {
			newStr = newStr + strings.ToUpper(string(letter))
			continue
		}
		newStr = newStr + string(letter)
	}
	return newStr
}
