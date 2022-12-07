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

	// local utility package for getting data from advent of code website
	data := adventofcode.GetData("https://adventofcode.com/2022/day/4/input")

	total := readData(data)
	fmt.Println("The elves have", total, "overlaps.")
}

// parses the excercise data
func readData(data string) int {

	scanner := bufio.NewScanner(strings.NewReader(data))

	total := 0
	for scanner.Scan() {

		line := scanner.Text()

		// fmt.Println("line", line)
		pairs := strings.Split(line, ",")

		left := getInts(strings.Split(pairs[0], "-"))
		right := getInts(strings.Split(pairs[1], "-"))

		if left[0] >= right[0] && left[1] <= right[1] {
			total += 1
		} else if right[0] >= left[0] && right[1] <= left[1] {
			total += 1
		}

		fmt.Println("")
	}

	return total
}

func getInts(sides []string) []int {

	sideInts := []int{}

	for _, side := range sides {
		sideInt, _ := strconv.Atoi(side)
		sideInts = append(sideInts, sideInt)
	}

	return sideInts
}
