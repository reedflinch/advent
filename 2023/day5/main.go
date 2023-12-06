package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// https://adventofcode.com/2023/day/4

type Seed struct {
	id, soil, fert, water, light, temp, Humidity, Location int64
}

type Entry struct {
	sourceStart, destStart, length int64
}

var fileName = "input.txt"
var seeds = []Seed{}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	seeds = append(seeds, getSeeds(s)...)
	fmt.Printf("\n\nseeds before soil\n")
	printSeeds(seeds)

	seedsToSoils(s)
	// fmt.Printf("\n\nseeds with soil\n")
	// printSeeds(seeds)

	soilsToFert(s)
	// fmt.Printf("\n\nseeds with fert\n")
	// printSeeds(seeds)

	fertToWater(s)
	// fmt.Printf("\n\nseeds with water\n")
	// printSeeds(seeds)

	waterToLight(s)
	// fmt.Printf("\n\nseeds with light\n")
	// printSeeds(seeds)

	lightToTemp(s)
	// fmt.Printf("\n\nseeds with light\n")
	// printSeeds(seeds)

	tempToHumidity(s)
	// fmt.Printf("\n\nseeds with light\n")
	// printSeeds(seeds)

	humidityToLocation(s)
	// fmt.Printf("\n\nseeds with light\n")
	// printSeeds(seeds)

	fmt.Printf("lowest Location (part 1) = %v\n", getLowestLocation())
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

func humidityToLocation(sc *bufio.Scanner) {
	headerName := "humidity-to-location map:"

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
				// populate map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false

		for _, entry := range mapping {
			if seed.Humidity >= entry.sourceStart && seed.Humidity <= entry.sourceStart+entry.length-1 {
				seeds[i].Location = seed.Humidity + (entry.destStart - entry.sourceStart)
				found = true
				break
			}
		}
		if !found {
			seeds[i].Location = seed.Humidity
		}
	}
}

func tempToHumidity(sc *bufio.Scanner) {
	headerName := "temperature-to-humidity map:"

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
				// populate map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false

		for _, entry := range mapping {
			if seed.temp >= entry.sourceStart && seed.temp <= entry.sourceStart+entry.length-1 {
				seeds[i].Humidity = seed.temp + (entry.destStart - entry.sourceStart)
				found = true
				break
			}
		}
		if !found {
			seeds[i].Humidity = seed.temp
		}
	}
}

func lightToTemp(sc *bufio.Scanner) {
	headerName := "light-to-temperature map:"

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
				// populate map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false

		for _, entry := range mapping {
			if seed.light >= entry.sourceStart && seed.light <= entry.sourceStart+entry.length-1 {
				seeds[i].temp = seed.light + (entry.destStart - entry.sourceStart)
				found = true
				break
			}
		}
		if !found {
			seeds[i].temp = seed.light
		}
	}
}

func waterToLight(sc *bufio.Scanner) {
	headerName := "water-to-light map:"

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
				// populate map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false

		for _, entry := range mapping {
			if seed.water >= entry.sourceStart && seed.water <= entry.sourceStart+entry.length-1 {
				seeds[i].light = seed.water + (entry.destStart - entry.sourceStart)
				found = true
				break
			}
		}
		if !found {
			seeds[i].light = seed.water
		}
	}
}

func fertToWater(sc *bufio.Scanner) {
	headerName := "fertilizer-to-water map:"

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
				// populate map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false

		for _, entry := range mapping {
			if seed.fert >= entry.sourceStart && seed.fert <= entry.sourceStart+entry.length-1 {
				seeds[i].water = seed.fert + (entry.destStart - entry.sourceStart)
				found = true
				break
			}
		}
		if !found {
			seeds[i].water = seed.fert
		}
	}
}

func soilsToFert(sc *bufio.Scanner) {
	headerName := "soil-to-fertilizer map:"

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
				// populate map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false

		for _, entry := range mapping {
			// lowerBound := entry.sourceStart
			// upperBound := entry.sourceStart + entry.length - 1
			if seed.soil >= entry.sourceStart && seed.soil <= entry.sourceStart+entry.length-1 {
				seeds[i].fert = seed.soil + (entry.destStart - entry.sourceStart)
				found = true
				break
			}
		}
		if !found {
			seeds[i].fert = seed.soil
		}
	}
}

func seedsToSoils(sc *bufio.Scanner) {
	headerName := "seed-to-soil map:"

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
				// populate map
				mapping = append(mapping, parseEntry(sc.Text()))
			}
		}
	}

	// process all seeds
	for i, seed := range seeds {
		found := false

		for _, entry := range mapping {
			if seed.id >= entry.sourceStart && seed.id <= entry.sourceStart+entry.length-1 {
				seeds[i].soil = seed.id + (entry.destStart - entry.sourceStart)
				found = true
				break
			}
		}
		if !found {
			seeds[i].soil = seed.id
		}
	}
}

func parseEntry(s string) Entry {
	parts := strings.Fields(s)
	return Entry{
		destStart:   parseInt(parts[0]),
		sourceStart: parseInt(parts[1]),
		length:      parseInt(parts[2]),
	}
}

func getSeeds(sc *bufio.Scanner) []Seed {
	line := ""
	allSeeds := []Seed{}

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
		allSeeds = append(allSeeds, Seed{id: id})
	}

	return allSeeds
}

func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func printSeeds(seeds []Seed) {
	for _, seed := range seeds {
		fmt.Printf("seed = %+v\n", seed)
	}
}

/*
Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, Humidity 78, Location 82.
Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, Humidity 43, Location 43.
Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, Humidity 82, Location 86.
Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, Humidity 35, Location 35.
*/
