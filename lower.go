package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func Lower(strArr []string) string {
	var convertedArray []string
	for i, c := range strArr {
		if c == "(low)" || c == "(low)," || c == "(low)." || c == "(low)!" || c == "(low)?" {
			strArr[i-1] = strings.ToLower(strArr[i-1])
		}
	}

	for _, c := range strArr {
		if c != "(low)" && c != "(low)," && c != "(low)." && c != "(low)!" && c != "(low)?" {
			convertedArray = append(convertedArray, c)
		}
	}

	return (strings.Join(convertedArray, " "))
}

func IndexLow(strArr []string) string {
	var convertedArray []string
	next := ""

	for i, c := range strArr {
		if c == "(low," {
			next = strArr[i+1]
			integer, _ := strconv.Atoi(next[:len(next)-1])

			if integer < i {
				tolow := strArr[i-integer : i]
				for in, v := range tolow {
					tolow[in] = strings.ToLower(v)
				}
			} else {
				fmt.Printf("The format index %d is longer than the words to format %d\n", integer, i)
			}

		}
	}

	for _, v := range strArr {
		if v != "(low," && v != next {
			convertedArray = append(convertedArray, v)
		}
	}
	return (strings.Join(convertedArray, " "))
}
