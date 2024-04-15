package goreloaded

import (
	"regexp"
	"strings"
	"unicode"
)
// Takes a string and Every instance of the punctuations ., ,, !, ?, : and ; should be
// close to the previous word and with space apart from the next one.
func Punctuation(str string) string {
	strArr := strings.Fields(str)
	punct := ""

	for i, c := range strArr {
		for _, v := range c {
			if unicode.IsPunct(v) {
				punct += string(v)
			}
			if string(c[0]) == punct && i != 0 {
				strArr[i-1] += punct
				strArr[i] = strings.Trim(strArr[i], punct)
				punct = ""
			}
		}

		punct = ""
	}
	return (strings.Join(strArr, " "))
}
// checks for a ' space or no space and any character not ^ and return a string with
// 'All char'
func Regex(str string) string {
	test := regexp.MustCompile(`'\s*([^']+)'`)
	str = test.ReplaceAllString(str, " '$1'")
	re := regexp.MustCompile(`\d+\)`)
	str = re.ReplaceAllString(str, "")

	return str
}
// checks for ( space word , space digit and ) and returns a string with space word
func Regex2(str string) string {
	re := regexp.MustCompile(`(\()(\s*)(\w+,\s+\d+\))`)
	str = re.ReplaceAllString(str, " $1$3 ")
	re = regexp.MustCompile(`(\()(\s+)(\w+\))`)
	str = re.ReplaceAllString(str, " $1$3 ")

	return str
}
