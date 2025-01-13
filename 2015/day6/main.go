package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type instruction struct {
	toggle bool
	start  point
	end    point
	on     bool
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instruction := instruction{}
		splitLine := strings.Split(line, " ")
		var nums1, nums2 []string
		if splitLine[0] == "turn" {
			instruction.toggle = false
			if splitLine[1] == "on" {
				instruction.on = true
			}
			nums1, nums2 = strings.Split(splitLine[2], ","), strings.Split(splitLine[4], ",")
		} else {
			instruction.toggle = true
			nums1, nums2 = strings.Split(splitLine[1], ","), strings.Split(splitLine[3], ",")
		}
		startX, _ := strconv.Atoi(nums1[0])
		startY, _ := strconv.Atoi(nums1[1])
		endX, _ := strconv.Atoi(nums2[0])
		endY, _ := strconv.Atoi(nums2[1])
		start := point{startX, startY}
		end := point{endX, endY}
		instruction.start = start
		instruction.end = end
		instructions[i] = instruction
	}

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

func part2(instructions []instruction) int {
	var matrix [1000][1000]int
	for _, inst := range instructions {
		for y := inst.start.y; y <= inst.end.y; y++ {
			for x := inst.start.x; x <= inst.end.x; x++ {
				if inst.toggle {
					matrix[y][x] += 2
					continue
				}
				if inst.on {
					matrix[y][x]++
					continue
				}
				matrix[y][x] = max(matrix[y][x]-1, 0)
			}
		}
	}
	ans := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			ans += matrix[y][x]
		}
	}
	return ans
}

func part1(instructions []instruction) int {
	var matrix [1000][1000]bool
	for _, inst := range instructions {
		for y := inst.start.y; y <= inst.end.y; y++ {
			for x := inst.start.x; x <= inst.end.x; x++ {
				if inst.toggle {
					matrix[y][x] = !matrix[y][x]
					continue
				}
				if inst.on {
					matrix[y][x] = true
					continue
				}
				matrix[y][x] = false
			}
		}
	}
	ans := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if matrix[y][x] {
				ans++
			}
		}
	}
	return ans
}
