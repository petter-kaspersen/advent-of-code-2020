package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (nums []int) {
	data, _ := ioutil.ReadFile("day-one/input.txt")

	lines := strings.Split(string(data), "\n")

	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, _ := strconv.Atoi(l)

		nums = append(nums, n)
	}

	return nums
}

func main() {
	input := readFile()

	partOneAnswer := partOne(input)
	partTwoAnswer := partTwo(input)

	fmt.Printf("Part one answer: %v \n", partOneAnswer)
	fmt.Printf("Part two answer: %v \n", partTwoAnswer)

}

func partOne(inp []int) int {
	nums := make(map[int]int)

	for _, num := range inp {
		if nums[2020-num] > 0 {
			return num * (2020 - num)
		}

		nums[num] = 1
	}
	return 0
}

// Don't open, dead inside
func partTwo(inp []int) int {
	for _, num1 := range inp {
		for _, num2 := range inp {
			for _, num3 := range inp {
				sum := num1 + num2 + num3

				if sum == 2020 {
					return num1 * num2 * num3
				}
			}
		}
	}
	return 0
}
