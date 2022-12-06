package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	// check(err) // does some error handling
	data := getData("https://adventofcode.com/2022/day/1/input")

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

// gets the excercise data from the code advent site
func getData(address string) string {

	// open up the json file to get our session key
	jsonFile, err := os.Open("../info.json")

	if err != nil {
		log.Fatalf("Error... %s", err.Error())
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(jsonBytes), &result)

	sessionId := fmt.Sprintf("%v", result["key"])

	// create a cookie with the session id for the request
	sessionCookie := &http.Cookie{
		Name:   "session",
		Value:  sessionId,
		MaxAge: 300,
	}

	// create an http request for the excercise data
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		log.Fatalf("Error... %s", err.Error())
	}

	req.AddCookie(sessionCookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error... %s", err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("response status: %d\n", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error... %s", err.Error())
	}

	stringBody := string(body[:])

	return stringBody
}
