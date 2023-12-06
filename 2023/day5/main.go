package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// https://adventofcode.com/2023/day/5

type Seeder struct {
	Seed, Soil, Fertilizer, Water, Light, Temperature, Humidity, Location int64
}

type Entry struct {
	sourceStart, destStart, length int64
}

type Pair struct {
	src, dest string
}

var pairs = []Pair{
	{src: "Seed", dest: "Soil"},
	{src: "Soil", dest: "Fertilizer"},
	{src: "Fertilizer", dest: "Water"},
	{src: "Water", dest: "Light"},
	{src: "Light", dest: "Temperature"},
	{src: "Temperature", dest: "Humidity"},
	{src: "Humidity", dest: "Location"},
}

var fileName = "input.txt"
var seeds []Seeder

func main() {
	runPart(1)
	runPart(2)
}

func runPart(part int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)

	seeds = []Seeder{}
	seeds = append(seeds, getSeeds(s, part)...)

	for _, pair := range pairs {
		// fmt.Printf("processing pair %+v part %v\n", pair, part)
		processPair(s, pair)
	}

	fmt.Printf("lowest Location (part %v) = %v\n", part, getLowestLocation())
}

// 	for _, seed := range seeds {

// 	}

// func (s *Seeder) traverse() {

// }

func processPair(sc *bufio.Scanner, p Pair) {
	src, dest := p.src, p.dest

	headerName := fmt.Sprintf("%s-to-%s map:", strings.ToLower(src), strings.ToLower(dest))

	mapping := []Entry{}
	foundMap := false

	// extract data from input
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), headerName) {
			foundMap = true
			continue
		}

		isEntryLine := len(sc.Text()) > 0 && unicode.IsDigit(rune(sc.Text()[0]))
		if foundMap {
			if !isEntryLine {
				break
			} else {
				// populate current map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false
		srcValue := reflect.ValueOf(seed).FieldByName(src).Int()

		for _, entry := range mapping {
			if srcValue >= entry.sourceStart && srcValue <= entry.sourceStart+entry.length-1 {
				newValue := srcValue + (entry.destStart - entry.sourceStart)
				seeds[i].setProperty(dest, newValue)
				found = true
				break
			}
		}
		if !found {
			seeds[i].setProperty(dest, srcValue)
		}
	}
}

func getSeeds(sc *bufio.Scanner, part int) []Seeder {
	if part == 1 {
		fmt.Printf("getting seeds part 1...\n")
		return getSeedsPart1(sc)
	}
	fmt.Printf("getting seeds part 2...\n")
	return getSeedsPart2(sc)
}

func getSeedsPart1(sc *bufio.Scanner) []Seeder {
	line := ""
	allSeeds := []Seeder{}

	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "seeds:") {
			line = sc.Text()
			break
		}
	}

	seedList := strings.TrimSpace(strings.TrimPrefix(line, "seeds:"))
	seedStrs := strings.Fields(seedList)

	for _, seedStr := range seedStrs {
		id := parseInt(seedStr)
		allSeeds = append(allSeeds, Seeder{Seed: id})
	}

	return allSeeds
}

func getSeedsPart2(sc *bufio.Scanner) []Seeder {
	line := ""
	allSeeds := []Seeder{}

	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "seeds:") {
			line = sc.Text()
			break
		}
	}

	seedList := strings.TrimSpace(strings.TrimPrefix(line, "seeds:"))
	seedStrs := strings.Fields(seedList)

	for i := 0; i < len(seedStrs)-1; i += 2 {
		start := parseInt(seedStrs[i])
		length := parseInt(seedStrs[i+1])

		for j := start; j < start+length; j++ {
			fmt.Printf("adding seed %v\n", j)
			allSeeds = append(allSeeds, Seeder{Seed: j})
		}
	}

	return allSeeds
}

func getLowestLocation() int64 {
	lowest := math.MaxInt64

	for _, seed := range seeds {
		if seed.Location < int64(lowest) {
			lowest = int(seed.Location)
		}
	}

	return int64(lowest)
}

func (s *Seeder) setProperty(name string, value int64) *Seeder {
	reflect.ValueOf(s).Elem().FieldByName(name).SetInt(value)
	return s
}

func parseEntry(s string) Entry {
	parts := strings.Fields(s)
	return Entry{
		destStart:   parseInt(parts[0]),
		sourceStart: parseInt(parts[1]),
		length:      parseInt(parts[2]),
	}
}

func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func printSeeds(seeds []Seeder) {
	for _, seed := range seeds {
		fmt.Printf("seed = %+v\n", seed)
	}
}
