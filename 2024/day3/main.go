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

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(content))
		fmt.Printf("PART 2: %v\n", part2(content))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(content))
		return
	}
	fmt.Printf("PART 1: %v", part1(content))
}

func part2(content []byte) int {
	r, err := regexp.Compile("(don't\\(\\))|(do\\(\\))|(mul\\([0-9]{1,3},[0-9]{1,3}\\))")
	if err != nil {
		print(err.Error())
		return -1
	}

	matches := r.FindAllString(string(content), -1)
	r, err = regexp.Compile("[0-9]{1,3},[0-9]{1,3}")
	if err != nil {
		print(err.Error())
		return -1
	}

	ans := 0
	prohibit := false
	for _, match := range matches {
		if match == "do()"{
			prohibit = false
			continue
		}
		if match == "don't()"{
			prohibit = true
			continue
		}
		if prohibit{
			continue
		}
		nums := strings.Split(r.FindString(match), ",")
		int1, err := strconv.Atoi(nums[0])
		if err != nil {
			print(err.Error())
			return -1
		}
		int2, err := strconv.Atoi(nums[1])
		if err != nil {
			print(err.Error())
			return -1
		}
		ans += int1 * int2
	}
	return ans
}

func part1(content []byte) int {
	r, err := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	if err != nil {
		print(err.Error())
		return -1
	}

	matches := r.FindAllString(string(content), -1)
	r, err = regexp.Compile("[0-9]{1,3},[0-9]{1,3}")
	if err != nil {
		print(err.Error())
		return -1
	}
	ans := 0
	for _, mul := range matches {
		nums := strings.Split(r.FindString(mul), ",")
		int1, err := strconv.Atoi(nums[0])
		if err != nil {
			print(err.Error())
			return -1
		}
		int2, err := strconv.Atoi(nums[1])
		if err != nil {
			print(err.Error())
			return -1
		}
		ans += int1 * int2
	}
	return ans
}
