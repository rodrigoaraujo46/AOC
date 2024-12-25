package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func (p point) getManhattan(v point) (int, int) {
	return v.x - p.x, v.y - p.y
}

type code []string
type codes []code

type numpad [4][3]string

func (np numpad) getPoint(str string) point {
	for i := range np {
		for j := range np[i] {
			if np[i][j] == str {
				return point{x: j, y: i}
			}
		}
	}
	return point{}
}

func (dp dirpad) getPoint(str string) point {
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] == str {
				return point{x: j, y: i}
			}
		}
	}
	return point{}
}

func (c *code) getDirNumpad(numpad numpad) code {
	current := point{x: 2, y: 3}
	path := make(code, 0)

	for _, char := range *c {
		currentCode := make(code, 0)
		end := numpad.getPoint(char)

		if current.x > 0 && current.y+1 == end.y {
			current.y++
			currentCode = append(currentCode, "v")
		}
		if current.x == 0 && end.x != 0 {
			current.x++
			currentCode = append(currentCode, ">")
		} else if current.y == 3 && end.y != 3 {
			if current.x == 2 && end.x == 1 {
				current.x--
				currentCode = append(currentCode, "<")
			}
			current.y--
			currentCode = append(currentCode, "^")

		}

		x, y := current.getManhattan(end)
		if len(currentCode) > 0 && currentCode[len(currentCode)-1] == "^" {
			currentCode.appendDirections(y, "^", "v")
			currentCode.appendDirections(x, "<", ">")
		} else {
			currentCode.appendDirections(x, "<", ">")
			currentCode.appendDirections(y, "^", "v")
		}
		currentCode = append(currentCode, "A")
		path = append(path, currentCode...)
		current = end
	}
	return path
}

func (code *code) appendDirections(count int, negativeChar, positiveChar string) {
	for i := 0; i < int(math.Abs(float64(count))); i++ {
		if count < 0 {
			*code = append(*code, negativeChar)
			continue
		}
		*code = append(*code, positiveChar)

	}
}

func (c *code) getDirDirpad(dirpad dirpad) code {
	current := point{x: 2, y: 0}
	path := make(code, 0)

	for _, char := range *c {
		currentCode := make(code, 0)
		end := dirpad.getPoint(char)

		if current.x == 0 && end.y != 1 {
			current.x++
			currentCode = append(currentCode, ">")
		} else if current.y == 0 && end.y != 0 {
			if end.x == current.x-1 {
				current.x--
				currentCode = append(currentCode, "<")
			}
			current.y++
			currentCode = append(currentCode, "v")

		}

		x, y := current.getManhattan(end)
		if len(currentCode) > 0 && currentCode[len(currentCode)-1] == "v" {
			currentCode.appendDirections(y, "^", "v")
			currentCode.appendDirections(x, "<", ">")
		} else {
			currentCode.appendDirections(x, "<", ">")
			currentCode.appendDirections(y, "^", "v")
		}
		currentCode = append(currentCode, "A")
		path = append(path, currentCode...)
		current = end
	}
	return path
}

type dirpad [2][3]string

type buttonPair struct {
	first  string
	second string
}

var paths = map[buttonPair]string{
	{"A", "0"}: "<A",
	{"0", "A"}: ">A",
	{"A", "1"}: "^<<A",
	{"1", "A"}: ">>vA",
	{"A", "2"}: "<^A",
	{"2", "A"}: "v>A",
	{"A", "3"}: "^A",
	{"3", "A"}: "vA",
	{"A", "4"}: "^^<<A",
	{"4", "A"}: ">>vvA",
	{"A", "5"}: "<^^A",
	{"5", "A"}: "vv>A",
	{"A", "6"}: "^^A",
	{"6", "A"}: "vvA",
	{"A", "7"}: "^^^<<A",
	{"7", "A"}: ">>vvvA",
	{"A", "8"}: "<^^^A",
	{"8", "A"}: "vvv>A",
	{"A", "9"}: "^^^A",
	{"9", "A"}: "vvvA",
	{"0", "1"}: "^<A",
	{"1", "0"}: ">vA",
	{"0", "2"}: "^A",
	{"2", "0"}: "vA",
	{"0", "3"}: "^>A",
	{"3", "0"}: "<vA",
	{"0", "4"}: "^<^A",
	{"4", "0"}: ">vvA",
	{"0", "5"}: "^^A",
	{"5", "0"}: "vvA",
	{"0", "6"}: "^^>A",
	{"6", "0"}: "<vvA",
	{"0", "7"}: "^^^<A",
	{"7", "0"}: ">vvvA",
	{"0", "8"}: "^^^A",
	{"8", "0"}: "vvvA",
	{"0", "9"}: "^^^>A",
	{"9", "0"}: "<vvvA",
	{"1", "2"}: ">A",
	{"2", "1"}: "<A",
	{"1", "3"}: ">>A",
	{"3", "1"}: "<<A",
	{"1", "4"}: "^A",
	{"4", "1"}: "vA",
	{"1", "5"}: "^>A",
	{"5", "1"}: "<vA",
	{"1", "6"}: "^>>A",
	{"6", "1"}: "<<vA",
	{"1", "7"}: "^^A",
	{"7", "1"}: "vvA",
	{"1", "8"}: "^^>A",
	{"8", "1"}: "<vvA",
	{"1", "9"}: "^^>>A",
	{"9", "1"}: "<<vvA",
	{"2", "3"}: ">A",
	{"3", "2"}: "<A",
	{"2", "4"}: "<^A",
	{"4", "2"}: "v>A",
	{"2", "5"}: "^A",
	{"5", "2"}: "vA",
	{"2", "6"}: "^>A",
	{"6", "2"}: "<vA",
	{"2", "7"}: "<^^A",
	{"7", "2"}: "vv>A",
	{"2", "8"}: "^^A",
	{"8", "2"}: "vvA",
	{"2", "9"}: "^^>A",
	{"9", "2"}: "<vvA",
	{"3", "4"}: "<<^A",
	{"4", "3"}: "v>>A",
	{"3", "5"}: "<^A",
	{"5", "3"}: "v>A",
	{"3", "6"}: "^A",
	{"6", "3"}: "vA",
	{"3", "7"}: "<<^^A",
	{"7", "3"}: "vv>>A",
	{"3", "8"}: "<^^A",
	{"8", "3"}: "vv>A",
	{"3", "9"}: "^^A",
	{"9", "3"}: "vvA",
	{"4", "5"}: ">A",
	{"5", "4"}: "<A",
	{"4", "6"}: ">>A",
	{"6", "4"}: "<<A",
	{"4", "7"}: "^A",
	{"7", "4"}: "vA",
	{"4", "8"}: "^>A",
	{"8", "4"}: "<vA",
	{"4", "9"}: "^>>A",
	{"9", "4"}: "<<vA",
	{"5", "6"}: ">A",
	{"6", "5"}: "<A",
	{"5", "7"}: "<^A",
	{"7", "5"}: "v>A",
	{"5", "8"}: "^A",
	{"8", "5"}: "vA",
	{"5", "9"}: "^>A",
	{"9", "5"}: "<vA",
	{"6", "7"}: "<<^A",
	{"7", "6"}: "v>>A",
	{"6", "8"}: "<^A",
	{"8", "6"}: "v>A",
	{"6", "9"}: "^A",
	{"9", "6"}: "vA",
	{"7", "8"}: ">A",
	{"8", "7"}: "<A",
	{"7", "9"}: ">>A",
	{"9", "7"}: "<<A",
	{"8", "9"}: ">A",
	{"9", "8"}: "<A",
	{"<", "^"}: ">^A",
	{"^", "<"}: "v<A",
	{"<", "v"}: ">A",
	{"v", "<"}: "<A",
	{"<", ">"}: ">>A",
	{">", "<"}: "<<A",
	{"<", "A"}: ">>^A",
	{"A", "<"}: "v<<A",
	{"^", "v"}: "vA",
	{"v", "^"}: "^A",
	{"^", ">"}: "v>A",
	{">", "^"}: "<^A",
	{"^", "A"}: ">A",
	{"A", "^"}: "<A",
	{"v", ">"}: ">A",
	{">", "v"}: "<A",
	{"v", "A"}: "^>A",
	{"A", "v"}: "<vA",
	{">", "A"}: "^A",
	{"A", ">"}: "vA",
	{"A", "A"}: "A",
	{"0", "0"}: "A",
	{"1", "1"}: "A",
	{"2", "2"}: "A",
	{"3", "3"}: "A",
	{"4", "4"}: "A",
	{"5", "5"}: "A",
	{"6", "6"}: "A",
	{"7", "7"}: "A",
	{"8", "8"}: "A",
	{"9", "9"}: "A",
	{"<", "<"}: "A",
	{"^", "^"}: "A",
	{"v", "v"}: "A",
	{">", ">"}: "A",
}

func (c *code) getDirNumpad2() code {
	path := make(code, 0)
	for j := 1; j < len(*c); j++ {
		i := j - 1
		start := (*c)[i]
		end := (*c)[j]
		path = append(path, paths[buttonPair{first: start, second: end}])
	}
	return path
}

func recursive(memo map[string]int, a, b string, depth int, dirpad dirpad, numpad numpad) int {
	key := a + b + string(depth)
	if val, ok := memo[key]; ok {
		return val
	}
	if depth == 1 {
		return len(paths[buttonPair{first: a, second: b}])
	}

	seq := code{"A"}
	seq = append(seq, strings.Split(paths[buttonPair{first: a, second: b}], "")...)

	length := 0
	for j := 1; j < len(seq); j++ {
		i := j - 1
		length += recursive(memo, seq[i], seq[j], depth-1, dirpad, numpad)
	}
	memo[key] = length
	return length
}

func part2(codes codes, numpad numpad, dirpad dirpad) int {
	ans := 0
	memo := make(map[string]int)
	for _, code_ := range codes {
		code_ = append(code{"A"}, code_...)
		inputs := strings.Split(strings.Join(append(code{"A"}, code_.getDirNumpad2()...), ""), "")

		length := 0
		for j := 1; j < len(inputs); j++ {
			i := j - 1
			length += recursive(memo, inputs[i], inputs[j], 25, dirpad, numpad)
		}
		numerical := extractNumber(strings.Join(code_, ""))
		complexity := length * numerical
		ans += complexity

	}
	return ans
}

func extractNumber(s string) int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindString(s)
	if matches != "" {
		number, err := strconv.Atoi(matches)
		if err == nil {
			return number
		}
	}
	return 0
}

func part1(codes codes, numpad numpad, dirpad dirpad) int {
	ans := 0

	for i := range codes {
		code := codes[i].getDirNumpad(numpad)
		for range 2 {
			code = code.getDirDirpad(dirpad)
		}

		numerical := extractNumber(strings.Join(codes[i], ""))
		complexity := len(code) * numerical
		ans += complexity

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

	codes := make(codes, len(lines))

	for i, line := range lines {
		code := strings.Split(line, "")
		codes[i] = code
	}

	numpad := numpad{{"7", "8", "9"}, {"4", "5", "6"}, {"1", "2", "3"}, {"#", "0", "A"}}
	dirpad := dirpad{{"#", "^", "A"}, {"<", "v", ">"}}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(codes, numpad, dirpad))
		fmt.Printf("PART 2: %v\n", part2(codes, numpad, dirpad))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(codes, numpad, dirpad))
		return
	}
	fmt.Printf("PART 1: %v", part1(codes, numpad, dirpad))
}
