package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func newPoint(x, y int) point {
	return point{x: x, y: y}
}

type points []point

type cpu [][]string

func (cpu cpu) inBounds(point point) bool {
	return point.x >= 0 && point.x < len(cpu[0]) && point.y >= 0 && point.y < len(cpu)
}

func (cpu cpu) String() string {
	var sb strings.Builder
	for i := 0; i < len(cpu); i++ {
		for j := 0; j < len(cpu[i]); j++ {
			sb.WriteString(cpu[i][j])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (cpu cpu) getValidNeighbours(visited map[point]int, point point) points {
	directions := [4][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	neighbours := make(points, 0)
	for _, direction := range directions {
		neighbour := newPoint(point.x+direction[0], point.y+direction[1])

		if cpu.inBounds(neighbour) && cpu[neighbour.y][neighbour.x] != "#" {
			if _, visited := visited[neighbour]; !visited {
				neighbours = append(neighbours, neighbour)
			}
		}
	}
	return neighbours
}

func (cpu cpu) shortestPath(start, end point) map[point]int {
	queue := points{start}

	visited := make(map[point]int)
	visited[start] = 0

	for len(queue) != 0 {
		myByte := queue[0]
		queue = queue[1:]
		if myByte == end {
			return visited
		}

		neighbours := cpu.getValidNeighbours(visited, myByte)
		if len(neighbours) == 0 {
			continue
		}
		for _, neighbour := range neighbours {
			visited[neighbour] = visited[myByte] + 1
		}
		queue = append(queue, neighbours...)
	}

	return nil
}

func part2(cpu cpu) int {
	ans := 0
	var start point
	var end point
	for y, line := range cpu {
		for x, char := range line {
			if char == "S" {
				start = newPoint(x, y)
			}
			if char == "E" {
				end = newPoint(x, y)
			}
		}
	}

	distanceMap := cpu.shortestPath(start, end)

	points := make(points, 0, len(distanceMap))

	for key := range distanceMap {
		points = append(points, key)
	}

	for i := range points {
		for j := i; j < len(points); j++ {
			manhattan := manhattan(points[i], points[j])
			if manhattan <= 20 {
				distance := distanceMap[points[j]] - distanceMap[points[i]]
				if distance < 0 {
					distance = -distance
				}

				cheatD := distance - manhattan
				if cheatD >= 100 {
					ans++
				}
			}

		}
	}
	return ans
}

func manhattan(a point, b point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func part1(cpu cpu) int {
	ans := 0
	var start point
	var end point
	for y, line := range cpu {
		for x, char := range line {
			if char == "S" {
				start = newPoint(x, y)
			}
			if char == "E" {
				end = newPoint(x, y)
			}
		}
	}

	distanceMap := cpu.shortestPath(start, end)

	points := make(points, 0, len(distanceMap))

	for key := range distanceMap {
		points = append(points, key)
	}

	for i := range points {
		for j := i; j < len(points); j++ {
			manhattan := manhattan(points[i], points[j])
			if manhattan <= 2 {
				distance := distanceMap[points[j]] - distanceMap[points[i]]
				if distance < 0 {
					distance = -distance
				}

				cheatD := distance - manhattan
				if cheatD >= 100 {
					ans++
				}
			}

		}
	}
	return ans
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	cpu := make(cpu, len(lines))

	for i, line := range lines {
		chars := strings.Split(line, "")
		cpu[i] = chars
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(cpu))
		fmt.Printf("PART 2: %v\n", part2(cpu))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(cpu))
		return
	}
	fmt.Printf("PART 1: %v", part1(cpu))
}
