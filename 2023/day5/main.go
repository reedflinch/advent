package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/5

var fileName = "input.txt"

type Entry struct {
	sourceStart, destStart, length int64
}

type Mapping []Entry

type Almanac struct {
	seedToSoil            Mapping
	soilToFertilizer      Mapping
	fertilizerToWater     Mapping
	waterToLight          Mapping
	lightToTemperature    Mapping
	temperatureToHumidity Mapping
	humidityToLocation    Mapping
}

func runPart(part int) {
	lowest := math.MaxInt64

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content := string(bytes)
	content = strings.Trim(content, "\n")
	groups := strings.Split(content, "\n\n")

	seeds := []int64{}
	seeds = append(seeds, getSeeds(groups[0], part)...)

	a := Almanac{
		groupToMapping(groups[1]),
		groupToMapping(groups[2]),
		groupToMapping(groups[3]),
		groupToMapping(groups[4]),
		groupToMapping(groups[5]),
		groupToMapping(groups[6]),
		groupToMapping(groups[7]),
	}

	for _, seed := range seeds {
		location := a.traverse(seed)
		if location < int64(lowest) {
			lowest = int(location)
		}
	}

	fmt.Printf("lowest location (part %v) = %v\n", part, lowest)
}

func (m Mapping) get(val int64) int64 {
	for _, entry := range m {
		if val >= entry.sourceStart && val <= entry.sourceStart+entry.length-1 {
			newVal := val + (entry.destStart - entry.sourceStart)
			return newVal
		}
	}

	return val
}

func (a *Almanac) traverse(seed int64) (location int64) {
	soil := a.seedToSoil.get(seed)
	fertilizer := a.soilToFertilizer.get(soil)
	water := a.fertilizerToWater.get(fertilizer)
	light := a.waterToLight.get(water)
	temperature := a.lightToTemperature.get(light)
	humidity := a.temperatureToHumidity.get(temperature)
	location = a.humidityToLocation.get(humidity)

	return location
}

func groupToMapping(group string) Mapping {
	lines := strings.Split(group, "\n")
	mappingLines := lines[1:]

	mappings := make(Mapping, 0, len(mappingLines))
	for _, line := range mappingLines {
		entry := parseEntry(line)
		mappings = append(mappings, entry)
	}

	return mappings
}

func parseEntry(s string) Entry {
	parts := strings.Fields(s)
	return Entry{
		destStart:   parseInt(parts[0]),
		sourceStart: parseInt(parts[1]),
		length:      parseInt(parts[2]),
	}
}

func getSeeds(line string, part int) []int64 {
	if part == 1 {
		fmt.Printf("getting seeds part 1...\n")
		return getSeedsPart1(line)
	}
	fmt.Printf("getting seeds part 2...\n")
	return getSeedsPart2(line)
}

func getSeedsPart1(line string) []int64 {
	allSeeds := []int64{}

	seedList := strings.TrimSpace(strings.TrimPrefix(line, "seeds:"))
	seedStrs := strings.Fields(seedList)

	for _, seedStr := range seedStrs {
		id := parseInt(seedStr)
		allSeeds = append(allSeeds, id)
	}

	return allSeeds
}

func getSeedsPart2(line string) []int64 {
	allSeeds := []int64{}

	seedList := strings.TrimSpace(strings.TrimPrefix(line, "seeds:"))
	seedStrs := strings.Fields(seedList)

	for i := 0; i < len(seedStrs)-1; i += 2 {
		start := parseInt(seedStrs[i])
		length := parseInt(seedStrs[i+1])

		for j := start; j < start+length; j++ {
			allSeeds = append(allSeeds, j)
		}
	}

	return allSeeds
}

func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func main() {
	runPart(1)
	runPart(2)
}
