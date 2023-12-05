package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

type SeedRange struct {
	Seed int
	Offset int
}

func (sr *SeedRange) LowerBound() int {
	return sr.Seed
}

func (sr *SeedRange) UpperBound() int {
	return sr.Seed + sr.Offset - 1
}


func getSeedRangesFromData(data string) (seedRanges []SeedRange) {
	// get seeds
	parts := strings.Split(data, ": ")
	rangeList := parts[1]
	seedsStrings := strings.Split(rangeList, " ")
	
	for i := 0; i < len(seedsStrings); i += 2 {
		seed, _ := strconv.Atoi(seedsStrings[i])
		offset, _ := strconv.Atoi(seedsStrings[i + 1])
		seedrange := SeedRange{seed, offset}
		seedRanges = append(seedRanges, seedrange)
	}
	return seedRanges
}

func (m *Map) mapSeeds(seedRanges []SeedRange) (newRanges []SeedRange) {

	fmt.Println("SeedRange length", len(seedRanges))
	for i:= 0; i < len(seedRanges); i++ {
		fmt.Println("\nSeedRange:", seedRanges[i])
		hasPassedThrough := false

		seedRange := seedRanges[i]

		for _, mapper := range m.Mappers {
			fmt.Println("Mapper:", mapper)
			// mapper does not in anyway intersect seedrange
			if mapper.UpperBound() < seedRange.LowerBound() || mapper.LowerBound() > seedRange.UpperBound() {
				fmt.Println(" Mapper does not intersect with seedRange")
				continue;
			}

			// mapper intersects lowerBound, upperBound out of range
			// seedRange is split in to two, lower version processed through mapper, upper version is recycled for further checks
			if (
					mapper.LowerBound() <= seedRange.LowerBound() && 
					mapper.UpperBound() >= seedRange.LowerBound() && 
					mapper.UpperBound() < seedRange.UpperBound()) {

				fmt.Println(" Mapper intersects lowerBound, upperBound out of range")
				lowerVersion := SeedRange{seedRange.LowerBound() + mapper.TransformValue(), mapper.UpperBound() - seedRange.LowerBound() + 1}
				newRanges = append(newRanges, lowerVersion)

				upperVersion := SeedRange{mapper.UpperBound() + 1, seedRange.Offset - lowerVersion.Offset}
				seedRanges = append(seedRanges, upperVersion)

				hasPassedThrough = true
				break;
			}

			// mapper intersects lowerBound, upperBound in range
			// entire seedRange is processed through mapper
			if mapper.LowerBound() <= seedRange.LowerBound() && mapper.UpperBound() >= seedRange.UpperBound() {
				fmt.Println(" Mapper completely intersects seedRange")
				newRange := SeedRange{seedRange.Seed + mapper.TransformValue(), seedRange.Offset}
				newRanges = append(newRanges, newRange)
				hasPassedThrough = true
				break;
			}

			// mapper intersects upperBound, lowerBound out of range
			// seedRange is split in to two, upper version processed through mapper, lower version is recycled for further checks
			if mapper.LowerBound() > seedRange.LowerBound() && mapper.LowerBound() <= seedRange.UpperBound() && mapper.UpperBound() > seedRange.UpperBound() {
				fmt.Println(" Mapper intersects upperBound, lowerBound out of range")
				upperVersion := SeedRange{mapper.LowerBound() + mapper.TransformValue(), seedRange.UpperBound() - mapper.LowerBound() + 1}
				newRanges = append(newRanges, upperVersion)

				lowerVersion := SeedRange{seedRange.LowerBound(), seedRange.Offset - upperVersion.Offset}
				seedRanges = append(seedRanges, lowerVersion)
				hasPassedThrough = true
				break;
			}
		}

		if !hasPassedThrough {
			fmt.Println(" Passing entire seedRange through")
			newRanges = append(newRanges, seedRange)
		}
	}

	return newRanges
}

func getLowestSeedValue(seedRanges []SeedRange) int {
	lowest := seedRanges[0].Seed

	for _, seedRange := range seedRanges {
		if seedRange.Seed < lowest {
			lowest = seedRange.Seed
		}
	}
	return lowest
}

func Puzzle05b() string {
	data := utils.FileReader("data/05.txt")

	maps := getMapValuesFromData(data)
	seedRanges := getSeedRangesFromData(data[0])
	fmt.Println("*** Base Seed Ranges:", seedRanges)

	soil := maps["seed-to-soil"].mapSeeds(seedRanges)
	fmt.Println("\n*** Soil Seed Ranges:", soil)

	fertilizer := maps["soil-to-fertilizer"].mapSeeds(soil)
	fmt.Println("\n*** Fertilizer Seed Ranges:", fertilizer)

	water := maps["fertilizer-to-water"].mapSeeds(fertilizer)
	fmt.Println("\n*** Water Seed Ranges:", water)

	light := maps["water-to-light"].mapSeeds(water)
	fmt.Println("\n*** Light Seed Ranges:", light)

	temperature := maps["light-to-temperature"].mapSeeds(light)
	fmt.Println("\n*** Temperature Seed Ranges:", temperature)

	humidity := maps["temperature-to-humidity"].mapSeeds(temperature)
	fmt.Println("\n*** Humidity Seed Ranges:", humidity)

	location := maps["humidity-to-location"].mapSeeds(humidity)
	fmt.Println("\n*** Location Seed Ranges:", location)

	lowest := getLowestSeedValue(location)
	return strconv.Itoa(lowest)
}