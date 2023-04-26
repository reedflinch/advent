package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type worryLevelFunc func(int) int

type testFunc func(int) bool

type Monkey struct {
	name        int
	items       []int
	operation   worryLevelFunc
	test        testFunc
	trueResult  int
	falseResult int
	inspected   int
}

var monkeys []Monkey

func main() {
	monkeys = []Monkey{}
	part1 := getMonkeyBusiness(20, true)
	fmt.Println("part1 =", part1)
}

func getMonkeyBusiness(rounds int, isPart1 bool) int {
	runRounds(rounds, isPart1)

	var mostActiveMonkeys = make([]int, 2)

	for _, monkey := range monkeys {
		if monkey.inspected > mostActiveMonkeys[0] {
			mostActiveMonkeys[1] = mostActiveMonkeys[0]
			mostActiveMonkeys[0] = monkey.inspected
		} else if monkey.inspected > mostActiveMonkeys[1] {
			mostActiveMonkeys[1] = monkey.inspected
		}
	}

	return mostActiveMonkeys[0] * mostActiveMonkeys[1]
}

func runRounds(rounds int, isPart1 bool) {
	monkeys = getMonkeys()

	for i := 0; i < rounds; i++ {
		processRound(monkeys, isPart1)
	}
}

func processRound(monks []Monkey, isPart1 bool) {
	for i, monkey := range monks {
		for len(monkeys[i].items) > 0 {
			item := monkey.items[0]

			monkey.inspected++
			worryLevel := monkey.operation(item)
			if isPart1 {
				worryLevel /= 3
			}
			test := monkey.test(worryLevel)

			if test {
				monks[monkey.trueResult].items = append(monks[monkey.trueResult].items, worryLevel)
			} else {
				monks[monkey.falseResult].items = append(monks[monkey.falseResult].items, worryLevel)
			}

			monkey.items = monkey.items[1:]
			monkeys[i] = monkey
		}
	}
}

func getMonkeys() (monkeys []Monkey) {
	var monkey []string

	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// reached end of monkey block, time to parse
		if line == "" {
			monkeys = append(monkeys, parseMonkey(monkey))
			monkey = []string{}
		} else {
			monkey = append(monkey, line)
		}
	}

	return monkeys
}

func parseMonkey(in []string) (out Monkey) {
	if len(in) < 6 {
		fmt.Println("INVALID INPUT WTF")
	}

	out.name = parseName(in[0])
	out.items = parseItems(in[1])
	out.operation = parseOperation(in[2])
	out.test = parseTest(in[3])
	out.trueResult = parseTrueResult(in[4])
	out.falseResult = parseFalseResult(in[5])

	return out
}

func parseFalseResult(in string) int {
	suffix := strings.TrimPrefix(strings.TrimSpace(in), "If false: throw to monkey ")
	value, _ := strconv.Atoi(suffix)
	return value
}

func parseTrueResult(in string) int {
	suffix := strings.TrimPrefix(strings.TrimSpace(in), "If true: throw to monkey ")
	value, _ := strconv.Atoi(suffix)
	return value
}

func parseTest(in string) testFunc {
	suffix := strings.TrimPrefix(strings.TrimSpace(in), "Test: divisible by ")
	value, _ := strconv.Atoi(suffix)

	return func(level int) bool {
		if level%value == 0 {
			return true
		}
		return false
	}
}

func isAddition(sign string) bool {
	if sign == "*" {
		return false
	}
	return true
}

func parseOperation(in string) worryLevelFunc {
	var value = 0
	var isOld = false

	suffix := strings.TrimPrefix(strings.TrimSpace(in), "Operation: new = old ")
	elements := strings.Split(suffix, " ")
	if elements[1] == "old" {
		isOld = true
	} else {
		value, _ = strconv.Atoi(elements[1])
	}

	if isAddition(elements[0]) {
		return func(old int) int {
			return old + value
		}
	} else if isOld {
		return func(old int) int {
			return old * old
		}
	}

	return func(old int) int {
		return old * value
	}
}

func parseItems(in string) (out []int) {
	items := strings.TrimPrefix(strings.TrimSpace(in), "Starting items: ")
	itemList := strings.Split(items, ", ")

	for _, item := range itemList {
		level, _ := strconv.Atoi(item)
		out = append(out, level)
	}

	return out
}

func parseName(in string) int {
	parts := strings.Split(in, " ")
	name, _ := strconv.Atoi(string(parts[1][0]))
	return name
}
