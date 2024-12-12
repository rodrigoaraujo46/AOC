package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		print(err.Error())
		return
	}

	var nums []int
	for _, char := range content {
		num, _ := strconv.Atoi(string(char))
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



func getDisk(nums []int) []string {
	var disk []string
	id := 0
	for i, num := range nums {
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				disk = append(disk, strconv.Itoa(i/2))
			}
			continue
		}
		for j := 0; j < num; j++ {
			disk = append(disk, ".")
		}
		id++
	}
	return disk
}

func continousBlocks(disk []string) bool {
	ans := 0
	previous := "1"
	for i := 0; i < len(disk); i++ {
		if disk[i] == "." && previous != "." {
			ans++
			if ans >= 2 {
				return false
			}
		}
		previous = disk[i]
	}
	if disk[len(disk)-1] != "." && ans != 0 {
		return false
	}
	return true
}

func transformDisk(nums []int) []string {
	disk := getDisk(nums)
	for i := len(disk) - 1; i >= 0 && !continousBlocks(disk); i-- {
		if disk[i] != "." {
			for j, char := range disk {
				if char == "." {
					disk[j] = disk[i]
					disk[i] = "."
					break
				}
			}
		}
	}
	return disk
}

func part1(nums []int) (ans int) {
	disk := transformDisk(nums)
	for i, num := range disk {
		if num != "." {
			numI, _ := strconv.Atoi(num)
			ans += i * numI
		}
	}
	return ans
}

func fileSize(disk []string, id int) int {
	ans := 0
	str := strconv.Itoa(id)
	for _, num := range disk {
		if num == str {
			ans++
		}
	}
	return ans
}

func firstEmpty(disk []string, size int) int {
	for i := 0; i <= len(disk)-size; i++ {
		if disk[i] == "." {
			count := 1
			for j := i + 1; j < i+size && j < len(disk); j++ {
				if disk[j] == "." {
					count++
				} else {
					break
				}
			}
			if count == size {
				return i
			}
		}
	}
	return -1
}

func findFirst(disk []string, r string) int {
	return slices.Index(disk, r)
}

func transformDiskByBlocks(nums []int) []string {
	disk := getDisk(nums)
	for i := len(disk); i >= 0; i-- {
		if i == '.' {
			continue
		}

		numid := findFirst(disk, strconv.Itoa(i))
		if numid == -1 {
			continue
		}

		fs := fileSize(disk, i)
		j := firstEmpty(disk, fs)
		if j == -1 || j > numid {
			continue
		}

		for j := range disk {
			if disk[j] == strconv.Itoa(i) {
				disk[j] = "."
			}
		}

		for x := j; x < j+fs; x++ {
			disk[x] = strconv.Itoa(i)
		}

	}
	return disk
}

func part2(nums []int) (ans int) {
	disk := transformDiskByBlocks(nums)
	for i, num := range disk {
		if num != "." {
			numI, _ := strconv.Atoi(num)
			ans += i * numI
		}
	}
	return
}
