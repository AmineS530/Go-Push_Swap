package solver

import (
	"push_swap/internal/ops"
	"push_swap/internal/stack"
)

func IsSorted(a *stack.Stack) bool {
	if a.Size <= 1 {
		return true
	}
	curr := a.Head
	for i := 0; i < a.Size-1; i++ {
		if curr.Index > curr.Next().Index {
			return false
		}
		curr = curr.Next()
	}
	return true
}

func positionOf(node *stack.Node, s *stack.Stack) int {
	curr := s.Head
	for i := 0; i < s.Size; i++ {
		if curr == node {
			return i
		}
		curr = curr.Next()
	}
	return -1 // should never happen
}

func costToTop(pos, size int) (int, bool) {
	forward := pos
	backward := size - pos
	if forward <= backward {
		return forward, true // true = forward direction
	}
	return backward, false // false = backward direction
}

func computeCost(nodeB *stack.Node, a, b *stack.Stack) int {
	// where is this node in B
	posB := positionOf(nodeB, b)

	// cost + direction to bring it to top of B
	costB, forwardB := costToTop(posB, b.Size)

	// find where it needs to go in A
	_, posA := findTargetInA(nodeB.Index, a)

	// cost + direction to bring that target to top of A
	costA, forwardA := costToTop(posA, a.Size)

	// same direction → pay max (use rr/rrr)
	if forwardB == forwardA {
		if costB > costA {
			return costB
		}
		return costA
	}

	// different directions → pay both
	return costB + costA
}

func bestCandidate(a, b *stack.Stack) *stack.Node {
	best := b.Head
	bestCost := computeCost(b.Head, a, b)

	curr := b.Head.Next()
	for i := 1; i < b.Size; i++ {
		cost := computeCost(curr, a, b)
		if cost < bestCost {
			bestCost = cost
			best = curr
		}
		curr = curr.Next()
	}

	return best
}

func posOfIndex(a *stack.Stack, target int) int {
	curr := a.Head

	for i := 0; i < a.Size; i++ {
		if curr.Index == target {
			return i
		}
		curr = curr.Next()
	}

	return -1
}

func moveIndexToTop(a *stack.Stack, target int) []string {
	moves := []string{}

	pos := posOfIndex(a, target)

	cost, forward := costToTop(pos, a.Size)

	for i := 0; i < cost; i++ {
		if forward {
			moves = append(moves, ops.Ra(a))
		} else {
			moves = append(moves, ops.Rra(a))
		}
	}

	return moves
}
