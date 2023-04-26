package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileName = "input.txt"

func main() {
	part1 := getIndividualPriorities()
	fmt.Println("part1 answer =", part1)

	part2 := getGroupPriorities()
	fmt.Println("part2 answer =", part2)
}

func getGroupPriorities() int {
	var count, sum, priority = 0, 0, 0
	var group = []string{}

	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()

		group = append(group, str)
		count++

		if count == 3 {
			commonChar := getGroupCommonChar(group)
			priority = charToPriority(commonChar)
			sum += priority
			group = []string{}
			count = 0
		}
	}

	return sum
}

func getIndividualPriorities() int {
	var sum, priority = 0, 0

	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		first := str[:len(str)/2]
		second := str[len(str)/2:]

		commonChar := getIndividualCommonChar(first, second)
		priority = charToPriority(commonChar)
		sum += priority
	}

	return sum
}

func getGroupCommonChar(group []string) byte {
	if len(group) != 3 {
		panic(fmt.Errorf("group of 3 not found"))
	}

	var groupChars = make(map[byte]int)
	var elfChars = make(map[byte]struct{})

	for _, elf := range group {

		for j := 0; j < len(elf); j++ {
			item := elf[j]
			_, inElfChars := elfChars[item]
			if !inElfChars {
				groupChars[item]++
			}
			elfChars[item] = struct{}{}
		}

		elfChars = make(map[byte]struct{})
	}

	for key, value := range groupChars {
		if value == 3 {
			return key
		}
	}

	return 0
}

func getIndividualCommonChar(a, b string) byte {
	var charsInA = make(map[byte]struct{})
	for i := 0; i < len(a); i++ {
		charsInA[a[i]] = struct{}{}
	}
	for j := 0; j < len(b); j++ {
		if _, exists := charsInA[b[j]]; exists {
			return b[j]
		}
	}

	return 0
}

func charToPriority(char byte) int {
	value := char - 'a' + 1
	if value > 26 {
		value -= 198
	}
	return int(value)
}
