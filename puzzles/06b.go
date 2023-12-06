package puzzles

import (
	"advent/utils"
	"regexp"
	"strconv"
)

func parseSingleRaceData(data []string) Race {
	r := regexp.MustCompile(`([0-9]+)`)
	timeVals := r.FindAllString(data[0], -1)
	recordVals := r.FindAllString(data[1], -1)	

	timeString := ""
	recordString := ""
	for i := 0; i < len(timeVals); i++ {
		timeString += timeVals[i]
		recordString += recordVals[i]
	}
	time, _ := strconv.Atoi(timeString)
	record, _ := strconv.Atoi(recordString)

	return Race{time, record}
}

func Puzzle06b() string {
	data := utils.FileReader("data/06.txt")
	race := parseSingleRaceData(data)

	count := race.getRecordBreakingChargeTimes()

	return strconv.Itoa(count)
}