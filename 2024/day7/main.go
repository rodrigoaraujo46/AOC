package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		print(err.Error())
		return
	}

	lines := strings.Split(string(content), ("\n"))
	lines = lines[:len(lines)-1]

	results := make([]int, len(lines))
	equations := make([][]int, len(lines))

	for i, line := range lines {
		line := strings.Split(line, ":")
		result, _ := strconv.Atoi(line[0])
		results[i] = result
		var equation []int
		for _, num := range strings.Split(line[1], " ")[1:] {
			numI, _ := strconv.Atoi(num)
			equation = append(equation, numI)
		}
		equations[i] = equation
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(results, equations))
		fmt.Printf("PART 2: %v\n", part2(results, equations))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(results, equations))
		return
	}
	fmt.Printf("PART 1: %v", part1(results, equations))
}

func part2(results []int, equations [][]int) int {
	ans := 0
	for i, val := range results {
		if calc2(equations[i], val) == val {
			ans += val
		}
	}
	return ans
}

func calc2(params []int, expected int) int {
	if len(params) == 1 {
		return params[0]
	}
	current := params[len(params)-1]
	mult := expected / current
	if expected%current == 0 {
		if mult == calc2(params[:len(params)-1], mult) {
			return expected
		}
	}
	add := expected - current
	if add == calc2(params[:len(params)-1], add) {
		return expected
	}
	currentString := strconv.Itoa(current)
	expectedString := strconv.Itoa(expected)

	conc, found := strings.CutSuffix(expectedString, currentString)
	if !found {
		return -1
	}
	concInt, _ := strconv.Atoi(conc)
	if concInt == calc2(params[:len(params)-1], concInt) {
		return expected
	}
	return -1
}

func calc(params []int, expected int) int {
	if len(params) == 1 {
		return params[0]
	}
	current := params[len(params)-1]
	mult := expected / current
	if expected%current == 0 {
		if mult == calc(params[:len(params)-1], mult) {
			return expected
		}
	}
	add := expected - current
	if add == calc(params[:len(params)-1], add) {
		return expected
	}
	return -1
}

func part1(results []int, equations [][]int) int {
	ans := 0
	for i, val := range results {
		if calc(equations[i], val) == val {
			ans += val
		}
	}
	return ans
}
