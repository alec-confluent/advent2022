package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"utils/adventofcode"
)

var client http.Client

// struct for elfs
type Elf struct {
	totalCalories int64
	bag           []Item
}

// adds up calories for items in the bag
func (el *Elf) getTotalCalories() int64 {

	var total int64 = 0

	for _, item := range el.bag {
		total += item.calories
		// fmt.Println("total currently", total)
	}
	el.totalCalories = total

	return total
}

// definition of an item
type Item struct {
	calories int64
}

// the main dingy
func main() {

	// local utility package for getting data from advent of code website
	data := adventofcode.GetData("https://adventofcode.com/2022/day/1/input")

	// fmt.Println("response data:")
	// fmt.Println(data)

	elves := readData(data)
	top := findTop(elves, 3)

	fmt.Println("The top elves have ", top, " total calories in tow.")
}

// finds the top n elves and sums their calories
func findTop(elves []Elf, top int) int64 {

	elfTotals := []int64{}

	for _, el := range elves {

		total := el.getTotalCalories()
		elfTotals = append(elfTotals, total)
	}

	sort.Slice(elfTotals, func(i, j int) bool { return elfTotals[i] < elfTotals[j] })
	fmt.Println(elfTotals)

	var total int64 = 0
	for i, calories := range elfTotals[len(elfTotals)-top:] {

		fmt.Println(i, calories)
		total += calories
	}

	return total
}

// parses the excercise data into structs
func readData(data string) []Elf {

	scanner := bufio.NewScanner(strings.NewReader(data))

	var elves = []Elf{}
	curr := Elf{0, []Item{}}

	for scanner.Scan() {

		item := scanner.Text()

		if calories, err := strconv.ParseInt(item, 10, 64); err == nil {

			curr.bag = append(curr.bag, Item{calories})
		} else {

			elves = append(elves, curr)
			curr = Elf{0, []Item{}}
		}
	}

	return elves
}
