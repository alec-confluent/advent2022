package main

import (
	"bufio"
	"fmt"
	"net/http"
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
	fattest := findFatty(elves)

	fmt.Println("The fattest elf has ", fattest.totalCalories, " total calories in tow.")
}

// finds the elf with the most calories in the tuck
func findFatty(elves []Elf) Elf {

	fattest := elves[0]
	fattest.getTotalCalories()

	for i, el := range elves {

		total := el.getTotalCalories()
		fmt.Println("total for elf", i, "is", total, "aka", el.totalCalories)

		if total > fattest.totalCalories {
			fmt.Println("elf", i, "is currently fattest")
			fattest = el
		}
	}

	return fattest
}

// parses the excercise data into structs
func readData(data string) []Elf {

	scanner := bufio.NewScanner(strings.NewReader(data))

	var elves = []Elf{}
	curr := Elf{0, []Item{}}

	for scanner.Scan() {

		fmt.Println(scanner.Text())

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
