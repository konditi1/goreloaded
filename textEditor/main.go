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
	str2 := goreloaded.Upper(str)
	upper := goreloaded.Lower(str2)
	lower := goreloaded.Capitalize(upper)
	hexa := goreloaded.HexaDeci(lower)
	bin := goreloaded.Binary(hexa)
	fmt.Println(bin)
}
