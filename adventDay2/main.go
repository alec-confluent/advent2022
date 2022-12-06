package main

import (
	"bufio"
	"fmt"
	"strings"
	"utils/adventofcode"
)

// the main dingy
func main() {

	// local utility package for getting data from advent of code website
	data := adventofcode.GetData("https://adventofcode.com/2022/day/2/input")

	total := readData(data)
	fmt.Println("The player scores", total, "in total.")
}

// parses the excercise data
func readData(data string) int {

	scanner := bufio.NewScanner(strings.NewReader(data))

	total := 0
	for scanner.Scan() {

		elf, player := getChoiceScores(scanner.Text())
		roundScore := getMatchScore(choiceMap[elf], choiceMap[player])

		total += (player + roundScore)
	}

	return total
}

// maps the strategy guide to its numeric value
var scoreMap = map[string]int{
	"A": 1, // rock
	"B": 2, // paper
	"C": 3, // scissors
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, // scissors
}

// maps the numeric value to the selection
var choiceMap = map[int]string{
	1: "rock",     // rock
	2: "paper",    // paper
	3: "scissors", // scissors
}

// determines if the player won, lost or tied the round
// and returns numeric score for the round
func getMatchScore(elf string, player string) int {

	if elf == player {
		return 3
	} else if elf == "rock" && player == "paper" {
		return 6
	} else if elf == "paper" && player == "scissors" {
		return 6
	} else if elf == "scissors" && player == "rock" {
		return 6
	}

	return 0
}

// assigns numeric value to each side of the strategy guide line
func getChoiceScores(strategyGuide string) (int, int) {

	// returns each "word",
	// aka the elfs choice and the recommended player choice
	sides := strings.Fields(strategyGuide)

	// returns each side numeric value
	return scoreMap[sides[0]], scoreMap[sides[1]]
}
