package main

import (
	"fmt"
	"os"

	"push_swap/internal/parser"
	"push_swap/internal/solver"
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

	indices := stack.Normalize(numbers)
	stackA := stack.NewFromSlice(numbers, indices)
	stackB := stack.New()

	moves := solver.Solve(stackA, stackB)
	for _, move := range moves {
		fmt.Println(move)
	}
}