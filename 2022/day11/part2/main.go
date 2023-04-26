package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type worryLevelFunc func(int64) int64

type testFunc func(int64) bool

type Monkey struct {
	name        int
	items       []int64
	operation   worryLevelFunc
	test        testFunc
	trueResult  int
	falseResult int
	inspected   int64
}

var monkeys []Monkey

func main() {
	monkeys = []Monkey{}
	part2 := getMonkeyBusiness(10000)
	fmt.Println("part2 =", part2)
}

func getMonkeyBusiness(rounds int) int64 {
	runRounds(rounds)

	var mostActiveMonkeys = make([]int64, 2)

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

func runRounds(rounds int) {
	monkeys = getMonkeys()

	for i := 0; i < rounds; i++ {
		processRound(monkeys)
	}
}

func processRound(monks []Monkey) {
	for i, monkey := range monks {
		for len(monkeys[i].items) > 0 {
			item := monkey.items[0]

			monkey.inspected++
			worryLevel := monkey.operation(item)
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

	return func(level int64) bool {
		if level%int64(value) == 0 {
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
		return func(old int64) int64 {
			return old + int64(value)
		}
	} else if isOld {
		return func(old int64) int64 {
			return old * old
		}
	}

	return func(old int64) int64 {
		return old * int64(value)
	}
}

func parseItems(in string) (out []int64) {
	items := strings.TrimPrefix(strings.TrimSpace(in), "Starting items: ")
	itemList := strings.Split(items, ", ")

	for _, item := range itemList {
		level, _ := strconv.Atoi(item)
		out = append(out, int64(level))
	}

	return out
}

func parseName(in string) int {
	parts := strings.Split(in, " ")
	name, _ := strconv.Atoi(string(parts[1][0]))
	return name
}
