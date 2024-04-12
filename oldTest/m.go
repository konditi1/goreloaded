package m

import (
	"fmt"
	"os"
	"strings"
)

func m() {
	if len(os.Args) != 3 {
		fmt.Println("Error! The number of arguments is not 3")
		return
	}
	fileToRead := os.Args[1]
	writeToFile := os.Args[2]

	file, err := os.ReadFile(fileToRead)
	if err != nil {
		fmt.Println("Error! Could not read file", err)
	}

	var arr []string
	str := ""

	for _, v := range file {
		if v != ' ' {
			str += string(v)
		} else {
			arr = append(arr, str)
			str = ""
		}
	}
	if str != "" {
		arr = append(arr, str)
	}

	for i, c := range arr {
		if c == "(up)" || c == "(up)," || c == "(up)." || c == "(up)!" || c == "(up)?" {
			arr[i-1] = strings.ToUpper(arr[i-1])
			if i != len(arr) {
				arr = append(arr[:i], arr[i+1:]...)
			}
		} else if c == "(low)" || c == "(low)," || c == "(low)." || c == "(low)!" || c == "(low)?" {
			arr[i-1] = strings.ToLower(arr[i-1])
			// arr = append(arr[:i], arr[i+1:]...)
		}
	}

	str2 := ""
	for i, v := range arr {
		if i != len(arr)-1 {
			str2 += v + " "
		} else {
			str2 += v
		}
	}

	file2 := []byte(str2)
	err = os.WriteFile(writeToFile, file2, 0o744)
	if err != nil {
		println("failed to write", err)
	} else {
		fmt.Println("successfully written")
	}
}
