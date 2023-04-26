package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tailMoves = 0
var seen = make(map[pair]bool)

type pair struct {
	long, lat int
}

func main() {
	part1 := getTailmoves()
	fmt.Println("part1 answer =", part1)
}

func getTailmoves() (moves int) {
	// head and tail start at the same place
	head := []int{0, 0}
	tail := []int{0, 0}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dir, dist := parseMove(scanner.Text())
		head = execHeadMove(head, dir, dist)
		tail = handleDiagonal(head, tail)
		tail = catchUp(head, tail)
	}

	return tailMoves
}

func isDiagonal(long, lat int) bool {
	long = abs(long)
	lat = abs(lat)

	if long > 1 && lat > 0 {
		return true
	} else if long > 0 && lat > 1 {
		return true
	}

	return false
}

func handleDiagonal(head, tail []int) []int {
	long := head[0] - tail[0]
	lat := head[1] - tail[1]

	isDiagonal := isDiagonal(long, lat)

	if !isDiagonal {
		return tail
	}

	if long >= 0 {
		tail[0]++
	} else {
		tail[0]--
	}
	if lat >= 0 {
		tail[1]++
	} else {
		tail[1]--
	}

	tailPair := pair{
		long: tail[0],
		lat:  tail[1],
	}

	if _, exists := seen[tailPair]; !exists {
		tailMoves++
	}
	seen[tailPair] = true

	return tail
}

func execTailMove(tail []int, coord, dist, dir int) []int {
	var tailPair pair

	// start at one to only catch up till touching
	for i := 1; i < abs(dist); i++ {
		tail[coord] += 1 * dir

		tailPair = pair{
			long: tail[0],
			lat:  tail[1],
		}

		if _, exists := seen[tailPair]; !exists {
			tailMoves++
		}
		seen[tailPair] = true
	}
	return tail
}

func catchUp(head, tail []int) []int {
	long := head[0] - tail[0]
	lat := head[1] - tail[1]

	if long > 1 {
		tail = execTailMove(tail, 0, long, 1)
	} else if long < -1 {
		tail = execTailMove(tail, 0, long, -1)
	} else if lat > 1 {
		tail = execTailMove(tail, 1, lat, 1)
	} else if lat < -1 {
		tail = execTailMove(tail, 1, lat, -1)
	}

	return tail
}

func parseMove(s string) (dir string, dist int) {
	parts := strings.Split(s, " ")
	dir = parts[0]
	dist, _ = strconv.Atoi(parts[1])
	return strings.ToLower(dir), dist
}

func execHeadMove(pos []int, dir string, dist int) []int {
	switch dir {
	case "r":
		pos[0] += dist
	case "l":
		pos[0] -= dist
	case "u":
		pos[1] += dist
	case "d":
		pos[1] -= dist
	}

	return pos
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
