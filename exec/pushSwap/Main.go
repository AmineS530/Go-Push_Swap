package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run Main.go <numbers>")
		return
	}

	// numbers, err := parser.ParseArgs(os.Args[1:])
	// fmt.Println(numbers, err)
}
