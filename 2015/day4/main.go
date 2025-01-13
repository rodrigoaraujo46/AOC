package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	secretKey := string(content[:len(content)-1])
	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(secretKey))
		fmt.Printf("PART 2: %v\n", part2(secretKey))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(secretKey))
		return
	}
	fmt.Printf("PART 1: %v", part1(secretKey))
}

func part1(secret string) int {
	check := "00000"
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(secret + strconv.Itoa(i)))
		hash16 := hex.EncodeToString(hash[:])
		if hash16[:5] == check {
			return i
		}
	}
}

func part2(secret string) int {
	check := "000000"
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(secret + strconv.Itoa(i)))
		hash16 := hex.EncodeToString(hash[:])
		if hash16[:6] == check {
			return i
		}
	}
}
