package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var fileName = "input.txt"

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	sum1, sum2 := 0, 0
	for s.Scan() {
		fmt.Printf("line = %v\n", s.Text())

		sum1 += getLineValue(s.Text(), 1)
		sum2 += getLineValue(s.Text(), 2)
	}

	fmt.Printf("part 1 answer = %v\n", sum1)
	fmt.Printf("part 2 answer = %v\n", sum2)
}

func getLineValue(l string, part int) int {
	var leftStr, rightStr string

	if part == 1 {
		leftStr, rightStr = getOuterDigitsPart1(l)
	} else if part == 2 {
		leftStr, rightStr = getOuterDigitsPart2(l)
	}

	left, _ := strconv.Atoi(leftStr)
	right, _ := strconv.Atoi(rightStr)

	return (left * 10) + right
}

func getOuterDigitsPart1(line string) (string, string) {
	digits := []string{}

	for i := 0; i < len(line); i++ {
		char := line[i]

		// store digits
		if unicode.IsDigit(rune(char)) {
			digits = append(digits, string(char))
		}
	}

	return digits[0], digits[len(digits)-1]
}

func getOuterDigitsPart2(line string) (string, string) {
	digits := []string{}

	for i := 0; i < len(line); i++ {
		char := line[i]

		// store digit
		if unicode.IsDigit(rune(char)) {
			digits = append(digits, string(char))
			// otherwise, check if the digit is spelled
		} else {
			// starting at the current character, advance to the right to find a spelled digit
			for j := i + 1; j < len(line); j++ {
				// if we hit a real digit in the substring, we know it's not spelled
				if unicode.IsDigit(rune(line[j])) {
					break
				}

				// store spelled digit
				if digit, ok := digitMap[line[i:j+1]]; ok {
					digits = append(digits, digit)
				}
			}
		}
	}

	return digits[0], digits[len(digits)-1]
}
