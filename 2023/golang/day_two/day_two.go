package day_two

import (
	"strconv"
	"strings"
)

func CubeConundrum(game string) (int, bool) {
	gameCubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	line := strings.Split(game, ":")

	_, gameId := getNumberAndString(line[0])

	gameInfo := line[1]

	rounds := strings.Split(gameInfo, ";")

	for _, round := range rounds {
		roundData := strings.Split(round, ",")
		for _, roundDataItem := range roundData {
			cube, num := getNumberAndString(strings.Trim(roundDataItem, " "))
			if gameCubes[cube] < num {
				return 0, false
			}
		}
	}

	return gameId, true
}

func getNumberAndString(s string) (string, int) {
	items := strings.Split(s, " ")

	if num, err := strconv.Atoi(items[0]); err == nil {
		return items[1], num
	}

	num, _ := strconv.Atoi(items[1])
	return items[0], num
}
