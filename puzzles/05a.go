package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

type Mapper struct {
	Source int
	Dest int
	Range int
}

func (m *Mapper) LowerBound() int {
	return m.Source
}

func (m *Mapper) UpperBound() int {
	return m.Source + m.Range - 1
}

func (m *Mapper) TransformValue() int {
	return m.Dest - m.Source
}

func (m *Mapper) Check(value int) (result int, inRange bool) {
	if value >= m.Source && value < m.Source + m.Range {
		depth := value - m.Source
		return m.Dest + depth, true
	} else {
		return value, false
	}
}

type Map struct {
	Mappers []*Mapper
}

func (m *Map) addMapper(mapper *Mapper) []*Mapper {
	m.Mappers = append(m.Mappers, mapper)
	return m.Mappers
}

func (m *Map) mapValue(value int)	int {
	for _, mapper := range m.Mappers {
		if result, inRange := mapper.Check(value); inRange {
			return result
		}
	}
	return value
}


func getSingleSeedsFromData(data string) (seeds []int) {
	// get seeds
	parts := strings.Split(data, ": ")
	seedsStrings := strings.Split(parts[1], " ")

	for _, seedString := range seedsStrings {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	return seeds
}

func getMapValuesFromData(data []string) (mappers map[string]*Map) {
	maps := map[string]*Map{}

	// get maps
	mapperTitle := ""
	for i := 2; i < len(data); i++ {
		if data[i] == "" {
			continue
		}
		
		if strings.Contains(data[i], "map:") {
			mapperTitle = strings.TrimSuffix(data[i], " map:")
			newMap := Map{
				Mappers: make([]*Mapper, 0),
			}

			maps[mapperTitle] = &newMap
			continue
		} else {
			var SourceString, DestString, RangeString string
			fmt.Sscanf(data[i], "%s %s %s", &DestString, &SourceString, &RangeString)
			Source, _ := strconv.Atoi(SourceString)
			Dest, _ := strconv.Atoi(DestString)
			Range, _ := strconv.Atoi(RangeString)

			mapper := Mapper{Source, Dest, Range}
			maps[mapperTitle].addMapper(&mapper)
		}
	}

	return maps
}

func Puzzle05a() string {
	data := utils.FileReader("data/05.txt")

	maps := getMapValuesFromData(data)
	seeds := getSingleSeedsFromData(data[0])
	var lowest *int

	for _, seed := range seeds {
			fmt.Println(" ")
			fmt.Println("*** NEW SEED", seed)
			soil := maps["seed-to-soil"].mapValue(seed)
			fmt.Println("soil:", soil)
			fertilizer := maps["soil-to-fertilizer"].mapValue(soil)
			fmt.Println("fertilizer:", fertilizer)
			water := maps["fertilizer-to-water"].mapValue(fertilizer)
			fmt.Println("water:", water)
			light := maps["water-to-light"].mapValue(water)
			fmt.Println("light:", light)
			temperature := maps["light-to-temperature"].mapValue(light)
			fmt.Println("temperature:", temperature)
			humidity := maps["temperature-to-humidity"].mapValue(temperature)
			fmt.Println("humidity:", humidity)
			location := maps["humidity-to-location"].mapValue(humidity)
			fmt.Println("location:", location)
			if lowest == nil || location < *lowest {
				lowest = &location
			}
	}

	return strconv.Itoa(*lowest)
}