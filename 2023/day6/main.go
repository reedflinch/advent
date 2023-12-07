package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/6

var fileName = "input.txt"

type race struct {
	time, dist int
}

func main() {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")

	answer := 1
	races := parseLines(lines)
	for _, race := range races {
		answer *= waysToWin(race)
	}

	fmt.Printf("product of ways to win = %v\n", answer)
}

func parseLines(lines []string) (races []race) {
	timeLine := strings.Fields(lines[0])
	distLine := strings.Fields(lines[1])

	for i := 1; i < len(timeLine); i++ {
		time := mustAtoi(strings.TrimSpace(timeLine[i]))
		dist := mustAtoi(strings.TrimSpace(distLine[i]))
		races = append(races, race{time, dist})
	}

	return races
}
func waysToWin(r race) int {
	answer := 0

	// relies on the fact that the distance distribution
	// across all possible charging times is symmetrical
	middle := r.time/2 + 1

	for chargeTime := middle; chargeTime <= r.time; chargeTime++ {
		timeLeft := r.time - chargeTime
		d := timeLeft * chargeTime

		if d > r.dist {
			answer++
		} else {
			break
		}
	}

	answer *= 2
	if r.time%2 == 0 {
		answer++
	}
	return answer
}

func mustAtoi(s string) int {
	if num, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return num
	}
}
