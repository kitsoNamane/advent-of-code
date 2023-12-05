package day_one

import (
	"fmt"
	"strconv"
	"strings"
)

func FindFirstAndLastDigit(s string) int {
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

func ReplaceWithDigit(s string) string {
	digits := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	firstDigit := ""
	firstDigitIndex := len(s) - 1
	lastDigit := ""
	lastDigitIndex := 0

	for key := range digits {
		first := strings.Index(s, key)
		if first <= firstDigitIndex && first != -1 {
			firstDigitIndex = first
			firstDigit = key
		}
	}
	if firstDigit != "" {
		replace := digits[firstDigit]
		s = strings.Replace(s, firstDigit, replace+string(firstDigit[len(firstDigit)-1]), 1)
	}
	for key := range digits {
		last := strings.LastIndex(s, key)
		if last >= lastDigitIndex && last != -1 {
			lastDigitIndex = last
			lastDigit = key
		}
	}
	if lastDigit != "" {
		replace := digits[lastDigit]
		s = strings.Replace(s, lastDigit, replace, 1)
	}
	return s
}
