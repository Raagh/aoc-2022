package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadFile("./resources/day04.txt")
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
	count := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		firstRanges := strings.Split(pairs[0], "-")
		secondRanges := strings.Split(pairs[1], "-")
    leftStart, _ := strconv.Atoi(firstRanges[0])
    leftEnd, _ := strconv.Atoi(firstRanges[1])
    rightStart, _ := strconv.Atoi(secondRanges[0])
    rightEnd, _ := strconv.Atoi(secondRanges[1])

		if rightStart >= leftStart && rightEnd <= leftEnd {
			count++
		} else if leftStart >= rightStart && leftEnd <= rightEnd {
			count++
		}
	}

	return count
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	count := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		firstRanges := strings.Split(pairs[0], "-")
		secondRanges := strings.Split(pairs[1], "-")
    leftStart, _ := strconv.Atoi(firstRanges[0])
    leftEnd, _ := strconv.Atoi(firstRanges[1])
    rightStart, _ := strconv.Atoi(secondRanges[0])
    rightEnd, _ := strconv.Atoi(secondRanges[1])

		if rightStart >= leftStart && rightStart <= leftEnd {
			count++
		} else if leftStart >= rightStart && leftStart <= rightEnd {
			count++
		}
  }

	return count
}
