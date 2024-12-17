package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		print(err.Error())
		return
	}

	lines := strings.Split(string(content), "\n")

	var operations [][]int
	var operaion []int
	for i, line := range lines {
		if line == "" {
			operations = append(operations, operaion)
			operaion = []int{}
			continue
		}
		if i%4 == 0 || i%4 == 1 {
			afterPoints := strings.Split(line, ":")[1]
			beforeComma, afterComma := strings.Split(afterPoints, ",")[0], strings.Split(afterPoints, ",")[1]
			xString, yString := strings.Split(beforeComma, "+")[1], strings.Split(afterComma, "+")[1]
			x, _ := strconv.Atoi(xString)
			y, _ := strconv.Atoi(yString)
			operaion = append(operaion, x, y)
			continue
		}
		if i%4 == 2 {
			afterPoints := strings.Split(line, ":")[1]
			beforeComma, afterComma := strings.Split(afterPoints, ",")[0], strings.Split(afterPoints, ",")[1]
			xString, yString := strings.Split(beforeComma, "=")[1], strings.Split(afterComma, "=")[1]
			x, _ := strconv.Atoi(xString)
			y, _ := strconv.Atoi(yString)
			operaion = append(operaion, x, y)
		}
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(operations))
		fmt.Printf("PART 2: %v\n", part2(operations))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(operations))
		return
	}
	fmt.Printf("PART 1: %v", part1(operations))
}

func part1(operations [][]int) int {
	ans := 0
	for i := range operations {
		ak1, ak2 := operations[i][0], operations[i][1]
		bk1, bk2 := operations[i][2], operations[i][3]
		wantX := operations[i][4]
		wantY := operations[i][5]
		found := false
		for b := 100; b >= 0; b-- {
			for a := range 100 {
				if wantX == ak1*a+bk1*b && wantY == ak2*a+bk2*b {
					ans += 3*a + b
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return ans
}

func solveSystem(A1, B1, X, A2, B2, Y int) (int, int) {

	if (X*B2 - Y*B1) % (A1*B2 - A2*B1) != 0{
		return 0,0
	}
	a := (X*B2 - Y*B1) / (A1*B2 - A2*B1)

	if (Y-A2*a)%B2 != 0 {
		return 0, 0
	}
	b := (Y - A2*a) / B2

	return a, b
}

func part2(operations [][]int) int {
	base := 10_000_000_000_000
	ans := 0
	for i := range operations {
		ak1, ak2 := operations[i][0], operations[i][1]
		bk1, bk2 := operations[i][2], operations[i][3]
		wantX := base + operations[i][4]
		wantY := base + operations[i][5]
		a, b:= solveSystem(ak1, bk1, wantX, ak2, bk2, wantY)
		ans += 3*a + b
	}
	return ans
}
