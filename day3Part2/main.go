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
	data := adventofcode.GetData("https://adventofcode.com/2022/day/3/input")

	total := readData(data)
	fmt.Println("The elves have", total, "in total priority points.")
}

// parses the excercise data
func readData(data string) int {

	scanner := bufio.NewScanner(strings.NewReader(data))

	letterValues := getMap(alphabet)

	total := 0

	curr := []string{}
	for scanner.Scan() {

		rucksack := scanner.Text()
		curr = append(curr, rucksack)

		if len(curr) == 3 {

			duplicate := findAinBandC(curr[0], curr[1], curr[2])

			total += letterValues[duplicate]

			// fmt.Println("Rucksack 1 has", curr[0])
			// fmt.Println("Rucksack 2 has", curr[1])
			// fmt.Println("Rucksack 3 has", curr[2])
			// fmt.Println("Letter", duplicate, "is in all 3 rucksacks")
			// fmt.Println("Its priority is", letterValues[duplicate])

			// fmt.Println("")

			curr = []string{}
		}
	}

	return total
}

// checks if any rune from a is contained in b
func findAinBandC(a string, b string, c string) string {
	for _, letter := range a {
		if contains(b, letter) {
			if contains(c, letter) {
				return string(letter)
			}
		}
	}
	return ""
}

// checks if the rune is in the string
func contains(s string, r rune) bool {
	for _, l := range s {
		if r == l {
			return true
		}
	}
	return false
}

// im lazy so im not gonna type the alphabet => number in a map literal
func getMap(alphabet []string) map[string]int {

	m := map[string]int{}

	for i, letter := range alphabet {
		m[string(letter)] = i + 1
		m[strings.ToUpper(string(letter))] = i + 1 + 26
	}

	return m
}

// a to z
var alphabet = []string{
	"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l",
	"m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x",
	"y", "z",
}
