package solver

import "push_swap/internal/stack"

func Solve(a, b *stack.Stack) []string {
	if IsSorted(a) {
		return nil
	}

	switch a.Size {
	case 2:
		return solveTwo(a)
	case 3:
		return solveThree(a)
	default:
		return solveBig(a, b)
	}
}

func solveBig(a, b *stack.Stack) []string {
	moves := []string{}

	return moves
}
