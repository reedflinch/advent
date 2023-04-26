package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileName = "input.txt"

func main() {
	part1, part2 := countPairs()
	fmt.Println("part1 answer =", part1)
	fmt.Println("part2 answer =", part2)
}

func countPairs() (int, int) {
	var contained, overlapping = 0, 0

	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := scanner.Text()
		pair := strings.Split(str, ",")
		mins, maxes := getPairs(pair)
		if doesContainTheOther(mins, maxes) {
			contained++
		}
		if doesOverlap(mins, maxes) {
			overlapping++
		}
	}

	return contained, overlapping
}

func doesOverlap(mins, maxes []int) bool {
	if mins[1] <= maxes[0] && mins[1] >= mins[0] {
		return true
	} else if mins[0] <= maxes[1] && mins[0] >= mins[1] {
		return true
	}

	return false
}

func getPairs(pair []string) ([]int, []int) {
	var mins []int
	var maxes []int

	for _, assignment := range pair {
		nums := strings.Split(assignment, "-")
		if len(nums) != 2 {
			panic(fmt.Errorf("pair not found"))

		}

		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])
		mins = append(mins, min)
		maxes = append(maxes, max)
	}

	return mins, maxes
}

func doesContainTheOther(mins, maxes []int) bool {
	if mins[0] <= mins[1] && maxes[0] >= maxes[1] {
		return true
	} else if mins[0] >= mins[1] && maxes[0] <= maxes[1] {
		return true
	}

	return false
}
