package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fileName = "../input.txt"
var theirPlayMap = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}
var neededOutcomeMap = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}
var playScoreMap = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}
var outcomeScoreMap = map[string]int{
	"win":  6,
	"draw": 3,
	"loss": 0,
}
var losesTo = map[string]string{
	"rock":     "scissors",
	"paper":    "rock",
	"scissors": "paper",
}
var beats = map[string]string{
	"rock":     "paper",
	"paper":    "scissors",
	"scissors": "rock",
}

func main() {
	part2 := getHighScore()
	fmt.Println("part2 answer =", part2)
}

func getHighScore() int {
	var outcomeScore = 0
	var myTotalScore = 0
	var myPlay string

	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		plays := strings.Split(str, " ")
		if len(plays) < 2 {
			fmt.Println("found less than 2 plays. exiting....")
			os.Exit(1)
		}

		theirPlaySymbol := plays[0]
		neededOutcomeSymbol := plays[1]

		theirPlay := theirPlayMap[theirPlaySymbol]
		neededOutcome := neededOutcomeMap[neededOutcomeSymbol]
		outcomeScore = outcomeScoreMap[neededOutcome]

		if neededOutcome == "draw" {
			myPlay = theirPlay
		} else if neededOutcome == "lose" {
			myPlay = losesTo[theirPlay]
		} else if neededOutcome == "win" {
			myPlay = beats[theirPlay]
		}

		myPlayScore := playScoreMap[myPlay]
		myTotalScore += myPlayScore + outcomeScore
	}

	return myTotalScore
}
