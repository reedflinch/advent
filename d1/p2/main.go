package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	freqDelta, _ := fileToLines("input.txt")
	findRepeatedFrequency(freqDelta)
}

func findRepeatedFrequency(deltas []string) int64 {
	var absString string
	var absInt int64
	var sum int64
	var freqMap map[int64]int64
	freqMap = make(map[int64]int64)

	for {
		for i := range deltas {
			if strings.Contains(deltas[i], "-") {
				absString = strings.Replace(deltas[i], "-", "", -1)
				absInt, _ = strconv.ParseInt(absString, 10, 64)

				sum -= absInt

			} else if strings.Contains(deltas[i], "+") {
				absString = strings.Replace(deltas[i], "+", "", -1)
				absInt, _ = strconv.ParseInt(absString, 10, 64)

				sum += absInt
			}

			if i > 0 {
				if _, ok := freqMap[sum]; ok {
					fmt.Printf("Frequency %d repeated\n", sum)
					return sum
				}
			}

			freqMap[sum] = absInt
		}
	}
}

func fileToLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return lines, nil
}
