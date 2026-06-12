package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseArgs(args []string) ([]int, error) {
	numStrs, err := validateArgs(args)
	if err != nil {
		return nil, err
	}
	numbers, err := convertNumbers(numStrs)
	if err != nil {
		return nil, err
	}
	if err := checkDups(numbers); err != nil {
		return nil, err
	}
	return numbers, nil
}

func validateArgs(args []string) ([]string, error) {
	numStrs := []string{}
	// under normal conditions it shouldn't enter this block, but just in case
	if len(args) == 0 {
		return nil, fmt.Errorf("No numbers provided")
	}
	if len(args) == 1 {
		numStrs = strings.Fields(args[0])
		if len(numStrs) == 0 {
			return nil, fmt.Errorf("No numbers provided")
		}
	} else if len(args) > 1 {
		numStrs = args
	}
	return numStrs, nil
}

func convertNumbers(numStrs []string) ([]int, error) {
	var numbers []int
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid number: %s", numStr)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func checkDups(numbers []int) error {
	seen := make(map[int]bool)
	for _, num := range numbers {
		if seen[num] {
			return fmt.Errorf("Duplicate number found: %d", num)
		}
		seen[num] = true
	}
	return nil
}
