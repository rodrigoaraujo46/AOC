package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Secret int

func (s Secret) getNextSecret() Secret {
	s = (s ^ (s << 6)) % 16_777_216
	s = (s ^ (s >> 5)) % 16_777_216
	s = (s ^ (s << 11)) % 16_777_216

	return s
}

func (s Secret) getSecretN(n int) Secret {
	for range n {
		s = s.getNextSecret()
	}

	return s
}

func (s Secret) getFirstNPrices(n int) Prices {
	prices := make(Prices, n)
	priceVal := int(s % 10)

	price := Price{val: priceVal}
	prices[0] = price

	for i := range n - 1 {
		s = s.getNextSecret()
		priceVal = int(s % 10)
		price = Price{val: priceVal, dif: priceVal - prices[i].val}
		prices[i+1] = price
	}

	return prices
}

type Secrets []Secret

type Price struct {
	val int
	dif int
}

type Prices []Price

func printAll(allMonkeyPrices []Prices) {
	for _, prices := range allMonkeyPrices {
		fmt.Println(prices)
	}
}

func part2(secrets Secrets) int {
	allMonkeyPrices := make([]Prices, len(secrets))

	for i, secret := range secrets {
		firstNPrices := secret.getFirstNPrices(2000)
		allMonkeyPrices[i] = firstNPrices
	}

	sequencePrice := make(map[[4]int]int, 0)
	for j := 0; j < len(allMonkeyPrices); j++ {
		seen := make(map[[4]int]bool, 0)
		for k := 4; k < len(allMonkeyPrices[0]); k++ {
			sequence := [4]int{allMonkeyPrices[j][k-3].dif,
				allMonkeyPrices[j][k-2].dif,
				allMonkeyPrices[j][k-1].dif,
				allMonkeyPrices[j][k].dif}
			if seen[sequence] {
				continue
			}
			seen[sequence] = true
			sequencePrice[sequence] += allMonkeyPrices[j][k].val
		}
	}

	maxPrice := 0
	for sequences := range sequencePrice {
		if sequencePrice[sequences] > maxPrice {
			maxPrice = sequencePrice[sequences]
		}
	}

	return maxPrice
}

func part1(secrets Secrets) int {
	ans := 0
	for _, secret := range secrets {
		ans += int(secret.getSecretN(2000))
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

	secrets := make(Secrets, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		secrets[i] = Secret(num)
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(secrets))
		fmt.Printf("PART 2: %v\n", part2(secrets))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(secrets))
		return
	}
	fmt.Printf("PART 1: %v", part1(secrets))
}
