package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	strings := strings.Split(string(content), "\n")
	strings = strings[:len(strings)-1]

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(strings))
		fmt.Printf("PART 2: %v\n", part2(strings))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(strings))
		return
	}
	fmt.Printf("PART 1: %v", part1(strings))
}

func part1(strs []string) int {
	ans := 0
Outer:
	for _, s := range strs {
		vowels := 0
		if strings.ContainsRune("aeiou", rune(s[0])) {
			vowels++
		}
		row := false
		for j := 1; j < len(s); j++ {
			if strings.ContainsRune("aeiou", rune(s[j])) {
				vowels++
			}
			if s[j] == s[j-1] {
				row = true
			}
			if slices.Contains([]string{"ab", "cd", "pq", "xy"}, string(s[j-1])+string(s[j])) {
				continue Outer
			}
		}
		if vowels >= 3 && row {
			ans++
		}
	}
	return ans
}

type pair struct {
	a any
	b any
}

func makePair(a, b any) pair {
	return pair{a, b}
}

func part2(strs []string) int {
	ans := 0
	for _, s := range strs {
		firstRule := false
		pairs := make(map[pair]int)
		for j := 1; j < len(s); j += 1 {
			pair := makePair(s[j-1], s[j])
			if k, ok := pairs[pair]; ok && k < j-1 {
				firstRule = true
				break
			}
			pairs[pair] = j
		}
		secondRule := false
		for j := 2; j < len(s); j += 1 {
			if s[j] == s[j-2] {
				secondRule = true
				break
			}
		}
		if firstRule && secondRule {
			ans++
		}
	}
	return ans
}
