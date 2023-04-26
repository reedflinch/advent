package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var fileName = "input.txt"

func main() {
	part1 := getStartOfPacket(4)
	fmt.Println("part1 =", part1)

	part2 := getStartOfPacket(14)
	fmt.Println("part2 =", part2)
}

func getStartOfPacket(length int) int {
	var str string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str = scanner.Text()
	}

	for i := 0; i < len(str)-length; i++ {
		if isValid(str[i:i+length], length) {
			return i + length
		}
	}

	return 0
}

func isValid(arr string, length int) bool {
	var seen = make(map[string]bool)

	for i := 0; i < length; i++ {
		if _, exists := seen[string(arr[i])]; exists {
			return false
		}
		seen[string(arr[i])] = true
	}
	return true
}
