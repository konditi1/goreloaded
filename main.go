package main

import (
	"fmt"
	"os"

	goreloaded "goreloaded/textEditor"
)

func main() {
	// check if the user entered three arguments
	if len(os.Args) != 3 {
		fmt.Println("Error The Number of Args must be 3")
	}
	// Assign the arguments one and two
	sourceFile := os.Args[1]
	destFile := os.Args[2]

	// Read from the file in arg one
	file, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Error! Could not read file", err)
	}

	str := string(file)
	editedText := goreloaded.TextEditor(str)

	// Write to a given file in arg 2
	err = os.WriteFile(destFile, []byte(editedText), 0644)

	if err != nil {
		fmt.Println()
	 } else {
		fmt.Printf("File %s written succesfuly to %s\n",sourceFile, destFile)
	}
	
}
