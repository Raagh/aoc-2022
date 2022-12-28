package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := utils.ReadFile("./resources/day02.txt")
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
	lines := strings.Split(input, "\n")
	totalScore := 0

	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		moves := strings.Split(line, " ")
		opponentMove := string(moves[0])
		playerMove := string(moves[1])
		roundScore := 0

		if playerMove == "X" {
			roundScore += 1
		} else if playerMove == "Y" {
			roundScore += 2
		} else if playerMove == "Z" {
			roundScore += 3
		}

		if opponentMove == "A" {
			if playerMove == "Z" {
				roundScore += 0
			} else if playerMove == "X" {
				roundScore += 3
			} else if playerMove == "Y" {
				roundScore += 6
			}
		}

		if opponentMove == "B" {
			if playerMove == "X" {
				roundScore += 0
			} else if playerMove == "Y" {
				roundScore += 3
			} else if playerMove == "Z" {
				roundScore += 6
			}
		}

		if opponentMove == "C" {
			if playerMove == "Y" {
				roundScore += 0
			} else if playerMove == "Z" {
				roundScore += 3
			} else if playerMove == "X" {
				roundScore += 6
			}
		}

		totalScore += roundScore
	}

	return totalScore
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	totalScore := 0

	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		moves := strings.Split(line, " ")
		opponentMove := string(moves[0])
		playerMove := string(moves[1])
		roundScore := 0

		if playerMove == "X" {
			roundScore += 0
		} else if playerMove == "Y" {
			roundScore += 3
		} else if playerMove == "Z" {
			roundScore += 6
		}

		if opponentMove == "A" {
			if playerMove == "X" {
				roundScore += 3
			} else if playerMove == "Y" {
				roundScore += 1
			} else if playerMove == "Z" {
				roundScore += 2
			}
		}

		if opponentMove == "B" {
			if playerMove == "X" {
				roundScore += 1
			} else if playerMove == "Y" {
				roundScore += 2
			} else if playerMove == "Z" {
				roundScore += 3
			}
		}

		if opponentMove == "C" {
			if playerMove == "X" {
				roundScore += 2
			} else if playerMove == "Y" {
				roundScore += 3
			} else if playerMove == "Z" {
				roundScore += 1
			}
		}

		totalScore += roundScore
	}

	return totalScore
}
