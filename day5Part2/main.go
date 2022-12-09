package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"utils/adventofcode"
)

// the main dingy
func main() {

	// local utility package for getting data from advent of code website
	data := adventofcode.GetData("https://adventofcode.com/2022/day/5/input")

	answer := readData(data)
	fmt.Println("The answer is:", answer)
}

// parses the excercise data
func readData(data string) string {

	scanner := bufio.NewScanner(strings.NewReader(data))

	answer := ""
	crateLines := []string{}
	stacks := [][]string{}
	for scanner.Scan() {

		line := scanner.Text()
		fmt.Println(line)

		if strings.HasPrefix(line, "move") {

			doMove(line, stacks)
		} else if len(line) != 0 {

			crateLines = append(crateLines, line)
		} else if len(line) == 0 {

			stacks = createStacks(crateLines)
		}
	}

	for _, stack := range stacks {
		answer += stack[len(stack)-1]
	}

	return answer
}

// reads the instructions and moves (n) crates from
// one stack to another stack
func doMove(line string, stacks [][]string) {

	move, from, to := getInstructions(line)

	size := len(stacks[from])

	movingCrates := stacks[from][size-move:]

	stacks[from] = stacks[from][:size-move]
	stacks[to] = append(stacks[to], movingCrates...)

	instructions := fmt.Sprintln("moved", move, "crate(s) from stack", from+1, "to stack", to+1)
	print(stacks, instructions)
}

// snags all the numbers out of the instruction string
// and for any stack numbers converts them to 0 based index locations
func getInstructions(line string) (int, int, int) {

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	instructions := []string{}

	matches := re.FindAllString(line, -1)
	for _, match := range matches {

		instructions = append(instructions, match)
	}

	// -1 to convert the instructions to the index locations
	return getInt(instructions[0]), getInt(instructions[1]) - 1, getInt(instructions[2]) - 1
}

func createStacks(lines []string) [][]string {

	// the bottom line is the stack number we
	// we use this line to get an index location
	// which will be later used to create the stacks with crates
	crateLocations := getLocations(lines[len(lines)-1])

	// add an empty slice for each stack, to be filled later
	// we could do this on the fly but it will make the next
	// loop a bit verbose...
	stacks := [][]string{}
	for i := 0; i < len(crateLocations); i++ {
		stacks = append(stacks, []string{})
	}

	// loop the slice in reverse order
	// we dont need the last line because
	// its just the stack number...so -2
	for i := len(lines) - 2; i >= 0; i-- {

		line := lines[i]

		for j, location := range crateLocations {

			content := line[location]

			// if the content is not empty
			// then add the crate to the stack
			if content != 32 {

				stacks[j] = append(stacks[j], string(content))
			}
		}
	}

	return stacks
}

// returns an array of index locations for crates in each line
// note that a crate may not exist on a line
// example
// [D]         <<< this is a line
// [N] [C]
// [Z] [M] [P]
//  1   2   3
func getLocations(line string) []int {

	locations := []int{}

	for i, r := range line {
		// if its not a space..
		if r != 32 {

			locations = append(locations, i)
		}
	}

	return locations
}

// just for fun we pretty print the stacks to the console
func print(stacks [][]string, instructions string) {
	fmt.Print("\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n")
	max := 0
	for _, v := range stacks {
		if len(v) > max {
			max = len(v)
		}
	}
	for i := max; i >= 0; i-- {
		for _, v := range stacks {
			if len(v) > i {
				fmt.Print("[", v[i], "] ")
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Printf("\n")
	}
	for i, _ := range stacks {
		fmt.Print(" ", i+1, "  ")
	}
	fmt.Printf("\n")
	fmt.Println(instructions)
	time.Sleep(time.Second / 2)
}

// converts a string to an int value
func getInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
