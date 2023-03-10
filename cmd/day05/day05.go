package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
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
  return 0
}

func part2(input string) int {
  return 0
}
