package main

import (
	"fmt"
	"os"

	"push_swap/internal/parser"
	"push_swap/internal/stack"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run Main.go <numbers>")
		return
	}

	numbers, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println("Error parsing arguments:", err)
		return
	}
	fmt.Println(numbers)

	numbers = stack.Normalize(numbers)
	
	fmt.Println(numbers)
}
