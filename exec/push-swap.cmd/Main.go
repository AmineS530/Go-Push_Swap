package main

import (
	"fmt"
	"os"

	"push_swap/internal/parser"
	"push_swap/internal/solver"
)

func main() {
	if len(os.Args) < 2 {
		// fmt.Println("Usage: go run exec/push-swap.cmd/Main.go <numbers> || ./push-swap \"<numbers>\"")
		return
	}

	numbers, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println("\033[31mError:", err, "\033[0m")
		return
	}

	moves := solver.Solve(numbers)
	for _, move := range moves {
		fmt.Println(move)
	}
}
