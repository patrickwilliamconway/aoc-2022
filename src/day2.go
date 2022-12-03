/*
--- Day 2: Rock Paper Scissors ---
The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.

Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape.

	Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent

	is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have

	been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape

	you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.

For example, suppose you were given the following strategy guide:

A Y
B X
C Z
This strategy guide predicts and recommends the following:

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
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

func toRPS(played string) string {
	switch played {
	case "A", "X":
		return "r"
	case "B", "Y":
		return "p"
	default:
		return "s"
	}
}

func score(us string, them string) int {
	if us == "r" && them == "s" {
		return 1 + 6
	} else if us == "r" && them == "p" {
		return 1 + 0
	} else if us == "r" && them == "r" {
		return 1 + 3
	} else if us == "p" && them == "p" {
		return 2 + 3
	} else if us == "p" && them == "s" {
		return 2
	} else if us == "p" && them == "r" {
		return 2 + 6
	} else if us == "s" && them == "r" {
		return 3
	} else if us == "s" && them == "s" {
		return 3 + 3
	} else {
		return 3 + 6
	}
}

func determineWinner(p1 string, p2 string) int {
	p1Move := toRPS(p1)
	p2Move := toRPS(p2)
	return score(p1Move, p2Move)
}

func main() {
	filename := "./data/day2"
	data, err := readLines(filename)
	check(err)

	var totalScore int = 0

	for _, element := range data {
		moves := strings.Split(element, " ")
		us := moves[1]
		them := moves[0]
		totalScore += determineWinner(us, them)
	}

	fmt.Println("total score:", totalScore)

}
