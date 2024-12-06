package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TO ANYONE HERE, PLEASE CLOSE YOUR EYES, I ALMOST BURNT MINE WHILE LOOKING AT THIS NUCLEAR WASTE
func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	var reports [][]int
	for _, string := range lines {
		line := strings.Split(string, " ")
		var report []int
		for _, num := range line {
			numI, _ := strconv.Atoi(num)
			report = append(report, numI)
		}
		reports = append(reports, report)
	}
	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(reports))
		fmt.Printf("PART 2: %v\n", part2(reports))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(reports))
		return
	}
	fmt.Printf("PART 1: %v", part1(reports))
}

func remove(slice []int, s int) []int {
	slice = append(slice[:s], slice[s+1:]...)[:len(slice)-1]
	return slice
}

func consider(report []int, indexes []int) bool {
	if indexes[1] > 1{
		indexes = append(indexes, indexes[1]-2)
	}
	for _,i := range indexes {
		opt := make([]int, len(report))
		copy(opt, report)
		opt = remove(opt, i)
		if checkCurrent(opt) {
			return true
		}
	}
	return false
}

func checkCurrent(report []int) bool {
	if report[0] == report[1] {
		return false
	}
	growing := report[0] < report[1]
	valid := true
	for i := 1; i < len(report); i++ {
		if !growing {
			if report[i-1] <= report[i] || report[i-1]-report[i] > 3 {
				valid = false
				break
			}
			continue
		}
		if report[i-1] >= report[i] || report[i-1]-report[i] < -3 {
			valid = false
			break
		}
	}
	return valid
}

func part2(reports [][]int) int {
	ans := 0
	for _, report := range reports {
		valid := true
		if report[0] == report[1] {
			if consider(report, []int{0,1}) {
				ans++
				continue
			}
			continue
		}
		growing := report[0] < report[1]
		bail := true
		for i := 1; i < len(report); i++ {
			if !growing {
				if report[i-1] <= report[i] || report[i-1]-report[i] > 3 {
					if !bail {
						valid = false
						break
					}
					if consider(report, []int{i-1,i}) {
						break
					}
					valid = false
					break
				}
				continue
			}
			if report[i-1] >= report[i] || report[i-1]-report[i] < -3 {
				if !bail {
					valid = false
					break
				}

				if consider(report, []int{i-1,i}) {
					break
				}
				valid = false
				break
			}
		}
		if valid {
			ans++
		}
	}
	return ans
}

func part1(reports [][]int) int {
	ans := 0
	for _, report := range reports {
		if report[0] == report[1] {
			continue
		}
		growing := report[0] < report[1]
		valid := true
		for i := 1; i < len(report); i++ {
			if !growing {
				if report[i-1] <= report[i] || report[i-1]-report[i] > 3 {
					valid = false
					break
				}
				continue
			}
			if report[i-1] >= report[i] || report[i-1]-report[i] < -3 {
				valid = false
				break
			}
		}
		if valid {
			ans++
		}
	}
	return ans
}
