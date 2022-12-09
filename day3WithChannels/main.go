package main

import (
	"bufio"
	"fmt"
	"strings"
	"utils/adventofcode"
)

func main() {

	// local utility package for getting data from advent of code website
	data := adventofcode.GetData("https://adventofcode.com/2022/day/3/input")

	// make some channel for data processing
	a := make(chan string)
	b := make(chan string)
	quit := make(chan string)

	// sends data to one of the two processing channels
	go readData(data, a, b, quit)

	// receives the results of processing
	total := receive(a, b, quit)
	fmt.Println("The elves have", total, "in total priority points.")
}

// parses the excercise data into either channel A or B
func readData(data string, a chan string, b chan string, q chan string) {

	scanner := bufio.NewScanner(strings.NewReader(data))

	lever := true
	for scanner.Scan() {

		rucksack := scanner.Text()

		if lever {
			a <- rucksack
		} else {
			b <- rucksack
		}

		lever = !lever
	}
	q <- "done"
}

// pulls data from processing channels
func receive(a, b, q <-chan string) int {

	letterValues := getMap(alphabet)
	total := 0

	for {
		select {

		case rucksack := <-a:
			total += searchRucksack(rucksack, letterValues)

		case rucksack := <-b:
			total += searchRucksack(rucksack, letterValues)

		case done := <-q:
			fmt.Println("aaaaand we're", done)
			return total
		}
	}
}

// checks if any rune from a is contained in b
func searchRucksack(rucksack string, letterValues map[string]int) int {

	compartmentA := rucksack[:len(rucksack)/2]
	compartmentB := rucksack[len(rucksack)/2:]

	duplicate := findAinB(compartmentA, compartmentB)
	return letterValues[duplicate]
}

// checks if any rune from a is contained in b
func findAinB(a string, b string) string {
	var duplicate string

	for _, letter := range a {
		if strings.Contains(b, string(letter)) {
			duplicate = string(letter)
			break
		}
	}

	return duplicate
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
