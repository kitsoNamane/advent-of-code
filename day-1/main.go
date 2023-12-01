package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var codes int

	file, err := os.Open("puzzle.txt")
	if err != nil {
		os.Exit(-1)
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		codes += findFirstAndLastDigit(scan.Text())
	}

	fmt.Println(codes)
}

func findFirstAndLastDigit(s string) int {
	var code string
	first := 0
	last := len(s) - 1

	for first <= last {
		if num, err := strconv.Atoi(string(s[first])); err == nil {
			code += fmt.Sprint(num)
			break
		}
		first++
	}

	for first <= last {
		if num, err := strconv.Atoi(string(s[last])); err == nil {
			code += fmt.Sprint(num)
			break
		}
		last--
	}
	num, _ := strconv.Atoi(code)
	return num
}
