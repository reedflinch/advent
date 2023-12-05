package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/4

var fileName = "input.txt"
var cardPile = make(map[int]*Card)

type Card struct {
	id             int
	count          int
	winning, mine  []string
	matches, score int
}

func main() {
	matchScore := 0

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(file)

	// part 1
	for s.Scan() {
		card := parseCard(s.Text())
		card.countMatches()
		card.calculateScore()
		card.save()
		matchScore += card.score
	}
	fmt.Printf("match score (part 1) = %v\n\n", matchScore)

	// part 2
	processCopies()
	fmt.Printf("total cards (part 2) = %v\n\n", countCards())
}

func parseCard(s string) Card {
	parts := strings.Split(s, ":")

	idStr := strings.TrimPrefix(parts[0], "Card ")
	dataParts := strings.Split(parts[1], "|")

	winning := strings.Fields(dataParts[0])
	mine := strings.Fields(dataParts[1])

	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	card := Card{
		id:      id,
		winning: winning,
		mine:    mine,
		count:   1,
	}

	cardPile[id] = &card
	return card
}

// part 1
func (c *Card) countMatches() {
	mineMap := make(map[string]struct{})

	for _, num := range c.mine {
		mineMap[num] = struct{}{}
	}

	for _, num := range c.winning {
		if _, ok := mineMap[num]; ok {
			c.matches++
		}
	}
}

func (c *Card) calculateScore() {
	if c.matches == 0 {
		c.score = 0
	} else {
		c.score = int(math.Pow(float64(2), float64((c.matches - 1))))
	}
}

func countCards() int {
	sum := 0
	for _, card := range cardPile {
		sum += card.count
	}
	return sum
}

// part 2
func (c *Card) save() {
	cardPile[c.id] = c
}

func processCopies() {
	for i := 0; i < len(cardPile); i++ {
		if card, ok := cardPile[i+1]; ok {
			offset := card.id + 1

			for i := offset; i < card.matches+offset; i++ {
				cardPile[i].count += card.count
			}
		}
	}
}
