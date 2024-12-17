package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	X int
	Y int
}

type Robot struct {
	Position *Vector
	Velocity Vector
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		print(err.Error())
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	var robots []*Robot
	for _, line := range lines {
		positionLine, velocityLine := strings.Split(line, " ")[0], strings.Split(line, " ")[1]

		afterEquals := strings.Split(positionLine, "=")[1]
		xString, yString := strings.Split(afterEquals, ",")[0], strings.Split(afterEquals, ",")[1]
		x, _ := strconv.Atoi(xString)
		y, _ := strconv.Atoi(yString)
		p := &Vector{
			X: x,
			Y: y,
		}

		afterEquals = strings.Split(velocityLine, "=")[1]
		xString, yString = strings.Split(afterEquals, ",")[0], strings.Split(afterEquals, ",")[1]
		x, _ = strconv.Atoi(xString)
		y, _ = strconv.Atoi(yString)
		v := Vector{
			X: x,
			Y: y,
		}
		r := &Robot{
			Position: p,
			Velocity: v,
		}
		robots = append(robots, r)
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(robots))
		fmt.Printf("PART 2: %v\n", part2(robots))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(robots))
		return
	}
	fmt.Printf("PART 1: %v", part1(robots))
}

func part1(robots []*Robot) int {
	width := 101
	height := 103
	n1, n2, n3, n4 := 0, 0, 0, 0
	for _, robot := range robots {
		robot.Position.X = ((robot.Position.X+100*robot.Velocity.X)%width + width) % width
		robot.Position.Y = ((robot.Position.Y+100*robot.Velocity.Y)%height + height) % height
		if robot.Position.X > width/2 && robot.Position.Y < height/2 {
			n1++
			continue
		}
		if robot.Position.X < width/2 && robot.Position.Y < height/2 {
			n2++
			continue
		}
		if robot.Position.X < width/2 && robot.Position.Y > height/2 {
			n3++
			continue
		}
		if robot.Position.X > width/2 && robot.Position.Y > height/2 {
			n4++
		}
	}
	return n1 * n2 * n3 * n4
}

func part2(robots []*Robot) int {
	width := 101
	height := 103
	best_it := -1
	min_safety := 999999999999999999
	for i := range width * height {
		n1, n2, n3, n4 := 0, 0, 0, 0
		for _, robot := range robots {
			robot.Position.X = ((robot.Position.X+robot.Velocity.X)%width + width) % width
			robot.Position.Y = ((robot.Position.Y+robot.Velocity.Y)%height + height) % height
			if robot.Position.X > width/2 && robot.Position.Y < height/2 {
				n1++
				continue
			}
			if robot.Position.X < width/2 && robot.Position.Y < height/2 {
				n2++
				continue
			}
			if robot.Position.X < width/2 && robot.Position.Y > height/2 {
				n3++
				continue
			}
			if robot.Position.X > width/2 && robot.Position.Y > height/2 {
				n4++
			}
		}
		safety := n1 * n2 * n3 * n4
		if safety < min_safety {
			min_safety = safety
			best_it = i
		}
	}
	return best_it + 1
}
