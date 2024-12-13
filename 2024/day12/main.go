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

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		row := make([]rune, len(line))
		for j, numS := range line {
			row[j] = numS
		}
		matrix[i] = row
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(matrix))
		fmt.Printf("PART 2: %v\n", part2(matrix))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(matrix))
		return
	}
	fmt.Printf("PART 1: %v", part1(matrix))
}

type Plant struct {
	X    int
	Y    int
	Type rune
}

func (p *Plant) Up() *Plant {
	return &Plant{
		X:    p.X,
		Y:    p.Y - 1,
		Type: p.Type,
	}
}

func (p *Plant) Right() *Plant {
	return &Plant{
		X:    p.X + 1,
		Y:    p.Y,
		Type: p.Type,
	}
}

func (p *Plant) Down() *Plant {
	return &Plant{
		X:    p.X,
		Y:    p.Y + 1,
		Type: p.Type,
	}
}

func (p *Plant) Left() *Plant {
	return &Plant{
		X:    p.X - 1,
		Y:    p.Y,
		Type: p.Type,
	}
}

func (p *Plant) Next(direction int) *Plant {
	switch direction {
	case 0:
		return p.Up()
	case 1:
		return p.Right()
	case 2:
		return p.Down()
	}
	return p.Left()
}

func (p *Plant) InBounds(matrix [][]rune) bool {
	return 0 <= p.X && p.X < len(matrix[0]) && 0 <= p.Y && p.Y < len(matrix)
}

type Side struct {
	A Plant
	B Plant
}

func (previous *Plant) GetAreaPrice(matrix [][]rune, visited [][]bool, perimeter *int, area *int, removed map[Side]bool) {
	for i := range 4 {
		p := previous.Next(i)
		if !p.InBounds(matrix) {
			continue
		}
		p.Type = matrix[p.Y][p.X]
		if p.Type != previous.Type {
			continue
		}

		s := Side{A: *previous, B: *p}
		s2 := Side{A: *p, B: *previous}
		if visited[p.Y][p.X] {
			if removed[s] || removed[s2] {
				continue
			}
			removed[s] = true
			*perimeter -= 2
			continue
		}
		*area++
		*perimeter += 4
		visited[p.Y][p.X] = true
		p.GetAreaPrice(matrix, visited, perimeter, area, removed)
	}
}

func part1(matrix [][]rune) int {
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	ans := 0
	for y, row := range matrix {
		for x, char := range row {
			if visited[y][x] {
				continue
			}

			plant := &Plant{
				X:    x,
				Y:    y,
				Type: char,
			}
			area := 1
			perimeter := 4
			visited[plant.Y][plant.X] = true
			removed := make(map[Side]bool)
			plant.GetAreaPrice(matrix, visited, &perimeter, &area, removed)
			ans += area * perimeter

		}
	}
	return ans
}
func printMatrix(matrix [][]bool) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%4v", val) // Print each value with padding
		}
		fmt.Println() // Move to the next row
	}
	println()
}
func part2(matrix [][]rune) int {
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	ans := 0
	for y, row := range matrix {
		for x, char := range row {
			if visited[y][x] {
				continue
			}

			plant := &Plant{
				X:    x,
				Y:    y,
				Type: char,
			}
			if plant.Type != 'B' {
				continue
			}
			area := 1
			perimeter := 4
			visited[plant.Y][plant.X] = true
			removed := make(map[Side]bool)
			plant.GetAreaPrice(matrix, visited, &perimeter, &area, removed)
			fmt.Println(perimeter)
			ans += area * perimeter

		}
	}
	return ans
}
