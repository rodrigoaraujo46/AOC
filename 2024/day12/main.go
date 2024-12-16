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

type Corner struct {
	X float64
	Y float64
}

func (c Corner) GetPlantsFromCorner() [4]Plant {
	halfX := 0.5
	halfY := 0.5

	return [4]Plant{
		{X: int(c.X - halfX), Y: int(c.Y - halfY)}, // Bottom-left
		{X: int(c.X + halfX), Y: int(c.Y - halfY)}, // Bottom-right
		{X: int(c.X - halfX), Y: int(c.Y + halfY)}, // Top-left
		{X: int(c.X + halfX), Y: int(c.Y + halfY)}, // Top-right
	}
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

func (p *Plant) GetCorners() [4]Corner {
	return [4]Corner{
		{float64(p.X) + 0.5, float64(p.Y) + 0.5}, // Top-right
		{float64(p.X) - 0.5, float64(p.Y) + 0.5}, // Top-left
		{float64(p.X) - 0.5, float64(p.Y) - 0.5}, // Bottom-left
		{float64(p.X) + 0.5, float64(p.Y) - 0.5}, // Bottom-right
	}
}

func (previous *Plant) CountCorners(region [][]rune, cornersChecked map[Corner]bool) int {
	corners := 0
	possibleCorners := previous.GetCorners()
	for _, corner := range possibleCorners {
		if cornersChecked[corner] {
			continue
		}
		cornersChecked[corner] = true
		plants := corner.GetPlantsFromCorner()
		empties := make([]bool, 4)
		count := 0
		for i, plant := range plants {
			if !plant.InBounds(region) {
				empties[i] = true
				count++
				continue
			}
			plant.Type = region[plant.Y][plant.X]
			if plant.Type != previous.Type {
				empties[i] = true
				count++
			}
		}
		if count == 1 {
			corners++
		}
		if count == 2 && (slices.Equal(empties, []bool{true, false, false, true}) || slices.Equal(empties, []bool{false, true, true, false})) {
			corners += 2
		}
		if count == 3 {
			corners++
		}
	}
	return corners
}

func printMatrix(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%c ", matrix[i][j]) // Print the rune as a character
		}
		fmt.Println() // Move to the next line after each row
	}
}

func (previous *Plant) GetAreaPrice2(region [][]rune, visited [][]bool, perimeter *int, area *int, cornersChecked map[Corner]bool) {
	for i := range 4 {
		p := previous.Next(i)
		if !p.InBounds(region) {
			continue
		}

		p.Type = region[p.Y][p.X]
		if p.Type != previous.Type {
			continue
		}

		if visited[p.Y][p.X] {
			continue
		}

		corners := p.CountCorners(region, cornersChecked)
		*perimeter += corners

		*area++
		visited[p.Y][p.X] = true
		p.GetAreaPrice2(region, visited, perimeter, area, cornersChecked)
	}
}

func (previous *Plant) GetArea(matrix [][]rune, region [][]rune, visited [][]bool) {
	for i := range 4 {
		p := previous.Next(i)
		if !p.InBounds(matrix) {
			continue
		}
		p.Type = matrix[p.Y][p.X]
		if p.Type != previous.Type {
			continue
		}

		if visited[p.Y][p.X] {
			continue
		}
		visited[p.Y][p.X] = true
		region[p.Y][p.X] = '!'
		p.GetArea(matrix, region, visited)
	}
}

func part2(matrix [][]rune) int {
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	visited2 := make([][]bool, len(matrix))
	for i := range visited {
		visited2[i] = make([]bool, len(matrix[0]))
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
			region := make([][]rune, len(matrix))
			for i := range matrix {
				region[i] = append([]rune(nil), matrix[i]...)
			}

			visited[plant.Y][plant.X] = true
			region[plant.Y][plant.X] = '!'
			plant.GetArea(matrix, region, visited2)

			plant.Type = '!'
			area := 1
			cornersChecked := make(map[Corner]bool)
			perimeter := plant.CountCorners(region, cornersChecked)
			plant.GetAreaPrice2(region, visited, &perimeter, &area, cornersChecked)

			ans += area * perimeter
		}
	}
	return ans
}
