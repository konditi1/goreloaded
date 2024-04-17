package goreloaded

import (
	"strconv"
	"strings"
)

func HexaDeci(strArr []string) string {
	var convertedArray []string
	for i, c := range strArr {
		if c == "(hex)" || c == "(hex)," || c == "(hex)." || c == "(hex)!" || c == "(hex)?" {
			outPut, _ := strconv.ParseInt(strArr[i-1], 16, 64)
			strArr[i-1] = strconv.FormatInt(outPut, 10)
		}
	}
	for _, v := range strArr {
		if v != "(hex)" && v != "(hex)," && v != "(hex)." && v != "(hex)!" && v != "(hex)?" {
			convertedArray = append(convertedArray, v)
		}
	}
	return (strings.Join(convertedArray, " "))
}
