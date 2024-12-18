package main

import (
	"fmt"
	"math"
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

	var maze [][]string
	for _, line := range lines {
		chars := strings.Split(line, "")
		maze = append(maze, chars)
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(maze))
		fmt.Printf("PART 2: %v\n", part2(maze))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(maze))
		return
	}
	fmt.Printf("PART 1: %v", part1(maze))
}

type Reindeer struct {
	X      int
	Y      int
	Dir    string
	Points int
}

type ReindeerKey struct {
	X   int
	Y   int
	Dir string
}

func Key(r Reindeer) ReindeerKey {
	return ReindeerKey{
		X:   r.X,
		Y:   r.Y,
		Dir: r.Dir,
	}
}

type Position struct {
	X int
	Y int
}

func (c *Reindeer) Turn(counter bool) {
	c.Points += 1000
	if c.Dir == "E" {
		if counter {
			c.Dir = "N"
		} else {
			c.Dir = "S"
		}
		return
	}
	if c.Dir == "S" {
		if counter {
			c.Dir = "E"
		} else {
			c.Dir = "W"
		}
		return
	}
	if c.Dir == "W" {
		if counter {
			c.Dir = "S"
		} else {
			c.Dir = "N"
		}
		return
	}
	if counter {
		c.Dir = "W"
	} else {
		c.Dir = "E"
	}
	return
}

type MazeError struct{}

func (e MazeError) Error() string {
	return fmt.Sprintf("That's a wall!")
}

func getScore(m map[ReindeerKey]int, key ReindeerKey) int {
	if value, exists := m[key]; exists {
		return value
	}
	return math.MaxInt
}

func (c *Reindeer) Forward(visited map[ReindeerKey]int, maze [][]string) error {
	if c.Dir == "E" {
		c.X += 1
		c.Points += 1
		if maze[c.Y][c.X] == "#" || getScore(visited, Key(*c)) <= c.Points {
			return MazeError{}
		}
		visited[Key(*c)] = c.Points
		return nil
	}
	if c.Dir == "S" {
		c.Y += 1
		c.Points += 1
		if maze[c.Y][c.X] == "#" || getScore(visited, Key(*c)) <= c.Points {
			return MazeError{}
		}
		visited[Key(*c)] = c.Points
		return nil
	}
	if c.Dir == "W" {
		c.X -= 1
		c.Points += 1
		if maze[c.Y][c.X] == "#" || getScore(visited, Key(*c)) <= c.Points {
			return MazeError{}
		}
		visited[Key(*c)] = c.Points
		return nil
	}
	c.Y -= 1
	c.Points += 1
	if maze[c.Y][c.X] == "#" || getScore(visited, Key(*c)) <= c.Points {
		return MazeError{}
	}
	visited[Key(*c)] = c.Points
	return nil
}

func (c *Reindeer) Copy() *Reindeer {
	copy := *c
	return &copy
}

func solveMaze(visited map[ReindeerKey]int, maze [][]string, c *Reindeer) int {
	fw, rw, lw, bw := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	if maze[c.Y][c.X] == "E" {
		return c.Points
	}
	f := c.Copy()
	if f.Forward(visited, maze) == nil {
		fw = solveMaze(visited, maze, f)
	}
	r := c.Copy()
	r.Turn(false)
	if r.Forward(visited, maze) == nil {
		rw = solveMaze(visited, maze, r)
	}

	l := c.Copy()
	l.Turn(true)
	if l.Forward(visited, maze) == nil {
		lw = solveMaze(visited, maze, l)
	}
	b := c.Copy()
	b.Turn(true)
	b.Turn(true)
	if b.Forward(visited, maze) == nil {
		bw = solveMaze(visited, maze, b)
	}

	return min(fw, rw, lw, bw)
}

func (c *Reindeer) Forward2(visited map[ReindeerKey]int, maze [][]string, best int) error {
	if c.Dir == "E" {
		c.X += 1
		c.Points += 1
		if maze[c.Y][c.X] == "#" || c.Points > best || getScore(visited, Key(*c)) < c.Points {
			return MazeError{}
		}
		visited[Key(*c)] = c.Points
		return nil
	}
	if c.Dir == "S" {
		c.Y += 1
		c.Points += 1
		if maze[c.Y][c.X] == "#" || c.Points > best || getScore(visited, Key(*c)) < c.Points {
			return MazeError{}
		}
		visited[Key(*c)] = c.Points
		return nil
	}
	if c.Dir == "W" {
		c.X -= 1
		c.Points += 1
		if maze[c.Y][c.X] == "#" || c.Points > best || getScore(visited, Key(*c)) < c.Points {
			return MazeError{}
		}
		visited[Key(*c)] = c.Points
		return nil
	}
	c.Y -= 1
	c.Points += 1
	if maze[c.Y][c.X] == "#" || c.Points > best || getScore(visited, Key(*c)) < c.Points {
		return MazeError{}
	}
	visited[Key(*c)] = c.Points
	return nil
}

func GetPaths(paths *[][]Position, path []Position, visited map[ReindeerKey]int, maze [][]string, c *Reindeer, best int) {
	path = append(path, Position{X: c.X, Y: c.Y})

	if maze[c.Y][c.X] == "E" && c.Points == best {
		*paths = append(*paths, append([]Position(nil), path...))
	}

	f := c.Copy()
	if f.Forward2(visited, maze, best) == nil {
		GetPaths(paths, path, visited, maze, f, best)
	}
	r := c.Copy()
	r.Turn(false)
	if r.Forward2(visited, maze, best) == nil {
		GetPaths(paths, path, visited, maze, r, best)
	}

	l := c.Copy()
	l.Turn(true)
	if l.Forward2(visited, maze, best) == nil {
		GetPaths(paths, path, visited, maze, l, best)
	}
	b := c.Copy()
	b.Turn(true)
	b.Turn(true)
	if b.Forward2(visited, maze, best) == nil {
		GetPaths(paths, path, visited, maze, b, best)
	}
}

func part1(maze [][]string) int {
	c := &Reindeer{}

Outer:
	for i, line := range maze {
		for j, char := range line {
			if char == "S" {
				c.X = j
				c.Y = i
				c.Dir = "E"
				break Outer
			}
		}
	}
	reindeerMap := make(map[ReindeerKey]int)
	reindeerMap[Key(*c)] = c.Points
	return solveMaze(reindeerMap, maze, c)
}

func part2(maze [][]string) int {
	bestPathSize := part1(maze)
	c := &Reindeer{}

Outer:
	for i, line := range maze {
		for j, char := range line {
			if char == "S" {
				c.X = j
				c.Y = i
				c.Dir = "E"
				break Outer
			}
		}
	}
	reindeerMap := make(map[ReindeerKey]int)
	var paths [][]Position
	GetPaths(&paths, []Position{}, reindeerMap, maze, c, bestPathSize)
	pathSet := make(map[Position]bool)
	for _, path := range paths {
		for _, pos := range path {
			pathSet[pos] = true
		}
	}
	return len(pathSet)
}
