package main

/*
--- Day 1: Calorie Counting ---
The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies.
One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's
inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000
(carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
*/

/*
--- Part Two ---
By the time you calculate the answer to the Elves' question, they've already realized that the Elf carrying the most Calories of food might eventually run out of snacks.

To avoid this unacceptable situation, the Elves would instead like to know the total Calories carried by the top three Elves carrying the most Calories. That way, even if one
of those Elves runs out of snacks, they still have two backups.

In the example above, the top three Elves are the fourth Elf (with 24000 Calories), then the third Elf (with 11000 Calories), then the fifth Elf (with 10000 Calories). The sum
of the Calories carried by these three elves is 45000.

Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// this reads the file into an array of stings
// a single space would mean a new line, a double space would mean a new elf
func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	check(err)

	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func elfSums(elfData []string) []int {
	var currVal int = 0
	var sums []int

	for _, element := range elfData {
		if strings.TrimSpace(element) == "" {
			sums = append(sums, currVal)
			currVal = 0
		} else {
			val, err := strconv.Atoi(element)
			check(err)
			currVal += val
		}
	}

	return sums
}

func topThree(sums []int) int {
	sort.Ints(sums)
	var len = len(sums)
	var topThree = sums[len-3 : len]

	var sum int = 0
	for _, element := range topThree {
		sum += element
	}

	return sum
}

func main() {
	filename := "./data/day1"
	data, err := readLines(filename)
	check(err)

	var elfSums []int = elfSums(data)
	var sum = topThree(elfSums)
	fmt.Println("max elf is:", sum)
}
