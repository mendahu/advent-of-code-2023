package puzzles

import (
	"advent/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Race struct {
	Time int
	Record int
}


func (r *Race) getMinChargeTime() int {
	min := 0

	for i := 0; i <= r.Record; i++ {
		if (i % 1000000) == 0 {
			fmt.Println(i, r.Record)
		}
		if ((r.Time - i) * i) > r.Record {
			min = i
			break
		}
	}

	return min
}

func (r *Race) getRecordBreakingChargeTimes() int {
	min := r.getMinChargeTime()

	return (r.Time + 1) - (min * 2)
}

func parseRaceData(data []string) []Race {
	r := regexp.MustCompile(`([0-9]+)`)
	times := r.FindAllString(data[0], -1)
	records := r.FindAllString(data[1], -1)	

	races := make([]Race, 0)

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		record, _ := strconv.Atoi(records[i])

		race := Race{time, record}
		races = append(races, race)
	}

	return races
}

func Puzzle06a() string {
	data := utils.FileReader("data/06.txt")
	raceData := parseRaceData(data)

	count := 1

	for _, race := range raceData {
		times := race.getRecordBreakingChargeTimes()
		fmt.Println(times)
		count = count * times
	}

	return strconv.Itoa(count)
}