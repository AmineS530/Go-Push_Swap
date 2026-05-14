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
	maxNode := a.Head
	maxPos := 0

	for i := 0; i < a.Size; i++ {
		// track the max node as we go
		if curr.Index > maxNode.Index {
			maxNode = curr
			maxPos = i
		}

		// does index fit in the gap between curr and curr.Next()?
		if curr.Index < index && index < curr.Next().Index {
			return curr, i
		}

		curr = curr.Next()
	}

	// no gap found → v is out of range of all gaps → goes after max
	return maxNode, maxPos
}
