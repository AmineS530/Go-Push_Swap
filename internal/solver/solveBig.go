package solver

import (
	"push_swap/internal/ops"
	"push_swap/internal/stack"
)

func solveBig(a, b *stack.Stack) []string {
	moves := []string{}

	// step 1: push all but 3 to B
	for a.Size > 3 {
		moves = append(moves, ops.Pb(a, b))
	}

	// step 2: sort the 3 remaining in A
	moves = append(moves, solveThree(a)...)

	// step 3: insert each B node back into A
	for b.Size > 0 {
		moves = append(moves, insertBest(a, b)...)
	}

	// step 4: rotate A until min is on top
	moves = append(moves, rotateToMin(a)...)

	return moves
}

func findTargetInA(index int, a *stack.Stack) (*stack.Node, int) {
	curr := a.Head

	for i := 0; i < a.Size; i++ {
		next := curr.Next()

		// normal ascending gap
		if curr.Index < next.Index {
			if index > curr.Index && index < next.Index {
				return next, (i + 1) % a.Size
			}
		} else {
			// wrap-around gap (max -> min)
			if index > curr.Index || index < next.Index {
				return next, (i + 1) % a.Size
			}
		}

		curr = next
	}

	return a.Head, 0
}
