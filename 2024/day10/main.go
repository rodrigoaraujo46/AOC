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

	matrix := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, numS := range line {
			row[j] = int(numS - '0')
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

type Point struct {
	X   int
	Y   int
	Val int
}

func (p *Point) Left() *Point {
	return &Point{X: p.X - 1, Y: p.Y, Val: p.Val}
}
func (p *Point) Up() *Point {
	return &Point{X: p.X, Y: p.Y - 1, Val: p.Val}
}
func (p *Point) Right() *Point {
	return &Point{X: p.X + 1, Y: p.Y, Val: p.Val}
}
func (p *Point) Down() *Point {
	return &Point{X: p.X, Y: p.Y + 1, Val: p.Val}
}
func (p *Point) InBounds(matrix [][]int) bool {
	return 0 <= p.X && p.X < len(matrix[0]) && 0 <= p.Y && p.Y < len(matrix)
}

func getScore(matrix [][]int, visited [][]int, p *Point) int {
	if !p.InBounds(matrix) || matrix[p.Y][p.X] != p.Val+1 || visited[p.Y][p.X] == -1 {
		return 0
	}
	p.Val++
	if p.Val == 9 {
		visited[p.Y][p.X] = -1
		return 1
	}
	return getScore(matrix, visited, p.Left()) + getScore(matrix, visited, p.Up()) + getScore(matrix, visited, p.Right()) + getScore(matrix, visited, p.Down())
}

func getScore2(matrix [][]int, p *Point) int {
	if !p.InBounds(matrix) || matrix[p.Y][p.X] != p.Val+1 {
		return 0
	}
	p.Val++
	if p.Val == 9 {
		return 1
	}
	return getScore2(matrix, p.Left()) + getScore2(matrix, p.Up()) + getScore2(matrix, p.Right()) + getScore2(matrix, p.Down())
}
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%4d", val) // Adjust width for alignment
		}
		fmt.Println()
	}
	println()
}
func part1(matrix [][]int) (ans int) {
	for y, row := range matrix {
		for x, num := range row {
			if num == 0 {
				visited := make([][]int, len(matrix))
				for i, row := range matrix {
					visited[i] = make([]int, len(row))
					copy(visited[i], row)
				}

				p := &Point{X: x, Y: y, Val: num}
				ans += getScore(matrix, visited, p.Left()) + getScore(matrix, visited, p.Up()) + getScore(matrix, visited, p.Right()) + getScore(matrix, visited, p.Down())
			}
		}
	}
	return ans
}

func part2(matrix [][]int) (ans int) {
	for y, row := range matrix {
		for x, num := range row {
			if num == 0 {
				visited := make([][]int, len(matrix))
				for i, row := range matrix {
					visited[i] = make([]int, len(row))
					copy(visited[i], row)
				}

				p := &Point{X: x, Y: y, Val: num}
				ans += getScore2(matrix, p.Left()) + getScore2(matrix, p.Up()) + getScore2(matrix, p.Right()) + getScore2(matrix, p.Down())
			}
		}
	}
	return ans
}
