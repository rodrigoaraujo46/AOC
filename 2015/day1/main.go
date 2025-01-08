package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	instructions := string(content)

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(instructions))
		fmt.Printf("PART 2: %v\n", part2(instructions))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(instructions))
		return
	}
	fmt.Printf("PART 1: %v", part1(instructions))
}

func part2(instructions string) int {
	floor := 0
	for i, char := range instructions {
		if char == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return i+1
		}
	}

	return floor
}

func part1(instructions string) int {
	floor := 0
	for _, char := range instructions {
		if char == '(' {
			floor++
		} else {
			floor--
		}
	}

	return floor
}
