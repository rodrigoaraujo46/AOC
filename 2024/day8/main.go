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

	var matrix [][]string

	for _, line := range lines {
		line := strings.Split(line, "")
		matrix = append(matrix, line)
	}
	fmt.Println(matrix[0][0])
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

type Point struct {
	X   int
	Y   int
	Val string
}

func AntiPoint(a *Point, b *Point) *Point {
	return &Point{
		X:   a.X + 2*(b.X-a.X),
		Y:   a.Y + 2*(b.Y-a.Y),
		Val: "#",
	}
}

func (p *Point) inBounds(matrix [][]string) bool {
	return 0 <= p.X && p.X < len(matrix[0]) && 0 <= p.Y && p.Y < len(matrix)
}

func getAntinodes1(antinodeMatrix [][]bool, matrix [][]string, startPoint *Point) {
	for y, row := range matrix {
		for x, char := range row {
			if char == startPoint.Val {
				endPoint := &Point{X: x, Y: y, Val: startPoint.Val}
				antiPoint := AntiPoint(startPoint, endPoint)
				if antiPoint.inBounds(matrix) && matrix[antiPoint.Y][antiPoint.X] != endPoint.Val{
					antinodeMatrix[antiPoint.Y][antiPoint.X] = true
				}
			}
		}
	}
}

func getAntinodes2(antinodeMatrix [][]bool, matrix [][]string, startPoint *Point) {
	for y, row := range matrix {
		for x, char := range row {
			if char == startPoint.Val {
				endPoint := &Point{X: x, Y: y, Val: startPoint.Val}
				antinodeMatrix[startPoint.Y][startPoint.X] = true
				antinodeMatrix[endPoint.Y][endPoint.X] = true
				antiPoint := AntiPoint(startPoint, endPoint)
				for antiPoint.inBounds(matrix) && matrix[antiPoint.Y][antiPoint.X] != endPoint.Val {
					antinodeMatrix[antiPoint.Y][antiPoint.X] = true
					helpPoint := antiPoint
					antiPoint = AntiPoint(endPoint, antiPoint)
					endPoint = helpPoint
				}
			}
		}
	}
}

func countAntiNodes(antiMatrix [][]bool) int {
	ans := 0
	for _, row := range antiMatrix {
		for _, in := range row {
			if in {
				ans++
			}
		}
	}
	return ans
}

func part1(matrix [][]string) int {
	antinodeMatrix := make([][]bool, len(matrix))
	for i := range antinodeMatrix {
		antinodeMatrix[i] = make([]bool, len(matrix[0]))
	}
	for y, row := range matrix {
		for x, char := range row {
			if char != "." {
				getAntinodes1(antinodeMatrix, matrix, &Point{X: x, Y: y, Val: char})
			}
		}
	}
	return countAntiNodes(antinodeMatrix)
}

func part2(matrix [][]string) int {
	antinodeMatrix := make([][]bool, len(matrix))
	for i := range antinodeMatrix {
		antinodeMatrix[i] = make([]bool, len(matrix[0]))
	}
	for y, row := range matrix {
		for x, char := range row {
			if char != "." {
				getAntinodes2(antinodeMatrix, matrix, &Point{X: x, Y: y, Val: char})
			}
		}
	}
	return countAntiNodes(antinodeMatrix)
}
