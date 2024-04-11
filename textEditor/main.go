package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"goreloaded"
)

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Println("Error The Number of Args must be 3")
	// }
	sourceFile := os.Args[1]
	// destFile := os.Args[2]

	file, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Error! Could not read file", err)
	}
	str := string(file)
	strArr := strings.Split(string(str), " ")

	for _, v := range strArr {
		if v == "(up)" || v == "(up)," || v == "(up)." || v == "(up)!" || v == "(up)?" {
			str = goreloaded.Upper(strArr)
			strArr = strings.Fields(str)
		} else if v == "(up," {
			str = goreloaded.IndexUpper(strArr)
			strArr = strings.Fields(str)
		} else if v == "(low)" || v == "(low)," || v == "(low)." || v == "(low)!" || v == "(low)?" {
			str = goreloaded.Lower(strArr)
			strArr = strings.Fields(str)
		} else if v == "(low," {
			str = goreloaded.IndexLow(strArr)
			strArr = strings.Fields(str)
		} else if v == "(cap)" || v == "(cap)," || v == "(cap)." || v == "(cap)!" || v == "(cap)?" {
			str = goreloaded.Capitalize(strArr)
			strArr = strings.Fields(str)
		} else if v == "(cap," {
			str = goreloaded.IndexCap(strArr)
			strArr = strings.Fields(str)
		} else if v == "(hex)" || v == "(hex)," || v == "(hex)." || v == "(hex)!" || v == "(hex)?" {
			str = goreloaded.HexaDeci(strArr)
			strArr = strings.Fields(str)
		} else if v == "(bin)" || v == "(bin)," || v == "(bin)." || v == "(bin)!" || v == "(bin)?" {
			str = goreloaded.Binary(strArr)
			strArr = strings.Fields(str)
		}
	}
	for _, v := range str {
		if unicode.IsPunct(v) {
			str = goreloaded.Punctuation(str)
		}
	}

	regexp := goreloaded.Regex(str)
	fmt.Println(regexp)
}
