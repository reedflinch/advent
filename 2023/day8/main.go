package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// https://adventofcode.com/2023/day/8

var fileName = "input.txt"

const (
	startNode = "AAA"
	endNode   = "ZZZ"
)

type Node struct {
	left, right string
}

var nodeMap = make(map[string]*Node)

func main() {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content := string(bytes)
	content = strings.TrimSpace(content)
	lines := strings.Split(content, "\n")

	instructions := strings.Split(lines[0], "")
	nodes := lines[2:]

	processNodes(nodes)

	if _, ok := nodeMap[startNode]; !ok {
		log.Fatal("no start node")
	}
	if _, ok := nodeMap[endNode]; !ok {
		log.Fatal("no finish node")
	}

	stepsTaken := stepsOnNormalPath(instructions)
	fmt.Printf("answer (part 1) = %v\n", stepsTaken)

	ghostStepsTaken := stepsOnParallelGhostPath(instructions)
	fmt.Printf("answer (part 2) = %v\n", ghostStepsTaken)
}

// part 1
func stepsOnNormalPath(instructions []string) int {
	index, steps := 0, 0
	currentNode := startNode

	for currentNode != endNode {
		if instructions[index] == "L" {
			currentNode = nodeMap[currentNode].left
		} else if instructions[index] == "R" {
			currentNode = nodeMap[currentNode].right
		} else {
			log.Fatalf("could not follow instruction: %v\n", instructions[index])
		}

		steps++
		index = (index + 1) % len(instructions)
	}

	return steps
}

func isGhostStartNode(name string) bool {
	return name[2] == 'A'
}

func isGhostEndNode(name string) bool {
	return name[2] == 'Z'
}

func ghostStartNodes() []string {
	var nodes []string

	for node := range nodeMap {
		if isGhostStartNode(node) {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

// part 2
// got some help on part 2: https://github.com/busser/adventofcode/blob/main/y2023/d08/solution.go
func stepsOnParallelGhostPath(instructions []string) int {
	startNodes := ghostStartNodes()

	var pathLengths []int
	for _, node := range startNodes {
		pathLengths = append(pathLengths, stepsOnSingleGhostPath(instructions, node))
	}

	if len(pathLengths) == 1 {
		return pathLengths[0]
	}

	parallelPathLength := pathLengths[0]
	for i := 1; i < len(pathLengths); i++ {
		parallelPathLength = lcm(parallelPathLength, pathLengths[i])
	}

	return parallelPathLength
}

func lcm(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}

	return abs(a*b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func stepsOnSingleGhostPath(instructions []string, start string) int {
	index, steps := 0, 0
	currentNode := start

	for !isGhostEndNode(currentNode) {
		if instructions[index] == "L" {
			currentNode = nodeMap[currentNode].left
		} else if instructions[index] == "R" {
			currentNode = nodeMap[currentNode].right
		} else {
			log.Fatalf("could not follow instruction: %v\n", instructions[index])
		}

		steps++
		index = (index + 1) % len(instructions)
	}

	return steps
}

func processNodes(nodes []string) {
	for _, node := range nodes {
		parts := strings.Split(node, " = ")
		name := parts[0]

		directions := strings.Fields(parts[1])
		left := strings.TrimFunc(directions[0], trimPunct)
		right := strings.TrimFunc(directions[1], trimPunct)

		nodeMap[name] = &Node{left, right}
	}
}

func trimPunct(r rune) bool {
	return unicode.IsPunct(r)
}
