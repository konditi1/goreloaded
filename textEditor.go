package goreloaded

import (
	"strings"
	"unicode"
)
// Calls all the function to edit the passed string and returns the edited string
func TextEditor(str string) string {
	strArr := strings.Split(string(str), " ")

	str = Vowel(strArr)

	for _, v := range str {
		if unicode.IsPunct(v) {
			str = Punctuation(str)
		}
	}

	str = Regex2(str)
	strArr = strings.Fields(str)

	for _, v := range strArr {
		if v == "(up)" || v == "(up)," || v == "(up)." || v == "(up)!" || v == "(up)?" {
			str = Upper(strArr)
			strArr = strings.Fields(str)
		} else if v == "(up," {
			str = IndexUpper(strArr)
			strArr = strings.Fields(str)
		} else if v == "(low)" || v == "(low)," || v == "(low)." || v == "(low)!" || v == "(low)?" {
			str = Lower(strArr)
			strArr = strings.Fields(str)
		} else if v == "(low," {
			str = IndexLow(strArr)
			strArr = strings.Fields(str)
		} else if v == "(cap)" || v == "(cap)," || v == "(cap)." || v == "(cap)!" || v == "(cap)?" {
			str = Capitalize(strArr)
			strArr = strings.Fields(str)
		} else if v == "(cap," {
			str = IndexCap(strArr)
			strArr = strings.Fields(str)
		} else if v == "(hex)" || v == "(hex)," || v == "(hex)." || v == "(hex)!" || v == "(hex)?" {
			str = HexaDeci(strArr)
			strArr = strings.Fields(str)
		} else if v == "(bin)" || v == "(bin)," || v == "(bin)." || v == "(bin)!" || v == "(bin)?" {
			str = Binary(strArr)
			strArr = strings.Fields(str)
		}
	}

	for _, v := range str {
		if unicode.IsPunct(v) {
			str = Punctuation(str)
		}
	}

	regexp := Regex(str)
	regexp1 := strings.Fields(regexp)
	regexp = strings.Join(regexp1, " ")

	return regexp
}
