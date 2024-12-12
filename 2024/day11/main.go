package main

import (
	"fmt"
	"math"
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
	line := strings.Split(lines[0], " ")

	var nums []int
	for _, numS := range line {
		num, _ := strconv.Atoi(numS)
		nums = append(nums, num)
	}
	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(nums))
		fmt.Printf("PART 2: %v\n", part2(nums))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(nums))
		return
	}
	fmt.Printf("PART 1: %v", part1(nums))
}
func numDigits(i int) int {
	return int(math.Floor(math.Log10(float64(i))) + 1)
}

func part1(nums []int) int {
	for range 25 {
		var newNums []int
		for _, num := range nums {
			if num == 0 {
				newNums = append(newNums, 1)
				continue
			}

			numDigits := numDigits(num)
			if numDigits%2 == 0 {
				numS := strconv.Itoa(num)

				stoneStr := ""
				for i := range numDigits / 2 {
					stoneStr += string(numS[i])
				}
				stone, _ := strconv.Atoi(stoneStr)
				newNums = append(newNums, stone)

				stoneStr = ""
				for i := range numDigits / 2 {
					stoneStr += string(numS[i+numDigits/2])
				}
				stone, _ = strconv.Atoi(stoneStr)
				newNums = append(newNums, stone)
				continue
			}
			newNums = append(newNums, num*2024)
		}
		nums = newNums
	}
	return len(nums)
}

func calc(num int, n int, memo map[Key]int) int {
	if n == 0 {
		return 1
	}
	ans := 0
	k := Key{N: n, Num: num}
	n--
	val, ok := memo[k]
	if ok {
		return val
	}
	if num == 0 {
		ans += calc(1, n, memo)
		memo[k] = ans
		return ans
	}
	numDigits := numDigits(num)
	if numDigits%2 == 0 {
		numS := strconv.Itoa(num)

		stoneStr := ""
		for i := range numDigits / 2 {
			stoneStr += string(numS[i])
		}
		stone, _ := strconv.Atoi(stoneStr)
		ans += calc(stone, n, memo)
		stoneStr = ""
		for i := range numDigits / 2 {
			stoneStr += string(numS[i+numDigits/2])
		}
		stone, _ = strconv.Atoi(stoneStr)
		ans += calc(stone, n, memo)
		memo[k] = ans
		return ans
	}
	ans += calc(num*2024, n, memo)
	memo[k] = ans
	return ans
}

type Key struct {
	N   int
	Num int
}

func part2(nums []int) int {
	memo := make(map[Key]int)
	n := 74
	ans := 0
	for _, num := range nums {
		if num == 0 {
			ans += calc(1, n, memo)
			continue
		}

		numDigits := numDigits(num)
		if numDigits%2 == 0 {
			numS := strconv.Itoa(num)

			stoneStr := ""
			for i := range numDigits / 2 {
				stoneStr += string(numS[i])
			}
			stone, _ := strconv.Atoi(stoneStr)
			ans += calc(stone, n, memo)
			stoneStr = ""
			for i := range numDigits / 2 {
				stoneStr += string(numS[i+numDigits/2])
			}
			stone, _ = strconv.Atoi(stoneStr)
			ans += calc(stone, n, memo)
			continue
		}
		ans += calc(num*2024, n, memo)
	}
	return ans
}
