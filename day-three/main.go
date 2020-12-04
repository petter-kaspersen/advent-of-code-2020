package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile() []string {
	data, _ := ioutil.ReadFile("day-three/input.txt")

	return strings.Split(string(data), "\n")
}

func main() {
	rows := readFile()

	fmt.Printf("Part one answer: %v \n", doGetTrees(rows, 3, 1))
	fmt.Printf("Part two answer: %v \n", partTwo(rows))
}

func partTwo(rows []string) int {
	trees1 := doGetTrees(rows, 1, 1)
	trees2 := doGetTrees(rows, 3, 1)
	trees3 := doGetTrees(rows, 5, 1)
	trees4 := doGetTrees(rows, 7, 1)
	trees5 := doGetTrees(rows, 1, 2)

	return trees1 * trees2 * trees3 * trees4 * trees5
}

func doGetTrees(rows []string, right int, down int) int {
	posX := 0
	posY := 0
	trees := 0
	gridWidth := len(rows[0])
	gridHeight := len(rows)

	for posY < gridHeight {
		if string(rows[posY][posX]) == "#" {
			trees++
		}
		posX = (posX + right) % gridWidth
		posY += down
	}

	return trees
}
