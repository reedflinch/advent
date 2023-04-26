package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var fileName = "input.txt"

type node struct {
	name     string
	size     int
	children []node
}

func main() {
	getLargestDirsSum()
}

// had some help
// https://github.com/aarneng/AdventOfCode2022/blob/main/day7/main.go
func getLargestDirsSum() {
	var (
		sum = 0

		spaceAvailable = 70000000
		spaceRequired  = 30000000
		dirToDelete    = math.MaxInt

		sizes   = map[string]int{"/": 0}
		dirPath = []string{"/"}
	)

	// pre-processing
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// pre-processing

	for scanner.Scan() {
		str := scanner.Text()

		switch {
		// we're counting file sizes, so continue
		case strings.HasPrefix(str, "$ ls"):
			continue
		// nothing to do for a dir listing, so continue
		case strings.HasPrefix(str, "dir "):
			continue
		// return to root
		case str == "$ cd /":
			dirPath = []string{"/"}
			continue
		// go up one level
		case str == "$ cd ..":
			dirPath = dirPath[:len(dirPath)-1]
			continue
		// change into new directory
		case strings.HasPrefix(str, "$ cd "):
			dir := str[4:] + "/"
			dirPath = append(dirPath, dir)
			continue
		}

		// if no other conditions, we're now counting file sizes
		currDirPath := "/"
		size := getFileSize(str)
		// count all sizes as part of root
		sizes[currDirPath] += size

		// we know the first dirPath is root, and that's already captured in currDirPath
		// so start at the second element
		for _, dir := range dirPath[1:] {
			currDirPath += string(dir)
			sizes[currDirPath] += size
		}
	}

	for _, dirSize := range sizes {
		if dirSize <= 100000 {
			sum += dirSize
		}
	}

	fmt.Println("part1 answer =", sum)

	// disk size minus root size is what's left
	unusedSpace := spaceAvailable - sizes["/"]
	spaceStillNeeded := spaceRequired - unusedSpace

	for _, size := range sizes {
		if size >= spaceStillNeeded && size < dirToDelete {
			dirToDelete = size
		}
	}

	fmt.Println("part2 answer =", dirToDelete)
}

func getFileSize(s string) int {
	parts := strings.Split(s, " ")
	size, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	return size
}
