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
	for scanner.Scan() {

		rucksack := scanner.Text()

		// split the sack in half
		compartmentA := rucksack[:len(rucksack)/2]
		compartmentB := rucksack[len(rucksack)/2:]

		duplicate := findAinB(compartmentA, compartmentB)

		// fmt.Println("Rucksack has", rucksack)
		// fmt.Println("Compartment A has", compartmentA)
		// fmt.Println("Compartment B has", compartmentB)
		// fmt.Println("Letter", duplicate, "is in both compartments")
		// fmt.Println("Its priority is", letterValues[duplicate])

		// fmt.Println("")

		total += letterValues[duplicate]
	}

	return total
}

// checks if any rune from a is contained in b
func findAinB(a string, b string) string {
	for _, letter := range a {
		if strings.Contains(b, string(letter)) {
			return string(letter)
		}
	}
	return ""
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
