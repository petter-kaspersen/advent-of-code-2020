package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readFile() (passports []string) {
	data, _ := ioutil.ReadFile("day-four/input.txt")

	lines := strings.Split(string(data), "\n\n")

	passports = make([]string, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		passports = append(passports, strings.ReplaceAll(l, "\n", " "))
	}

	return passports
}

func getFields() []string {
	return []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
}

func main() {
	input := readFile()

	fmt.Printf("Part one answer: %v \n", partOne(input))
	fmt.Printf("Part two answer: %v \n", partTwo(input))
}

func partOne(passports []string) int {
	fields := getFields()

	parsedPassports := parsePassports(passports)

	validPassports := 0
	for _, passport := range parsedPassports {
		valid := allPresent(passport, fields)

		if valid {
			validPassports++
		}
	}

	return validPassports
}

func partTwo(passports []string) int {
	fields := getFields()

	parsedPassports := parsePassports(passports)
	validPassports := make([][]string, 0)

	fullValid := 0

	for _, passport := range parsedPassports {
		valid := allPresent(passport, fields)

		if valid {
			validPassports = append(validPassports, passport)
		}
	}

	for _, passport := range validPassports {
		valid := isCompletelyValid(passport)

		if valid {
			fullValid++
		}
	}

	return fullValid
}

func parsePassports(passports []string) [][]string {
	newPassports := make([][]string, 0, len(passports))

	for _, passport := range passports {
		currFields := make([]string, 0)
		fields := strings.Split(passport, " ")

		for _, field := range fields {
			currFields = append(currFields, field)
		}

		newPassports = append(newPassports, currFields)
	}

	return newPassports
}

func allPresent(passport []string, fields []string) bool {
	for _, prop := range fields {

		_, match := Find(passport, prop)

		if !match {
			return false
		}
	}

	return true
}

func isCompletelyValid(passport []string) bool {

	for _, val := range passport {
		currVal := strings.Split(val, ":")

		switch currVal[0] {
		case "byr":
			if !isBetween(currVal[1], 1920, 2002) {
				return false
			}
		case "iyr":
			if !isBetween(currVal[1], 2010, 2020) {
				return false
			}

		case "eyr":
			if !isBetween(currVal[1], 2020, 2030) {
				return false
			}
		case "hgt":
			if !isHeightValid(currVal[1]) {
				return false
			}
		case "hcl":
			if !isHairColorValid(currVal[1]) {
				return false
			}
		case "ecl":
			if !isEyeColorValid(currVal[1]) {
				return false
			}
		case "pid":
			if !isPassportIdValid(currVal[1]) {
				return false
			}
		}
	}

	return true
}

func isPassportIdValid(val string) bool {
	isValid, _ := regexp.MatchString("^[0-9]{9}$", val)

	return isValid
}

func isEyeColorValid(val string) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, color := range colors {
		if color == val {
			return true
		}
	}

	return false
}

func isHairColorValid(val string) bool {
	isValid, _ := regexp.MatchString("^#[a-z0-9]{6}$", val)

	return isValid
}

func isBetween(val string, min int, max int) bool {
	intVal, err := strconv.Atoi(val)

	if err != nil {
		return false
	}

	return intVal >= min && intVal <= max
}

func isHeightValid(val string) bool {
	if strings.HasSuffix(val, "in") {
		intVal, _ := strconv.Atoi(strings.ReplaceAll(val, "in", ""))

		return intVal >= 59 && intVal <= 76
	} else if strings.HasSuffix(val, "cm") {
		intVal, _ := strconv.Atoi(strings.ReplaceAll(val, "cm", ""))

		return intVal >= 150 && intVal <= 193
	}
	return false
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if strings.Contains(item, val) {
			return i, true
		}
	}
	return -1, false
}
