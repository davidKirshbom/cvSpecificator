package utils

import "strings"

func Titleize(str ...string) string {
	var result string
	for _, s := range str {
		result += strings.ToTitle(s[0:1]) + strings.ToLower(s[1:]) + " "
		
	}
	return result
}
func NewLineToBr(str string) string {
	return strings.Replace(str, "\n", "<br>", -1)
}