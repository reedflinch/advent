package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/9

var fileName = "input.txt"

func main() {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content := string(bytes)
	content = strings.TrimSpace(content)
	lines := strings.Split(content, "\n")

	endSum, beginningSum := 0, 0
	for _, line := range lines {
		diffSet := processLine(line)

		endSum += sumFromDiffSetEnd(diffSet)
		beginningSum += sumFromDiffSetBeginning(diffSet)
	}

	fmt.Printf("answer (part 1) = %v\n", endSum)
	fmt.Printf("answer (part 2) = %v\n", beginningSum)
}

func sumFromDiffSetBeginning(d [][]int) int {
	prevValue := 0

	// start from the bottom of the diffSet
	for i := len(d) - 1; i >= 0; i-- {
		sequence := d[i]
		firstValue := sequence[0]
		prevValue = firstValue - prevValue
	}

	return prevValue
}

func sumFromDiffSetEnd(d [][]int) int {
	sum := 0
	prevValue := 0

	// start from the bottom of the diffSet
	for i := len(d) - 1; i >= 0; i-- {
		sequence := d[i]
		lastValue := sequence[len(sequence)-1]
		sum += prevValue + lastValue
	}

	return sum
}

func processLine(line string) [][]int {
	diffSet := [][]int{}

	nums := numsFromLine(line)
	diffs := nums

	for !isZeroSequence(diffs) {
		diffSet = append(diffSet, diffs)
		diffs = lineDiffs(diffs)
	}

	return diffSet
}

func numsFromLine(line string) []int {
	chars := strings.Fields(line)
	nums := make([]int, len(chars))

	for i, c := range chars {
		nums[i] = mustAtoi(c)
	}

	return nums
}

func lineDiffs(nums []int) []int {
	diffs := make([]int, len(nums)-1)

	for i := 1; i < len(nums); i++ {
		diffs[i-1] = nums[i] - nums[i-1]
	}

	return diffs
}

func isZeroSequence(s []int) bool {
	for _, n := range s {
		if n != 0 {
			return false
		}
	}
	return true
}

func mustAtoi(s string) int {
	if num, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return num
	}
}
