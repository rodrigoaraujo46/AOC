package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	var matrix [][]rune
	for _, line := range lines {
		runes := []rune(line)
		matrix = append(matrix, runes)
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

func part2(matrix [][]rune) int {
	ans := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == 'A' {
				if !valid(matrix, y-1, x-1) {
					continue
				}
				topL := matrix[y-1][x-1]

				if !valid(matrix, y-1, x+1) {
					continue
				}
				topR := matrix[y-1][x+1]

				if !valid(matrix, y+1, x-1) {
					continue
				}
				botL := matrix[y+1][x-1]

				if !valid(matrix, y+1, x+1) {
					continue
				}
				botR := matrix[y+1][x+1]

				if botR != topL && botL != topR {
					ans++
				}
			}
		}
	}
	return ans
}

func valid(matrix [][]rune, y int, x int) bool {
	if 0 <= x && x < len(matrix[0]) && 0 <= y && y < len(matrix) && matrix[y][x] != 'A' && matrix[y][x] != 'X'{
		return true
	}
	return false
}

func part1(matrix [][]rune) int {
	ans := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == 'X' {
				ans += countForPoint(matrix, y, x)
			}
		}
	}
	return ans

}

type Vec2 struct {
	X int
	Y int
	D string
	L rune
}

func CheckVec(matrix [][]rune, x int, y int, d string) bool {
	vec := &Vec2{
		X: x,
		Y: y,
		D: d,
		L: 'X',
	}
	return vec.CheckDir(matrix)
}

func (p *Vec2) CheckDir(matrix [][]rune) bool {
	switch p.L {
	case 'X':
		p.L = 'M'
	case 'M':
		p.L = 'A'
	case 'A':
		p.L = 'S'
	case 'S':
		p.L = 'M'
	}

	switch p.D {
	case "N":
		p.Y++
	case "NE":
		p.Y++
		p.X++
	case "E":
		p.X++
	case "SE":
		p.X++
		p.Y--
	case "S":
		p.Y--
	case "SW":
		p.Y--
		p.X--
	case "W":
		p.X--
	case "NW":
		p.X--
		p.Y++
	}
	if 0 <= p.X && p.X < len(matrix[0]) && 0 <= p.Y && p.Y < len(matrix) && p.L == matrix[p.Y][p.X] {
		if p.L == 'S' {
			return true
		}
		return p.CheckDir(matrix)
	}
	return false
}

func countForPoint(matrix [][]rune, y int, x int) int {
	count := 0
	dirs := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	for _, d := range dirs {
		if CheckVec(matrix, x, y, d) {
			count++
		}
	}
	return count
}
