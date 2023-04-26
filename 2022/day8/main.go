package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var fileName = "input.txt"

func main() {
	part1, part2 := getVisibleTrees()
	fmt.Println("part1 answer =", part1)
	fmt.Println("part2 answer =", part2)
}

func getVisibleTrees() (visibleTrees, scenicScore int) {
	var trees = getTrees()
	var cols = getColumns(trees)

	for i := 0; i < len(trees); i++ {
		row := trees[i]

		for j := 0; j < len(row); j++ {
			// part 1
			height := row[j]

			// add edge trees, they're all visible
			if i == 0 || j == 0 || i == len(trees)-1 || j == len(row)-1 {
				visibleTrees++
				// right
			} else if isVisible(height, row[j+1:]) {
				visibleTrees++
				// left
			} else if isVisible(height, row[:j]) {
				visibleTrees++
				// down
			} else if isVisible(height, cols[j][i+1:]) {
				visibleTrees++
				// up
			} else if isVisible(height, cols[j][:i]) {
				visibleTrees++
			}

			// part 2
			right := getNumberVisible(height, row[j+1:], false)
			left := getNumberVisible(height, row[:j], true)
			down := getNumberVisible(height, cols[j][i+1:], false)
			up := getNumberVisible(height, cols[j][:i], true)

			score := right * left * down * up
			if score > scenicScore {
				scenicScore = score
			}
		}
	}

	return visibleTrees, scenicScore
}

func getNumberVisible(value int, rowOrColumn []int, reverse bool) (count int) {
	if !reverse {
		for i := 0; i < len(rowOrColumn); i++ {
			count++
			if rowOrColumn[i] >= value {
				return count
			}
		}
	} else {
		for i := len(rowOrColumn) - 1; i >= 0; i-- {
			count++
			if rowOrColumn[i] >= value {
				return count
			}
		}
	}

	return count
}

func isVisible(value int, rowOrColumn []int) bool {
	for i := 0; i < len(rowOrColumn); i++ {
		if rowOrColumn[i] >= value {
			return false
		}
	}

	return true
}

func getColumns(trees [][]int) (cols [][]int) {
	cols = make([][]int, len(trees[0]))

	for i := 0; i < len(trees[0]); i++ {
		cols[i] = make([]int, 0)

		for j := 0; j < len(trees); j++ {
			cols[i] = append(cols[i], trees[j][i])
		}
	}

	return cols
}

func getTrees() (trees [][]int) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		trees = append(trees, stringToInts(scanner.Text()))
	}

	return trees
}

func stringToInts(s string) (ints []int) {
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		ints = append(ints, num)
	}

	return ints
}
