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
var myPlayMap = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
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

func main() {
	part1 := getHighScore()
	fmt.Println("part1 answer =", part1)
}

func getHighScore() int {
	var outcomeScore = 0
	var myTotalScore = 0

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
		myPlaySymbol := plays[1]

		theirPlay := theirPlayMap[theirPlaySymbol]
		myPlay := myPlayMap[myPlaySymbol]
		myPlayScore := playScoreMap[myPlay]

		theirPlayThatWinsForMe := losesTo[myPlay]

		if theirPlay == myPlay {
			outcomeScore = outcomeScoreMap["draw"]
		} else if theirPlay == theirPlayThatWinsForMe {
			outcomeScore = outcomeScoreMap["win"]
		} else {
			outcomeScore = outcomeScoreMap["loss"]
		}

		myTotalScore += myPlayScore + outcomeScore
	}

	return myTotalScore
}
