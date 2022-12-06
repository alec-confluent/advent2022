package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"utils/adventofcode"
)

// the main dingy
func main() {

	// check(err) // does some error handling
	data := adventofcode.GetData("https://adventofcode.com/2022/day/1/input")

	highest := readData(data)
	fmt.Println("The fattest elf has ", highest, " total calories in tow.")
}

// parses the excercise data into structs
func readData(data string) int64 {

	scanner := bufio.NewScanner(strings.NewReader(data))

	var highest int64 = 0
	var curr int64 = 0

	for scanner.Scan() {

		item := scanner.Text()

		if calories, err := strconv.ParseInt(item, 10, 64); err == nil {

			curr += calories
		} else {

			curr = 0
		}

		if curr > highest {
			highest = curr
		}
	}

	return highest
}
