package solver

import "push_swap/internal/stack"

func Solve(nums []int) []string {
	stackA := stack.NewFromSlice(nums)
	stackB := stack.New()
	
	if IsSorted(stackA) {
		return nil
	}

	switch stackA.Size {
	case 2:
		return solveTwo(stackA)
	case 3:
		return solveThree(stackA)
	default:
		return solveBig(stackA, stackB)
	}
}
