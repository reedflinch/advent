package main

// https://adventofcode.com/2023/day/2

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type game struct {
	id   int
	sets []Set
}

type Set struct {
	Red, Blue, Green int
}

var fileName = "input.txt"
var possibleCount, totalPower = 0, 0

var (
	blueLimit  = 14
	redLimit   = 12
	greenLimit = 13
)

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		game := parseLine(s.Text())

		if game.isPossible() {
			possibleCount += game.id
		}
		totalPower += game.getPower()
	}

	fmt.Printf("part 1 answer = %v\n", possibleCount)
	fmt.Printf("part 2 answer = %v\n", totalPower)
}

func (g *game) getPower() int {
	var bMax, gMax, rMax = 0, 0, 0

	for _, set := range g.sets {
		if set.Blue > bMax {
			bMax = set.Blue
		}
		if set.Green > gMax {
			gMax = set.Green
		}
		if set.Red > rMax {
			rMax = set.Red
		}
	}

	return bMax * gMax * rMax
}

func (g *game) isPossible() bool {
	for _, set := range g.sets {
		if set.Blue > blueLimit || set.Green > greenLimit || set.Red > redLimit {
			return false
		}
	}
	return true
}

func parseLine(line string) game {
	parts := strings.Split(line, ":")

	return game{
		id:   getID(parts[0]),
		sets: getSets(parts[1]),
	}
}

func (s *Set) setProperty(name string, value int) *Set {
	reflect.ValueOf(s).Elem().FieldByName(name).Set(reflect.ValueOf(value))
	return s
}

func getSet(s string) Set {
	set := Set{}

	sets := strings.Split(strings.TrimSpace(s), ",")
	for _, part := range sets {
		parts := strings.Split(strings.TrimSpace(part), " ")
		count := parts[0]
		color := parts[1]

		countNum, _ := strconv.Atoi(count)
		set = *set.setProperty(strings.Title(color), countNum)
	}

	return set
}

func getSets(s string) []Set {
	sets := []Set{}

	parts := strings.Split(s, ";")
	for _, part := range parts {
		sets = append(sets, getSet(part))
	}

	return sets
}

func getID(s string) int {
	var id = strings.Builder{}

	for _, char := range s {
		if unicode.IsDigit(char) {
			id.WriteRune(char)
		}
	}

	idNum, _ := strconv.Atoi(id.String())
	return idNum
}
