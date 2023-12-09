package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var fileName = "input.txt"

type Hand struct {
	cards      []string
	bid, score int
}

func main() {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content := string(bytes)
	content = strings.TrimSpace(content)
	lines := strings.Split(content, "\n")

	run(1, lines)
	run(2, lines)
}

func run(part int, lines []string) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		hand := Hand{
			cards: strings.Split(parts[0], ""),
			bid:   mustAtoi(parts[1]),
		}

		if part == 1 {
			hand.classify1()
		} else {
			hand.classify2()
		}
		hands[i] = hand
	}

	hands = sortHands(hands, part)
	answer := 0
	for i, hand := range hands {
		answer += hand.bid * (i + 1)
	}

	fmt.Printf("answer (part %v) = %v\n", part, answer)
}

func sortHands(hands []Hand, part int) []Hand {
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].score == hands[j].score {
			return compareCards(hands[i].cards, hands[j].cards, part)
		}
		return hands[i].score < hands[j].score
	})
	return hands
}

func compareCards(h1, h2 []string, part int) bool {
	for i := 0; i < len(h1); i++ {
		if cardRank(h1[i], part) == cardRank(h2[i], part) {
			continue
		}
		return cardRank(h1[i], part) < cardRank(h2[i], part)
	}
	return false
}

func (h *Hand) classify1() {
	counts := make(map[string]int)
	for _, card := range h.cards {
		counts[card]++
	}

	highestCount := 0
	secondHighestCount := 0
	for _, count := range counts {
		if count > highestCount {
			secondHighestCount = highestCount
			highestCount = count
		} else if count > secondHighestCount {
			secondHighestCount = count
		}
	}

	switch {
	case highestCount == 5: // Five of a kind
		h.score = 7
	case highestCount == 4: // Four of a kind
		h.score = 6
	case highestCount == 3 && secondHighestCount == 2: // Full house
		h.score = 5
	case highestCount == 3: // Three of a kind
		h.score = 4
	case highestCount == 2 && secondHighestCount == 2: // Two pair
		h.score = 3
	case highestCount == 2: // One pair
		h.score = 2
	case highestCount == 1: // High card
		h.score = 1
	default:
		log.Fatal("could not classify card")
	}
}

func (h *Hand) classify2() {
	counts := make(map[string]int)
	jacks := 0

	for _, card := range h.cards {
		if card == "J" {
			jacks++
		} else {
			counts[card]++
		}
	}

	highestCount := 0
	secondHighestCount := 0
	for _, count := range counts {
		if count > highestCount {
			secondHighestCount = highestCount
			highestCount = count
		} else if count > secondHighestCount {
			secondHighestCount = count
		}
	}

	switch {
	case jacks == 5:
		h.score = 7
	case highestCount == 5: // Five of a kind
		h.score = 7
	case highestCount == 4: // Four of a kind
		if jacks == 1 {
			h.score = 7 // five
		} else {
			h.score = 6
		}
	case highestCount == 3 && secondHighestCount == 2: // Full house
		h.score = 5
	case highestCount == 3: // Three of a kind
		if jacks == 2 {
			h.score = 7 // five
		} else if jacks == 1 {
			h.score = 6 // four
		} else {
			h.score = 4
		}
	case highestCount == 2 && secondHighestCount == 2: // Two pair
		if jacks == 1 {
			h.score = 5 // full house
		} else {
			h.score = 3
		}
	case highestCount == 2: // One pair
		if jacks == 3 {
			h.score = 7 // five
		} else if jacks == 2 {
			h.score = 6 // four
		} else if jacks == 1 {
			h.score = 4 // three
		} else {
			h.score = 2
		}
	case highestCount == 1: // High card
		if jacks == 4 {
			h.score = 7
		} else if jacks == 3 {
			h.score = 6
		} else if jacks == 2 {
			h.score = 4
		} else if jacks == 1 {
			h.score = 2
		} else {
			h.score = 1
		}
	default:
		log.Fatalf("could not classify hand %+v\n", h)
	}
}

func cardRank(card string, part int) int {
	var rankedCards []string
	if part == 1 {
		rankedCards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	} else {
		rankedCards = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	}

	for i, c := range rankedCards {
		if card == c {
			return i
		}
	}
	return -1
}

func mustAtoi(s string) int {
	if num, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return num
	}
}
