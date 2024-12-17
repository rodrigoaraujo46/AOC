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
	//	lines = lines[:len(lines)-1]

	var ocean [][]string
	index := 0
	for i, line := range lines {
		if line == "" {
			index += i + 1
			break
		}
		chars := strings.Split(line, "")
		ocean = append(ocean, chars)
	}

	var moves []string
	for ; index < len(lines); index++ {
		line := lines[index]
		moves = append(moves, line)
	}
	movesString := strings.Join(moves, "")
	directions := strings.Split(movesString, "")
	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(ocean, directions))
		fmt.Printf("PART 2: %v\n", part2(ocean, directions))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(ocean, directions))
		return
	}
	fmt.Printf("PART 1: %v", part1(ocean, directions))
}

type Vector struct {
	X int
	Y int
}

func (c *Vector) Next(direction string) *Vector {
	if direction == "^" {
		return &Vector{
			X: c.X,
			Y: c.Y - 1,
		}
	}
	if direction == ">" {
		return &Vector{
			X: c.X + 1,
			Y: c.Y,
		}
	}
	if direction == "v" {
		return &Vector{
			X: c.X,
			Y: c.Y + 1,
		}
	}
	return &Vector{
		X: c.X - 1,
		Y: c.Y,
	}
}

func getNewOcean(ocean [][]string, direction string, c *Vector) {
	n := c.Next(direction)
	if ocean[n.Y][n.X] == "#" {
		return
	}
	if ocean[n.Y][n.X] == "." {
		ocean[n.Y][n.X] = ocean[c.Y][c.X]
		ocean[c.Y][c.X] = "."
		if ocean[n.Y][n.X] == "@" {
			c.X = n.X
			c.Y = n.Y
		}
		return
	}
	getNewOcean(ocean, direction, n)
	if ocean[n.Y][n.X] == "." {
		ocean[n.Y][n.X] = ocean[c.Y][c.X]
		ocean[c.Y][c.X] = "."
		if ocean[n.Y][n.X] == "@" {
			c.X = n.X
			c.Y = n.Y
		}
		return
	}
	return
}

func part1(ocean [][]string, moves []string) int {
	var c *Vector

Outer:
	for i, line := range ocean {
		for j, char := range line {
			if char == "@" {
				c = &Vector{
					X: j,
					Y: i,
				}
				break Outer
			}
		}
	}
	for _, char := range moves {
		getNewOcean(ocean, char, c)
	}
	ans := 0
	for i, line := range ocean {
		for j, char := range line {
			if char == "O" {
				ans += 100*i + j
			}
		}
	}
	return ans
}
func printCharMatrix(matrix [][]string) {
	// Iterate through each row of the matrix
	for _, row := range matrix {
		// Print each character in the row without spaces between them
		for _, char := range row {
			fmt.Print(char)
		}
		// Move to the next line after each row
		fmt.Println()
	}
}
func getNewOcean2(ocean [][]string, direction string, c *Vector) [][]string {
	n := c.Next(direction)
	if ocean[n.Y][n.X] == "#" {
		return ocean
	}
	if ocean[n.Y][n.X] == "." {
		ocean[n.Y][n.X] = ocean[c.Y][c.X]
		ocean[c.Y][c.X] = "."
		if ocean[n.Y][n.X] == "@" {
			c.X = n.X
			c.Y = n.Y
		}
		return ocean
	}
	if direction != "v" && direction != "^" {
		ocean = getNewOcean2(ocean, direction, n)
		if ocean[n.Y][n.X] == "." {
			ocean[n.Y][n.X] = ocean[c.Y][c.X]
			ocean[c.Y][c.X] = "."
			if ocean[n.Y][n.X] == "@" {
				c.X = n.X
				c.Y = n.Y
			}
			return ocean
		}
		return ocean
	}

	var n2 *Vector
	if ocean[n.Y][n.X] == "[" {
		n2 = &Vector{
			X: n.X + 1,
			Y: n.Y,
		}
	} else {
		n2 = &Vector{
			X: n.X - 1,
			Y: n.Y,
		}
	}

	possible := make([][]string, len(ocean))
	for i := range ocean {
		possible[i] = make([]string, len(ocean[i]))
		copy(possible[i], ocean[i])
	}
	possible = getNewOcean2(possible, direction, n)
	possible = getNewOcean2(possible, direction, n2)

	if possible[n.Y][n.X] == "." && possible[n2.Y][n2.X] == "." {
		ocean = possible
		ocean[n.Y][n.X] = ocean[c.Y][c.X]
		ocean[c.Y][c.X] = "."
		if ocean[n.Y][n.X] == "@" {
			c.X = n.X
			c.Y = n.Y
		}
		return ocean
	}
	return ocean
}

func duplicateChars(ocean [][]string) [][]string {
	var newOcean [][]string

	for _, row := range ocean {
		var newRow []string
		for _, char := range row {
			if char == "O" {
				newRow = append(newRow, "[", "]")
				continue
			}
			if char == "@" {
				newRow = append(newRow, char, ".")
				continue
			}
			newRow = append(newRow, char, char)
		}
		newOcean = append(newOcean, newRow)
	}
	return newOcean
}

func part2(ocean [][]string, moves []string) int {
	ocean = duplicateChars(ocean)

	var c *Vector

Outer:
	for i, line := range ocean {
		for j, char := range line {
			if char == "@" {
				c = &Vector{
					X: j,
					Y: i,
				}
				break Outer
			}
		}
	}
	for _, char := range moves {
		ocean = getNewOcean2(ocean, char, c)
	}
	ans := 0
	for i, line := range ocean {
		for j, char := range line {
			if char == "[" {
				ans += 100*i + j
			}
		}
	}
	return ans
}
