package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func Cap(str1 string) string {
	newStr := ""
	for i, v := range str1 {
		if i == 0 && (v >= 97 && v <= 122) {
			v -= 32
			newStr += string(v)
		} else {
			newStr += string(v)
		}
	}
	return newStr
}

func Capitalize(strArr []string) string {
	var convertedArray []string
	for i, c := range strArr {
		if c == "(cap)" || c == "(cap)," || c == "(cap)." || c == "(cap)!" || c == "(cap)?" {
			strArr[i-1] = Cap(strArr[i-1])
		}
	}
	for _, v := range strArr {
		if v != "(cap)" && v != "(cap)," && v != "(cap)." && v != "(cap)!" && v != "(cap)?" {
			convertedArray = append(convertedArray, v)
		}
	}
	return (strings.Join(convertedArray, " "))
}

func IndexCap(strArr []string) string {
	var convertedArray []string
	next := ""

	for i, c := range strArr {
		if c == "(cap," {
			next = strArr[i+1]
			integer, _ := strconv.Atoi(next[:len(next)-1])

			if integer <= i {
				tocap := strArr[i-integer : i]
				for in, v := range tocap {
					tocap[in] = Cap(v)
				}
			} else {
				fmt.Printf("The format index %d is longer than the words to format %d\n", integer, i)
			}

		}
	}

	for _, v := range strArr {
		if v != "(cap," && v != next {
			convertedArray = append(convertedArray, v)
		}
	}
	return (strings.Join(convertedArray, " "))
}
