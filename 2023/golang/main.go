package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kitsoNamane/advent-of-code/day_two"
)

func main() {
	var count int
	file, err := os.Open("day_two/puzzle.txt")
	if err != nil {
		os.Exit(-1)
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		if gameId, isPossible := day_two.CubeConundrum(scan.Text()); isPossible {
			count += gameId
		}
	}

	fmt.Println(count)
}
