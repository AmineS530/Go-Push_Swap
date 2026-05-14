package solver

import (
	"push_swap/internal/ops"
	"push_swap/internal/stack"
)

func insertBest(a, b *stack.Stack) []string {
	moves := []string{}

	node := bestCandidate(a, b)
	posB := positionOf(node, b)
	_, posA := findTargetInA(node.Index, a)

	costB, forwardB := costToTop(posB, b.Size)
	costA, forwardA := costToTop(posA, a.Size)

	// simultaneous rotations
	simultaneous := 0
	if forwardB == forwardA {
		if costB < costA {
			simultaneous = costB
		} else {
			simultaneous = costA
		}
		for i := 0; i < simultaneous; i++ {
			if forwardB {
				moves = append(moves, ops.Rr(a, b))
			} else {
				moves = append(moves, ops.Rrr(a, b))
			}
		}
		costB -= simultaneous
		costA -= simultaneous
	}

	// remaining B
	for i := 0; i < costB; i++ {
		if forwardB {
			moves = append(moves, ops.Rb(b))
		} else {
			moves = append(moves, ops.Rrb(b))
		}
	}

	// remaining A
	for i := 0; i < costA; i++ {
		if forwardA {
			moves = append(moves, ops.Ra(a))
		} else {
			moves = append(moves, ops.Rra(a))
		}
	}

	moves = append(moves, ops.Pa(a, b))
	return moves
}

func rotateToMin(a *stack.Stack) []string {
	moves := []string{}

	// find position of minimum (index 0)
	minPos := 0
	curr := a.Head
	for i := 0; i < a.Size; i++ {
		if curr.Index == 0 {
			minPos = i
			break
		}
		curr = curr.Next()
	}

	_, forward := costToTop(minPos, a.Size)
	cost, _ := costToTop(minPos, a.Size)

	for i := 0; i < cost; i++ {
		if forward {
			moves = append(moves, ops.Ra(a))
		} else {
			moves = append(moves, ops.Rra(a))
		}
	}

	return moves
}
