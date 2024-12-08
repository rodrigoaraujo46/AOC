package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		print(err.Error())
		return
	}

	r, err := regexp.Compile("[1-9]{2}\\|[1-9]{2}")
	if err != nil {
		print(err.Error())
	}
	pageOrderLines := r.FindAllString(string(content), -1)
	var pageOrder [][]int
	for _, po := range pageOrderLines {
		r := strings.Split(po, "|")
		var m []int
		for _, t := range r {
			y, _ := strconv.Atoi(t)
			m = append(m, y)
		}
		pageOrder = append(pageOrder, m)
	}
	r, err = regexp.Compile("\\d+(,\\d+)+")
	if err != nil {
		print(err.Error())
	}
	pageNumberLines := r.FindAllString(string(content), -1)
	var pageNumber [][]int
	for _, line := range pageNumberLines {
		var lineInt []int
		split := strings.Split(line, ",")
		for _, v := range split {
			num, _ := strconv.Atoi(v)
			lineInt = append(lineInt, num)
		}
		pageNumber = append(pageNumber, lineInt)
	}
	myMap := make(map[int][]int)
	for _, line := range pageOrder {
		myMap[line[0]] = append(myMap[line[0]], line[1])
	}
	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(pageNumber, myMap))
		fmt.Printf("PART 2: %v\n", part2(pageNumber, myMap))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(pageNumber, myMap))
		return
	}
	fmt.Printf("PART 1: %v", part1(pageNumber, myMap))
}

func getIncorrect(pageNumber [][]int, pageWrong *[][]int, incorrect *[]int, myMap map[int][]int) {
	for _, line := range pageNumber {
		for j := 0; j < len(line); j++ {
			possible := myMap[line[j]]
			if !is_Subset(line[j+1:], possible) {
				*incorrect = append(*incorrect, j)
				*pageWrong = append(*pageWrong, line)
				break
			}
		}
	}
}

func getIncorrectLine(line []int, myMap map[int][]int) int {
	for j := 0; j < len(line); j++ {
		possible := myMap[line[j]]
		if !is_Subset(line[j+1:], possible) {
			return j
		}
	}
	return -99
}

func SolveLine(myMap map[int][]int, wIdx int, line []int) int {
	ans := 0
	afterError := append([]int(nil), line[wIdx+1:]...)
	for i := 0; i < len(afterError); i++ {
		nowTesting := afterError[i]
		val, ok := myMap[nowTesting]
		if !ok {
			continue
		}

		subset := append([]int(nil), afterError[0:i]...)
		subset = append(subset, afterError[i+1:]...)
		subset = append(subset, line[wIdx])

		if is_Subset(subset, val) {
			help := line[wIdx]
			line[wIdx] = line[wIdx+1+i]
			line[wIdx+1+i] = help
			afterError = line[wIdx+1:]

			if valid(line, myMap) {
				ans += line[(len(line)-1)/2]
				return ans
			}
			ans += SolveLine(myMap, wIdx+1, line)
			return ans
		}
	}
	ans += SolveLine(myMap, wIdx+1, line)
	return ans
}

func part2(pageNumber [][]int, myMap map[int][]int) int {
	ans := 0
	pageWrong := make([][]int, 0)
	incorrect := make([]int, 0)
	getIncorrect(pageNumber, &pageWrong, &incorrect, myMap)

	for i, line := range pageWrong {
		wIdx := incorrect[i]
		ans += SolveLine(myMap, wIdx, line)
	}
	return ans
}

func valid(line []int, myMap map[int][]int) bool {
	for j := 0; j < len(line); j++ {
		possible := myMap[line[j]]
		if !is_Subset(line[j+1:], possible) {
			return false
		}
	}
	return true
}

func part1(pageNumber [][]int, myMap map[int][]int) int {
	ans := 0
	for _, line := range pageNumber {
		if valid(line, myMap) {
			ans += line[(len(line)-1)/2]
		}

	}
	return ans

}

func is_Subset(subset []int, superset []int) bool {
	checkset := make(map[int]bool)
	for _, element := range subset {
		checkset[element] = true
	}
	for _, value := range superset {
		if checkset[value] {
			delete(checkset, value)
		}
	}
	return len(checkset) == 0
}
