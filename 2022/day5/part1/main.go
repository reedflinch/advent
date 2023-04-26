package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileName = "../input.txt"

// [G]                 [D] [R]
// [W]         [V]     [C] [T] [M]
// [L]         [P] [Z] [Q] [F] [V]
// [J]         [S] [D] [J] [M] [T] [V]
// [B]     [M] [H] [L] [Z] [J] [B] [S]
// [R] [C] [T] [C] [T] [R] [D] [R] [D]
// [T] [W] [Z] [T] [P] [B] [B] [H] [P]
// [D] [S] [R] [D] [G] [F] [S] [L] [Q]
//  1   2   3   4   5   6   7   8   9

type instruction struct {
	count, source, dest int
}

func main() {
	part1 := getTopOfStacks()
	fmt.Println("part1 answer =", part1)
}

func getTopOfStacks() string {
	var result string
	var stacks = map[int][]string{
		1: []string{"D", "T", "R", "B", "J", "L", "W", "G"},
		2: []string{"S", "W", "C"},
		3: []string{"R", "Z", "T", "M"},
		4: []string{"D", "T", "C", "H", "S", "P", "V"},
		5: []string{"G", "P", "T", "L", "D", "Z"},
		6: []string{"F", "B", "R", "Z", "J", "Q", "C", "D"},
		7: []string{"S", "B", "D", "J", "M", "F", "T", "R"},
		8: []string{"L", "H", "R", "B", "T", "V", "M"},
		9: []string{"Q", "P", "D", "S", "V"},
	}

	fmt.Printf("stacks = %+v\n", stacks)

	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := scanner.Text()
		instruction := parseInstruction(str)

		stacks = executeMove(instruction, stacks)
	}

	for i := 1; i <= len(stacks); i++ {
		result += getTopOfStack(stacks[i])
	}

	return result
}

func getTopOfStack(stack []string) string {
	return stack[len(stack)-1]
}

func executeMove(in *instruction, stacks map[int][]string) map[int][]string {
	for i := 1; i <= in.count; i++ {
		sourceStack := stacks[in.source]
		// get the last value (top) of the stack to move
		sourceValue := sourceStack[len(sourceStack)-1]

		sourceStack = sourceStack[:len(sourceStack)-1]
		destStack := stacks[in.dest]
		// move the sourceValue to the top of the destStack
		destStack = append(destStack, sourceValue)

		stacks[in.source] = sourceStack
		stacks[in.dest] = destStack
	}

	return stacks
}

func parseInstruction(i string) *instruction {
	strs := strings.Split(i, " ")
	if len(strs) != 6 {
		panic(fmt.Errorf("found instruction %v of invalid length\n", i))
	}

	count, _ := strconv.Atoi(strs[1])
	source, _ := strconv.Atoi(strs[3])
	dest, _ := strconv.Atoi(strs[5])

	in := &instruction{
		count:  count,
		source: source,
		dest:   dest,
	}

	return in
}
