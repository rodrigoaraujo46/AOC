package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type myByte struct {
	x int
	y int
}

func newByte(x, y int) myByte {
	return myByte{x: x, y: y}
}

type myBytes []myByte

const memSize = 71

type memory [memSize][memSize]bool

func (memory *memory) writeBytes(myBytes myBytes, start, size int) {
	if start+size > len(myBytes) {
		panic("Memory overflow.")
	}

	for i := start; i < start+size; i++ {
		if !memory.inBounds(myBytes[i]) {
			panic("Byte out of memory bounds.")
		}
		memory[myBytes[i].y][myBytes[i].x] = true
	}
}

func (memory memory) inBounds(myByte myByte) bool {
	return myByte.x >= 0 && myByte.x < len(memory) && myByte.y >= 0 && myByte.y < len(memory)
}
func (memory memory) String() string {
	var sb strings.Builder
	for i := 0; i < len(memory); i++ {
		for j := 0; j < len(memory[i]); j++ {
			if memory[i][j] {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (memory memory) getValidNeighbours(visited map[myByte]int, myByte myByte) myBytes {
	directions := [4][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	neighbours := make(myBytes, 0)
	for _, direction := range directions {
		neighbour := newByte(myByte.x+direction[0], myByte.y+direction[1])

		if memory.inBounds(neighbour) && !memory[neighbour.y][neighbour.x] {
			if _, visited := visited[neighbour]; !visited {
				neighbours = append(neighbours, neighbour)
			}
		}
	}
	return neighbours
}

func (memory memory) shortestPath() int {
	end := newByte(memSize-1, memSize-1)
	start := newByte(0, 0)

	queue := myBytes{start}

	visited := make(map[myByte]int)
	visited[start] = 0

	for len(queue) != 0 {
		myByte := queue[0]
		queue = queue[1:]
		if myByte == end {
			return visited[myByte]
		}

		neighbours := memory.getValidNeighbours(visited, myByte)
		if len(neighbours) == 0 {
			continue
		}
		for _, neighbour := range neighbours {
			visited[neighbour] = visited[myByte] + 1
		}
		queue = append(queue, neighbours...)
	}

	return -1
}

func part2(myBytes myBytes) myByte {
	mem := new(memory)

	spare := new(memory)

	lower := 0
	upper := len(myBytes) - 1
	for lower < upper {
		pivot := (lower + upper) / 2
		*spare = *mem
		spare.writeBytes(myBytes, 0, pivot+1)

		if spare.shortestPath() == -1 {
			upper = pivot
			continue
		}
		lower = pivot + 1
	}
	return myBytes[lower]
}

func part1(myBytes myBytes) int {
	memory := new(memory)
	memory.writeBytes(myBytes, 0, 1024)

	return memory.shortestPath()
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	myBytes := make(myBytes, len(lines))

	for i := range lines {
		pointS := strings.Split(lines[i], ",")

		x, err := strconv.Atoi(pointS[0])
		if err != nil {
			panic(err.Error())
		}

		y, err := strconv.Atoi(pointS[1])
		if err != nil {
			panic(err.Error())
		}

		myBytes[i] = newByte(x, y)
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(myBytes))
		fmt.Printf("PART 2: %v\n", part2(myBytes))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(myBytes))
		return
	}
	fmt.Printf("PART 1: %v", part1(myBytes))
}
