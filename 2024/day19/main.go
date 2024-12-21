package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func countPossible(towel string, patterns []string, memo map[string]int, maxSize int) int {
	if towel == "" {
		return 1
	}

	if val, ok := memo[towel]; ok {
		return val
	}

	ans := 0
	for i := 0; i < min(maxSize, len(towel)); i++ {
		if slices.Contains(patterns, towel[:i+1]) {
			ans += countPossible(towel[i+1:], patterns, memo, maxSize)
		}
	}
	memo[towel] = ans
	return ans
}

func part2(patterns, towels []string) int {
	ans := 0

	maxPattern := 0
	for _, pattern := range patterns {
		if len(pattern) > maxPattern {
			maxPattern = len(pattern)
		}
	}

	memo := make(map[string]int)
	for _, towel := range towels {
		ans += countPossible(towel, patterns, memo, maxPattern)
	}
	return ans
}

func isPossible(towel string, patterns []string, memo map[string]bool, maxSize int) bool {
	if towel == "" {
		return true
	}

	if val, ok := memo[towel]; ok {
		return val
	}

	for i := 0; i < min(maxSize, len(towel)); i++ {
		if slices.Contains(patterns, towel[:i+1]) {
			if isPossible(towel[i+1:], patterns, memo, maxSize) {
				memo[towel] = true
				return true
			}
		}
	}
	memo[towel] = false
	return false
}

func part1(patterns, towels []string) int {
	ans := 0

	maxPattern := 0
	for _, pattern := range patterns {
		if len(pattern) > maxPattern {
			maxPattern = len(pattern)
		}
	}

	memo := make(map[string]bool)
	for _, towel := range towels {
		if isPossible(towel, patterns, memo, maxPattern) {
			ans++
		}
	}
	return ans
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	patterns := strings.Split(lines[0], ", ")
	towels := lines[2:]

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(patterns, towels))
		fmt.Printf("PART 2: %v\n", part2(patterns, towels))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(patterns, towels))
		return
	}
	fmt.Printf("PART 1: %v", part1(patterns, towels))
}
