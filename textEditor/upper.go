package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// Takes a slice of string and returns a string with the letter before (low) to upperCase
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

// Takes a slice of string and returns a string previously specified number of words in upperCase
func IndexUpper(strArr []string) string {
	var convertedArray []string
	next := ""

	for i, c := range strArr {
		if c == "(up," {
			next = strArr[i+1]
			integer, _ := strconv.Atoi(next[:len(next)-1])
			
			if integer <= i {
				tocap := strArr[i-integer : i]	
				for in, v := range tocap {
					tocap[in] = strings.ToUpper(v)
				}
			} else {
				fmt.Printf("The format index %d is longer than the words to format %d\n", integer, i)
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
