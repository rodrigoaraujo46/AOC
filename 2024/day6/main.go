package main

import (
	"fmt"
	"os"
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

	var runes [][]rune

	for _, line := range lines {
		runes = append(runes, []rune(line))
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(runes))
		fmt.Printf("PART 2: %v\n", part2(runes))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(runes))
		return
	}
	fmt.Printf("PART 1: %v", part1(runes))
}

type Visited struct {
	X   int
	Y   int
	Dir int
}

func part2(runes [][]rune) int {
	directions := [4][2]int{

		{0, -1},

		{1, 0},

		{0, 1},

		{-1, 0},
	}

	X, Y := 0, 0
	for l, line := range runes {
		for r, rune := range line {
			if rune == '^' {
				Y = l
				X = r
				break
			}
		}
	}

	ans := 0
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes); j++ {

			directionIdx := 0
			direction := directions[directionIdx]
			visited := make(map[Visited]bool)
			x, y := X, Y

			for 0 <= x && x < len(runes[0]) && 0 <= y && y < len(runes) {
				if runes[y][x] == '#' || (y == i && x == j) {
					x -= direction[0]
					y -= direction[1]
					directionIdx = (directionIdx + 1) % 4
					direction = directions[directionIdx]
				} else {
					V := Visited{X: x, Y: y, Dir: directionIdx}
					if visited[V] {
						ans++
						break
					}
					visited[V] = true
				}
				x += direction[0]
				y += direction[1]
			}
		}
	}
	return ans
}

func part1(runes [][]rune) int {
	directions := [4][2]int{

		{0, -1},

		{1, 0},

		{0, 1},

		{-1, 0},
	}

	x, y := 0, 0
	for l, line := range runes {
		for r, rune := range line {
			if rune == '^' {
				y = l
				x = r
				break
			}
		}
	}
	directionIdx := 0
	direction := directions[directionIdx]
	ans := 0
	type Visited struct {
		X int
		Y int
	}

	visited := make(map[Visited]bool)
	for x < len(runes[0]) && y < len(runes) {
		if runes[y][x] == '#' {
			x -= direction[0]
			y -= direction[1]
			directionIdx = (directionIdx + 1) % 4
			direction = directions[directionIdx]
		} else {
			V := Visited{X: x, Y: y}
			if !visited[V] {
				ans++
				visited[V] = true
			}
		}
		x += direction[0]
		y += direction[1]
	}
	return ans
}
