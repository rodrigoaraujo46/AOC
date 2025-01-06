package main

import (
	"fmt"
	"os"
	"strings"
)

type lock [5]int
type key [5]int

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	locks := make([]lock, 0)
	keys := make([]key, 0)

	matrix := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			if lock, key := transformMatrix(matrix); lock == nil {
				keys = append(keys, *key)
			} else {
				locks = append(locks, *lock)
			}
			matrix = make([]string, 0)
			continue
		}
		matrix = append(matrix, line)
	}
	if lock, key := transformMatrix(matrix); lock == nil {
		keys = append(keys, *key)
	} else {
		locks = append(locks, *lock)
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v", part1(locks, keys))
		return
	}
	fmt.Printf("PART 1: %v", part1(locks, keys))
	return
}

func transformMatrix(matrix []string) (*lock, *key) {
	if !strings.ContainsRune(matrix[0], '.') {
		lock := new(lock)
		for _, line := range matrix {
			for j, char := range line {
				if char == '#' {
					lock[j]++
				}
			}
		}

		return lock, nil
	}

	key := new(key)
	for _, line := range matrix {
		for i, char := range line {
			if char == '#' {
				key[i]++
			}
		}
	}

	return nil, key
}

func canFit(lock lock, key key) bool {
	for i := range lock {
		if lock[i]+key[i] > 7 {
			return false
		}
	}
	return true
}

func part1(locks []lock, keys []key) int {
	ans := 0
	for _, lock := range locks {
		for _, key := range keys {
			if canFit(lock, key) {
				ans++
			}
		}
	}
	return ans
}
