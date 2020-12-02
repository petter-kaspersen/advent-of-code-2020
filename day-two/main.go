package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (lines []string) {
	data, _ := ioutil.ReadFile("day-two/input.txt")

	return strings.Split(string(data), "\n")
}

func getParts(line string) (string, string, string) {
	_parts := strings.Split(line, " ")

	return _parts[0], _parts[1], _parts[2]
}

func getMinMax(minMax string) (int, int) {
	values := strings.Split(minMax, "-")

	val1, _ := strconv.Atoi(values[0])
	val2, _ := strconv.Atoi(values[1])

	return val1, val2
}

func getLetter(letterRaw string) string {
	values := strings.Split(letterRaw, ":")

	return values[0]
}

func partOne(lines []string) int {
	valid := 0

	for _, line := range lines {
		occ, lett, str := getParts(line)

		min, max := getMinMax(occ)
		letter := getLetter(lett)

		total := 0

		for _, char := range str {
			val := string(char)

			if val == letter {
				total++
			}
		}

		if total >= min && total <= max {
			valid++
		}
	}

	return valid
}

func partTwo(lines []string) int {
	valid := 0

	for _, line := range lines {
		occ, lett, str := getParts(line)

		x, y := getMinMax(occ)
		letter := getLetter(lett)

		if string(str[x-1]) == letter && string(str[y-1]) != letter {
			valid++
		} else if string(str[x-1]) != letter && string(str[y-1]) == letter {
			valid++
		}
	}

	return valid
}

func main() {
	lines := readFile()

	fmt.Printf("Part one answer: %v \n", partOne(lines))
	fmt.Printf("Part two answer: %v \n", partTwo(lines))

}
