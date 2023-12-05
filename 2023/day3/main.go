package main

// https://adventofcode.com/2023/day/3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var fileName = "input.txt"
var grid [][]rune

var checkDirections = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

type Position struct {
	Row int
	Col int
}

func isPotentialGear(char rune) bool {
	return char == '*'
}

func isSymbol(char rune) bool {
	symbols := []rune{'@', '#', '$', '%', '&', '*', '=', '+', '-', '/'}
	for _, symbol := range symbols {
		if char == symbol {
			return true
		}
	}
	return false
}

func isNumber(char rune) bool {
	return unicode.IsDigit(char)
}

func isInBounds(row, col int) bool {
	if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row]) {
		return true
	}
	return false
}

func extractNumberFromStart(row, col int) (int, int) {
	numStr := strings.Builder{}

	for j := col; j < len(grid[row]) && isNumber(grid[row][j]); j++ {
		numStr.WriteRune(grid[row][j])
	}

	num, _ := strconv.Atoi(numStr.String())
	return num, len(numStr.String())
}

func extractNumberFromAnywhere(row, col int) (int, int, int) {
	start, end := col, col
	numStr := string(grid[row][col])

	// look left for more number digits
	for j := col - 1; j >= 0 && isNumber(grid[row][j]); j-- {
		// prepend the digit
		numStr = string(grid[row][j]) + numStr
		start = j
	}

	// look right for more number digits
	for j := col + 1; j < len(grid[row]) && isNumber(grid[row][j]); j++ {
		// append the digit
		numStr += string(grid[row][j])
		end = j
	}
	num, _ := strconv.Atoi(numStr)
	return num, start, end
}

func getGearRatio(row, col int) int {
	visited := make(map[Position]bool)
	nums := []int{}

	// check adjacent cells for symbols
	for _, dir := range checkDirections {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if isInBounds(newRow, newCol) && isNumber(grid[newRow][newCol]) && !visited[Position{newRow, newCol}] {
			fullNum, startCol, endCol := extractNumberFromAnywhere(newRow, newCol)
			nums = append(nums, fullNum)

			// mark all found digit positions as visited
			for j := startCol; j <= endCol; j++ {
				visited[Position{newRow, j}] = true
			}
		}
	}

	if len(nums) != 2 {
		return 0
	}
	return nums[0] * nums[1]
}

func hasAdjacentSymbols(row, col, length int) bool {
	// check adjacent cells for symbols
	for _, dir := range checkDirections {
		// look far enough to the right to account for each digit
		for l := 0; l < length; l++ {
			newRow := row + dir[0]
			newCol := col + dir[1] + l

			if isInBounds(newRow, newCol) && isSymbol(grid[newRow][newCol]) {
				return true
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}

	visited := make(map[Position]bool)
	partNumbersSum, gearRatiosSum := 0, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			char := grid[i][j]

			// find the first digit of a number that hasn't yet been visited
			if isNumber(char) && !visited[Position{i, j}] {
				// get the full number, as it may have multiple digits
				number, length := extractNumberFromStart(i, j)

				if hasAdjacentSymbols(i, j, length) {
					partNumbersSum += number
				}

				// mark the location of each digit as visited
				for l := j; l < j+length; l++ {
					visited[Position{i, l}] = true
				}
			} else if isPotentialGear(char) {
				gearRatiosSum += getGearRatio(i, j)
			}
		}
	}

	fmt.Printf("Sum of numbers with adjacent symbols (part 1): %v\n", partNumbersSum)
	fmt.Printf("Sum of gear ratios (part 2): %v\n", gearRatiosSum)
}
