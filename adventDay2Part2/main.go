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

		elf, roundScore := getChoiceScores(scanner.Text())
		player := getMatchScore(elf, roundScore)

		total += (player + roundScore)
	}

	return total
}

// maps the strategy guide to its numeric value
var scoreMap = map[string]int{
	"X":        0, // lose
	"Y":        3, // draw
	"Z":        6, // win
	"rock":     1, // rock
	"paper":    2, // paper
	"scissors": 3, // scissors
}

// maps the choices from the guide
var choiceMap = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

// maps the numeric value to the selection
var winningChoices = map[string]string{
	"scissors": "rock",
	"paper":    "scissors",
	"rock":     "paper",
}

// maps the numeric value to the selection
var losingChoices = map[string]string{
	"rock":     "scissors",
	"scissors": "paper",
	"paper":    "rock",
}

// determines if the player won, lost or tied the round
// and returns numeric score for the round
func getMatchScore(elf string, roundScore int) int {

	if roundScore == 3 {
		return scoreMap[elf]
	} else if roundScore == 6 {
		return scoreMap[winningChoices[elf]]
	}

	return scoreMap[losingChoices[elf]]
}

func getChoiceScores(strategyGuide string) (string, int) {

	// returns each "word",
	// aka the elfs choice and the recommended player choice
	sides := strings.Fields(strategyGuide)

	// returns each side numeric value
	return choiceMap[sides[0]], scoreMap[sides[1]]
}
