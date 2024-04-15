package goreloaded

import "strings"

// Takes a slice of string and returns a string with Every instance of a should be turned
// into an if the next word begins with a vowel 
func Vowel(strArr []string) string {
	for i, val := range strArr {
		if len(val) == 1 && val == "a" && i != len(strArr)-1 {
			next := strArr[i+1]
			if Isvowel(rune(next[0])) {
				strArr[i] = "an"
			}
		}
		if len(val) == 1 && val == "A" && i != len(strArr)-1 {
			next := strArr[i+1]
			if Isvowel(rune(next[0])) {
				strArr[i] = "An"
			}
		}

	}
	return (strings.Join(strArr, " "))
}

// Takes a rune and returns true or false if vowel or not
func Isvowel(c rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'h'}
	for _, v := range vowels {
		if v == c {
			return true
		}
	}
	return false
}
