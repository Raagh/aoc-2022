package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadFile("./resources/day01.txt")
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
	caloriesPerElf := getCaloriesPerElf(input)
	return max(caloriesPerElf) 
}

func part2(input string) int {
	caloriesPerElf := getCaloriesPerElf(input)
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))
  return sum(caloriesPerElf[0:3])
}

func getCaloriesPerElf(input string) []int {
	elfCalories := strings.Split(input, "\n\n")
	caloriesPerElf := make([]int, 0)
	for _, line := range elfCalories {
		caloriesArray := strings.Split(line, "\n")
		calories := make([]int, 0)
		for _, v := range caloriesArray {
			if len(v) > 0 {
				j, err := strconv.Atoi(v)
				if err != nil {
					fmt.Println(v)
					panic(v)
				}
				calories = append(calories, j)
			}
		}

		result := sum(calories)
		caloriesPerElf = append(caloriesPerElf, result)
	}

	return caloriesPerElf
}

func max(array []int) int {
	max := 0
	for _, v := range array {
		if v > max {
			max = v
		}
	}

	return max
}

func sum(array []int) int {
	result := 0
	for _, j := range array {
		result += j
	}

	return result
}
