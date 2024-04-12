package goreloaded

import (
	"strconv"
	"strings"
)

func Upper(strArr []string) string {
	var convertedArray []string
	for i, c := range strArr {
		if c == "(up)" || c == "(up)," || c == "(up)." || c == "(up)!" || c == "(up)?" {
			strArr[i-1] = strings.ToUpper(strArr[i-1])
		}
	}

	for _, v := range strArr {
		if v != "(up)" && v != "(up)," && v != "(up)." && v != "(up)!" && v != "(up)?" {
			convertedArray = append(convertedArray, v)
		}
	}

	return (strings.Join(convertedArray, " "))
}

func IndexUpper(strArr []string) string {
	var convertedArray []string
	next := ""

	for i, c := range strArr {
		if c == "(up," {
			next = strArr[i+1]
			integer, _ := strconv.Atoi(next[:len(next)-1])
			tocap := strArr[i-integer : i]

			for in, v := range tocap {
				tocap[in] = strings.ToUpper(v)
			}

		}
	}

	for _, v := range strArr {
		if v != "(up," && v != next {
			convertedArray = append(convertedArray, v)
		}
	}
	return (strings.Join(convertedArray, " "))
}
