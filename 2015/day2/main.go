package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dimension struct {
	l int
	w int
	h int
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	dimensions := make([]dimension, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, "x")

		l, _ := strconv.Atoi(nums[0])
		w, _ := strconv.Atoi(nums[1])
		h, _ := strconv.Atoi(nums[2])

		dimensions[i] = dimension{l, w, h}

	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(dimensions))
		fmt.Printf("PART 2: %v\n", part2(dimensions))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(dimensions))
		return
	}
	fmt.Printf("PART 1: %v", part1(dimensions))
}

func part2(dimensions []dimension) int {
	feet := 0
	for _, d := range dimensions {
		bow := d.l * d.w * d.h

		biggest := max(d.l, d.w, d.h)
		wrap := 2*d.l + 2*d.w + 2*d.h - 2*biggest

		feet += bow + wrap
	}

	return feet
}

func part1(dimensions []dimension) int {
	squareFeet := 0
	for _, d := range dimensions {
		side1 := d.l * d.w
		side2 := d.w * d.h
		side3 := d.h * d.l
		squareFeet += 2*side1 + 2*side2 + 2*side3 + min(side1, side2, side3)
	}

	return squareFeet
}
