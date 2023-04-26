package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	cycles, delta int
}

var crt = make([]string, 240)

func main() {
	part1 := getSignalStrengthSums()
	fmt.Println("part1 answer =", part1)

	fmt.Println("part2 answer =")
	print(crt)
}

func getSignalStrengthSums() int {
	var cycles = 0
	var register = 1
	var sum = 0

	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inst := parseInstruction(scanner.Text())

		for i := 0; i < inst.cycles; i++ {
			pixel := cycles % 40
			if pixel >= register-1 && pixel <= register+1 {
				crt[cycles] = "#"
			} else {
				crt[cycles] = "."
			}

			cycles++

			if isInterestingCycle(cycles) {
				fmt.Printf("adding %d to sum %d on cycle %d\n", register, sum, cycles)
				sum += register * cycles
			}
		}

		register += inst.delta
	}

	return sum
}

func isInterestingCycle(cycles int) bool {
	return cycles%40 == 20
}

func parseInstruction(str string) instruction {
	var cycles int
	var toAdd = 0

	parts := strings.Split(str, " ")

	op := parts[0]
	if op == "addx" {
		cycles = 2
	} else if op == "noop" {
		cycles = 1
	}

	if len(parts) > 1 {
		toAddStr := parts[1]
		toAdd, _ = strconv.Atoi(toAddStr)
	}

	return instruction{
		cycles: cycles,
		delta:  toAdd,
	}
}

func print(crt []string) {
	for i := 40; i <= 240; i += 40 {
		fmt.Println(strings.Join(crt[i-40:i], ""))
	}
}
