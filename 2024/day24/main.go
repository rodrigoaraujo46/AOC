package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type GateAndWires struct {
	wireA string
	gate  string
	wireB string
	wireC string
}

func part2(gatesAndWires []GateAndWires) string {
	mymap := make(map[string]GateAndWires)
	for i := range gatesAndWires {
		mymap[gatesAndWires[i].wireC] = gatesAndWires[i]
	}

	swaps := make([]string, 0)
	for range 4 {
		baseline := progress(mymap)
		for x := range mymap {
			found := false
			for y := range mymap {
				if x == y {
					continue
				}

				aux := mymap[x]
				mymap[x], mymap[y] = mymap[y], aux

				if progress(mymap) > baseline {
					found = true
					swaps = append(swaps, x, y)
					break
				}

				aux = mymap[x]
				mymap[x], mymap[y] = mymap[y], aux
			}
			if !found {
				continue
			}
			break
		}
	}
	sort.Strings(swaps)
	return strings.Join(swaps, ",")
}

func progress(mymap map[string]GateAndWires) int {
	i := 0
	for ; ; i++ {
		if !verify(mymap, i) {
			break
		}
	}
	return i
}

func verify(mymap map[string]GateAndWires, num int) bool {
	return verifyZ(mymap, makeWire("z", num), num)
}

func verifyZ(mymap map[string]GateAndWires, wire string, num int) bool {
	if _, ok := mymap[wire]; !ok {
		return false
	}
	gw := mymap[wire]
	if gw.gate != "XOR" {
		return false
	}
	if num == 0 {
		return gw.wireA == "x00" && gw.wireB == "y00" || gw.wireA == "y00" && gw.wireB == "x00"
	}
	return verifyInter(mymap, gw.wireA, num) && verifyCarry(mymap, gw.wireB, num) || verifyInter(mymap, gw.wireB, num) && verifyCarry(mymap, gw.wireA, num)
}

func verifyInter(mymap map[string]GateAndWires, wire string, num int) bool {
	gw := mymap[wire]
	if gw.gate != "XOR" {
		return false
	}
	return gw.wireA == makeWire("x", num) && gw.wireB == makeWire("y", num) || gw.wireA == makeWire("y", num) && gw.wireB == makeWire("x", num)
}

func verifyCarry(mymap map[string]GateAndWires, wire string, num int) bool {
	gw := mymap[wire]
	if num == 1 {
		if gw.gate != "AND" {
			return false
		}
		return gw.wireA == "x00" && gw.wireB == "y00" || gw.wireA == "y00" && gw.wireB == "x00"
	}
	if gw.gate != "OR" {
		return false
	}
	return verifyDirect(mymap, gw.wireA, num-1) && verifyRecarry(mymap, gw.wireB, num-1) || verifyDirect(mymap, gw.wireB, num-1) && verifyRecarry(mymap, gw.wireA, num-1)
}

func verifyDirect(mymap map[string]GateAndWires, wire string, num int) bool {
	gw := mymap[wire]
	if gw.gate != "AND" {
		return false
	}
	return gw.wireA == makeWire("x", num) && gw.wireB == makeWire("y", num) || gw.wireA == makeWire("y", num) && gw.wireB == makeWire("x", num)
}

func verifyRecarry(mymap map[string]GateAndWires, wire string, num int) bool {
	gw := mymap[wire]
	if gw.gate != "AND" {
		return false
	}
	return verifyInter(mymap, gw.wireA, num) && verifyCarry(mymap, gw.wireB, num) || verifyInter(mymap, gw.wireB, num) && verifyCarry(mymap, gw.wireA, num)
}

func makeWire(c string, num int) string {
	wire := ""
	if num < 10 {
		wire = c + "0" + strconv.Itoa(num)
	} else {
		wire = c + strconv.Itoa(num)
	}
	return wire
}

func gateResult(i1, i2 int, gate string) int {
	switch gate {
	case "AND":
		return i1 & i2
	case "OR":
		return i1 | i2
	case "XOR":
		return i1 ^ i2
	default:
		panic(fmt.Sprintf("Gate %q not expected", gate))
	}

}

func computeEndMap(wireValueMap map[string]int, gatesAndWires []GateAndWires) {
	for i := 0; i < len(gatesAndWires); i++ {
		gW := gatesAndWires[i]

		a, ok := wireValueMap[gW.wireA]
		if !ok {
			continue
		}

		b, ok := wireValueMap[gW.wireB]
		if !ok {
			continue
		}

		wireValueMap[gW.wireC] = gateResult(a, b, gW.gate)
		gatesAndWires = append(gatesAndWires[:i], gatesAndWires[i+1:]...)

		i = -1
	}
}

func part1(wireValueMap map[string]int, gatesAndWires []GateAndWires) int {
	computeEndMap(wireValueMap, gatesAndWires)
	var keysWithZ []string
	for key := range wireValueMap {
		if key[0] == 'z' {
			keysWithZ = append(keysWithZ, key)
		}
	}
	slices.Sort(keysWithZ)

	ans := 0b0
	for i, key := range keysWithZ {
		value := wireValueMap[key]
		ans |= (value << i)
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

	i := 0

	wireValueMap := make(map[string]int)
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}

		splitLine := strings.Split(lines[i], ": ")

		value, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(fmt.Sprintf("Couldnt convert value from string to int for %q in line %q", splitLine[1], lines[i]))
		}

		wireValueMap[splitLine[0]] = value
	}

	gatesAndWires := make([]GateAndWires, len(lines)-i)
	for ; i < len(lines); i++ {
		splitLine := strings.Split(lines[i], " -> ")
		splitGate := strings.Split(splitLine[0], " ")
		gatesAndWires[len(gatesAndWires)-len(lines)+i] = GateAndWires{splitGate[0], splitGate[1], splitGate[2], splitLine[1]}
	}

	if len(os.Args) < 3 {
		panic("missing arguments")
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(gatesAndWires))
		return
	}
	fmt.Printf("PART 1: %v", part1(wireValueMap, gatesAndWires))
}
