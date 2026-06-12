package main

import (
	"bufio"
	"fmt"
	"os"

	"push_swap/internal/ops"
	"push_swap/internal/parser"
	"push_swap/internal/solver"
	"push_swap/internal/stack"
)

func main() {
	if len(os.Args) < 2 {
		// fmt.Println("Usage: go run exec/checker.cmd/Main.go <numbers> || ./checker <numbers>")
		return
	}

	nums, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println("Error")
		return
	}

	a := stack.NewFromSlice(nums)
	b := stack.New()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		move := scanner.Text()
		if move == "" {
			continue
		}
		if err := executeMove(move, a, b); err != nil {
			fmt.Println("Error")
			return
		}
	}

	if solver.IsSorted(a) && b.Size == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func executeMove(move string, a, b *stack.Stack) error {
	switch move {
	case "sa":
		ops.Sa(a)
	case "sb":
		ops.Sb(b)
	case "ss":
		ops.Ss(a, b)
	case "pa":
		ops.Pa(a, b)
	case "pb":
		ops.Pb(a, b)
	case "ra":
		ops.Ra(a)
	case "rb":
		ops.Rb(b)
	case "rr":
		ops.Rr(a, b)
	case "rra":
		ops.Rra(a)
	case "rrb":
		ops.Rrb(b)
	case "rrr":
		ops.Rrr(a, b)
	default:
		return fmt.Errorf("invalid move: %s", move)
	}
	return nil
}
