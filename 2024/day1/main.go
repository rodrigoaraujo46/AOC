package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	var col1 []int
	var col2 []int
	for _, string := range lines {
		row := strings.Split(string, "   ")

		num1, _ := strconv.Atoi(row[0])
		num2, _ := strconv.Atoi(row[1])

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}
	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(col1, col2))
		fmt.Printf("PART 2: %v\n", part2(col1, col2))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(col1, col2))
		return
	}
	fmt.Printf("PART 1: %v", part1(col1, col2))
}
func part2(col1 []int, col2 []int) int {
	count := make(map[int]int)
	ans := 0
	for _, val := range col2 {
		count[val]++
	}

	for _, val := range col1 {
		ans += val * count[val]
	}

	return ans
}

func part1(col1 []int, col2 []int) int {
	sort.Ints(col1)
	sort.Ints(col2)

	ans := 0
	for i, num1 := range col1 {
		dif := num1 - col2[i]
		if dif < 0 {
			dif = -dif
		}
		ans += dif
	}
	return ans
}
