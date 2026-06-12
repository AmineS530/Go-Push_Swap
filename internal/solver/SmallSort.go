package solver

import (
	"push_swap/internal/ops"
	"push_swap/internal/stack"
)

func solveTwo(a *stack.Stack) []string {
	if a.Head.Index > a.Head.Next().Index {
		return []string{ops.Sa(a)}
	}
	return nil
}

func solveThree(a *stack.Stack) []string {
	moves := []string{}

	for !IsSorted(a) {
		top := a.Head.Index
		mid := a.Head.Next().Index
		bot := a.Head.Prev().Index

		if top > mid && top > bot {
			moves = append(moves, ops.Ra(a))
		} else if mid > top && mid > bot {
			moves = append(moves, ops.Rra(a))
		} else {
			moves = append(moves, ops.Sa(a))
		}
	}

	return moves
}

func solveFive(a, b *stack.Stack) []string {
	moves := []string{}

	// push index 0
	moves = append(moves, moveIndexToTop(a, 0)...)
	moves = append(moves, ops.Pb(a, b))

	// push index 1
	moves = append(moves, moveIndexToTop(a, 1)...)
	moves = append(moves, ops.Pb(a, b))

	// sort remaining 3
	moves = append(moves, solveThree(a)...)

	// bring back
	moves = append(moves, ops.Pa(a, b))
	moves = append(moves, ops.Pa(a, b))

	return moves
}
