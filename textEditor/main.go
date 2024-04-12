package main

import (
	"fmt"
	"os"

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
	editedText := goreloaded.TextEditor(str)
	fmt.Println(editedText)
	
}
