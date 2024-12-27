package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
	"strings"
)

func allowedInSet(connectionsMap map[string][]string, neighbour string, req map[string]bool) bool {
	for query := range req {
		if !slices.Contains(connectionsMap[neighbour], query) {
			return false
		}
	}
	return true
}

func search(connectionsMap map[string][]string, pc string, req map[string]bool, sets map[string]bool) {
	keys := make([]string, 0, len(req))
	for key := range req {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	key := strings.Join(keys, ",")

	if sets[key] == true {
		return
	}
	sets[key] = true

	for _, neighbour := range connectionsMap[pc] {
		if req[neighbour] == true {
			continue
		}

		if !allowedInSet(connectionsMap, neighbour, req) {
			continue
		}

		clone := maps.Clone(req)
		clone[neighbour] = true
		search(connectionsMap, neighbour, clone, sets)
	}
}

func part2(connectionsmap map[string][]string) string {
	sets := make(map[string]bool)

	for key := range connectionsmap {
		networkset := make(map[string]bool)
		networkset[key] = true
		search(connectionsmap, key, networkset, sets)
	}

	maxlen := 0
	maxstr := ""
	for key := range sets {
		if len(key) > maxlen {
			maxlen = len(key)
			maxstr = key
		}
	}
	return maxstr
}

func sliceToKey(slice []string) string {
	sort.Strings(slice)
	return strings.Join(slice, ",")
}

func part1(connectionsMap map[string][]string) int {
	tripleCon := make(map[string]bool)
	for pc, connections := range connectionsMap {
		for i := 0; i < len(connections); i++ {
			pcI := connections[i]
			pcICons := connectionsMap[pcI]
			for j := i + 1; j < len(connections); j++ {
				pcJ := connections[j]
				if slices.Contains(pcICons, pcJ) {
					key := sliceToKey([]string{pc, pcI, pcJ})
					tripleCon[key] = true
				}
			}
		}
	}

	ans := 0
	for triples := range tripleCon {
		pcs := strings.Split(triples, ",")
		for _, pc := range pcs {
			if pc[0] == 't' {
				ans++
				break
			}
		}
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

	connections := make(map[string][]string)
	for _, line := range lines {
		pcs := strings.Split(line, "-")
		pc1 := pcs[0]
		pc2 := pcs[1]
		connections[pc1] = append(connections[pc1], pc2)
		connections[pc2] = append(connections[pc2], pc1)
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(connections))
		fmt.Printf("PART 2: %v\n", part2(connections))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(connections))
		return
	}
	fmt.Printf("PART 1: %v", part1(connections))
}
