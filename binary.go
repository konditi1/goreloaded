package goreloaded

import (
	"strconv"
	"strings"
)
// Takes a slice of string and and returns a string with word after bin
//converted to decimal
func Binary(strArr []string) string {
	var convertedArray []string
	for i, c := range strArr {
		if c == "(bin)" || c == "(bin)," || c == "(bin)." || c == "(bin)!" || c == "(bin)?" {
			outPut, _ := strconv.ParseInt(strArr[i-1], 2, 64)
			strArr[i-1] = strconv.FormatInt(outPut, 10)
		}
	}
	for _, v := range strArr {
		if v != "(bin)" && v != "(bin)," && v != "(bin)." && v != "(bin)!" && v != "(bin)?" {
			convertedArray = append(convertedArray, v)
		}
	}
	return (strings.Join(convertedArray, " "))
}
