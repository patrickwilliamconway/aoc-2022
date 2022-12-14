package main

/*
--- Day 3: Rucksack Reorganization ---
One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so
 a few items now need to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this
rule for exactly one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single
 lowercase or uppercase letter (that is, a and A refer to different types of items).

The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first
 half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp.
The only item type that appears in both compartments is lowercase p.
The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
The fourth rucksack's compartments only share item type v.
The fifth rucksack's compartments only share item type t.
The sixth rucksack's compartments only share item type s.
To help prioritize item rearrangement, every item type can be converted to a priority:

Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.
In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?

total: 8243
--- Part Two ---
As you finish identifying the misplaced items, the Elves come to you with another issue.

For safety, the Elves are divided into groups of three. Every Elf carries a badge that identifies their group. For efficiency, within each group of three Elves, the badge is the only item type carried by all
 three Elves. That is, if a group's badge is item type B, then all three Elves will have item type B somewhere in their rucksack, and at most two of the Elves will be carrying any other item type.

The problem is that someone forgot to put this year's updated authenticity sticker on the badges. All of the badges need to be pulled out of the rucksacks so the new authenticity stickers can be attached.

Additionally, nobody wrote down which item type corresponds to each group's badges. The only way to tell which item type is the right one is by finding the one item type that is common between all three Elves
 in each group.

Every set of three lines in your list corresponds to a single group, but each group can have a different badge item type. So, in the above example, the first group's rucksacks are the first three lines:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
And the second group's rucksacks are the next three lines:

wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
In the first group, the only item type that appears in all three rucksacks is lowercase r; this must be their badges. In the second group, their badge item type must be Z.

Priorities for these items must still be found to organize the sticker attachment efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities of those item types?

total: 2631
*/

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

type Pair struct {
	sackContents string
	priority     int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func log(s string) {
	fmt.Println(s)
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

func processData(sacks []string) int {
	total := 0
	for _, sack := range sacks {
		total += processSack(sack)
	}
	return total
}

func processGroups(sacks []string) int {
	total := 0

	for i := 0; i < len(sacks); i += 3 {
		group := sacks[i : i+3]
		badge, err := intersectGroup(group[0], group[1], group[2])
		check(err)

		// log(fmt.Sprintf("found %q as intersection of: \n%s\n%s\n%s", badge, group[0], group[1], group[2]))
		total += priority(badge)
	}

	return total
}

type key struct {
	first  bool
	second bool
	third  bool
}

func mts(m map[rune]key) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%q: 1=%v, 2=%v, 3=%v", key, value.first, value.second, value.third)
	}
	return b.String()
}

func intersectGroup(one string, two string, three string) (rune, error) {
	m := make(map[rune]key)
	iterateString(m, one, 1)
	iterateString(m, two, 2)
	iterateString(m, three, 3)

	for k, v := range m {
		if v.first && v.second && v.third {
			return k, nil
		}
	}

	return ' ', errors.New("no intersection")
}

func iterateString(m map[rune]key, s string, pack int) map[rune]key {
	for _, ch := range s {
		curr := m[ch]
		if pack == 1 {
			curr.first = true
		} else if pack == 2 {
			curr.second = true
		} else {
			curr.third = true
		}
		m[ch] = curr
	}
	return m
}

func processSack(sack string) int {
	length := len(sack)
	left := sack[:length/2]
	right := sack[length/2:]

	r, err := intersection(left, right)
	check(err)

	return priority(r)
}

func intersection(left string, right string) (rune, error) {
	m := make(map[rune]bool)

	// track characters in left
	for _, ch := range left {
		m[ch] = true
	}

	// iterate right to find intersection
	for _, ch := range right {
		if m[ch] {
			return ch, nil
		}
	}

	return ' ', errors.New("no intersection")
}

func priority(s rune) int {
	// "A" is 65 in decimal (hex 46), "Z" is 90 (5a hex)
	// "a" is 97 in decimal (hex 61), "z" is 122 (7a hex)
	v := int(s) - 97 + 1

	// if v is negative, that means we're dealing with uppercase
	if v < 0 {
		// 32 offset gets us back to 1 for "a"
		// 26 gets offsets for letters
		v += 32 + 26
	}

	return v
}

func test(input string, expectedValue int) int {
	val := processSack(input)

	if val != expectedValue {
		panic(fmt.Sprintf("%s :: expected: %d, got %d", input, expectedValue, val))
	}

	return val
}

func runTests(pairs []Pair, expectedTotal int) {
	testTotal := 0
	for _, p := range pairs {
		testTotal += test(p.sackContents, p.priority)
	}

	if testTotal != expectedTotal {
		panic(fmt.Sprintf("expectedTotal: %d, testTotal %d", expectedTotal, testTotal))
	}
}

func runGroupTests(data []string, expectedTotal int) {
	testTotal := processGroups(data)
	if testTotal != expectedTotal {
		panic(fmt.Sprintf("expectedTotal: %d, testTotal %d", expectedTotal, testTotal))
	}
}

func main() {

	tests := []Pair{
		Pair{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		Pair{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		Pair{"PmmdzqPrVvPwwTWBwg", 42},
		Pair{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		Pair{"ttgJtRGJQctTZtZT", 20},
		Pair{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
	}
	runTests(tests, 157)

	groupTests := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	runGroupTests(groupTests, 70)

	filedata, err := readLines("./data/day3")
	check(err)

	// total := processData((filedata))
	total := processGroups(filedata)

	fmt.Println("total:", total)
}
