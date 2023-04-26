package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var fileName = "input.txt"

func main() {
	// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
	part1 := getHighestCalories()
	fmt.Println("part1 answer =", part1)

	// Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
	part2 := getThreeHighestCalories()
	fmt.Println("part2 answer =", part2)
}

func getHighestCalories() int {
	var sum, totalSum = 0, 0

	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()

		// empty line denotes the end of a particular elf's sum
		if str == "" {
			if sum > totalSum {
				totalSum = sum
			}
			// reset sum for the next elf
			sum = 0
		} else {
			num := getNumFromString(str)
			sum += num
		}
	}

	return totalSum
}

func getThreeHighestCalories() int {
	var sum, totalSum = 0, 0
	var threeSums = []int{0, 0, 0}

	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()

		if str == "" {
			if sum >= threeSums[0] {
				// shift sums by one to make room
				threeSums[2] = threeSums[1]
				threeSums[1] = threeSums[0]
				// set highest sum to current sum
				threeSums[0] = sum
			} else if sum >= threeSums[1] {
				threeSums[2] = threeSums[1]
				threeSums[1] = sum
			} else if sum >= threeSums[2] {
				threeSums[2] = sum
			}
			sum = 0
		} else {
			num := getNumFromString(str)
			sum += num
		}
	}

	for _, count := range threeSums {
		totalSum += count
	}

	return totalSum
}

func getNumFromString(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
