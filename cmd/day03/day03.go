package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadFile("./resources/day03.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

func part1(input string) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	totalPriorities := 0
	for _, line := range lines {
		firstHalf := string(line[:len(line)/2])
		secondHalf := string(line[len(line)/2:])
		for _, l := range firstHalf {
			if strings.ContainsRune(secondHalf, l) {
				totalPriorities += charToInt(l)
				break
			}
		}
	}

	return totalPriorities
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	totalPriorities := 0
	previousLines := make([]string, 0)
	for _, line := range lines {
		if len(previousLines) == 2 {
			line1 := previousLines[0]
			line2 := previousLines[1]

			for _, l := range line {
				if strings.ContainsRune(line1, l) && strings.ContainsRune(line2, l) {
					totalPriorities += charToInt(l)
					previousLines = make([]string, 0)
					break
				}
			}
		} else {
			previousLines = append(previousLines, line)
		}
	}

	return totalPriorities
}

func IsIntDivisibleBy3(n int) bool {
	digits := strconv.Itoa(n)
	sumOfDigits := 0
	for _, digit := range digits {
		d, _ := strconv.Atoi(string(digit))
		sumOfDigits += d
	}

	return (sumOfDigits % 3) == 0
}

func charToInt(c rune) int {
	if int(c) > 90 {
		// lower case
		return int(c) - 96
	} else {
		// upper case
		return int(c) - 38
	}
}
