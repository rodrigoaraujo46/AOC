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

	directions := string(content)

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(directions))
		fmt.Printf("PART 2: %v\n", part2(directions))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(directions))
		return
	}
	fmt.Printf("PART 1: %v", part1(directions))
}

func part2(directions string) int {
	p, pR := point{0, 0}, point{0, 0}
	visited := make(map[point]struct{})
	visited[p] = struct{}{}

	for i := 0; i < len(directions); i += 2 {
		dir := directions[i]
		switch dir {
		case '^':
			p.y += 1
		case '>':
			p.x += 1
		case 'v':
			p.y -= 1
		case '<':
			p.x -= 1
		}
		visited[p] = struct{}{}

		if i == len(directions)-1 {
			continue
		}
		dirR := directions[i+1]
		switch dirR {
		case '^':
			pR.y += 1
		case '>':
			pR.x += 1
		case 'v':
			pR.y -= 1
		case '<':
			pR.x -= 1
		}
		visited[pR] = struct{}{}
	}
	return len(visited)
}

type point struct {
	x int
	y int
}

func part1(directions string) int {
	p := point{0, 0}
	visited := make(map[point]struct{})
	visited[p] = struct{}{}

	for _, dir := range directions {
		switch dir {
		case '^':
			p.y += 1
		case '>':
			p.x += 1
		case 'v':
			p.y -= 1
		case '<':
			p.x -= 1
		}
		visited[p] = struct{}{}
	}
	return len(visited)
}
