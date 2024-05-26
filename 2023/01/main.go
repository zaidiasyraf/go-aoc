package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbersString := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	total := 0
	total2 := 0
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		var t [2]int
		line := scanner.Text()

		lowestIndex := -1
		highestIndex := -1
		for _, n := range numbersString {
			i := strings.Index(line, n)
			if i != -1 && (i < lowestIndex || lowestIndex == -1) {
				lowestIndex = i
			}
			i = strings.LastIndex(line, n)
			if i != -1 && (i > highestIndex || highestIndex == -1) {
				highestIndex = i
			}
		}
		if lowestIndex == -1 {
			lowestIndex = len(line)
		}
		if highestIndex == -1 {
			highestIndex = 0
		}

		for _, c := range line {
			if unicode.IsDigit(c) {
				i := int(c - '0')
				if t[0] == 0 {
					t[0] = int(i)
				} else {
					t[1] = int(i)
				}
			}
		}
		if t[1] == 0 {
			t[1] = t[0]
		}

		total = total + t[0]*10 + t[1]

		i := strings.Index(line, strconv.Itoa(t[0]))
		if i != -1 && i < lowestIndex {
			lowestIndex = i
		}
		i = strings.LastIndex(line, strconv.Itoa(t[1]))
		if i != -1 && i > highestIndex {
			highestIndex = i
		}

		low := line[lowestIndex:]
		high := line[highestIndex:]
		if !unicode.IsDigit(rune(low[0])) {
			for _, num := range numbersString {
				if strings.HasPrefix(low, num) {
					t[0] = convertNum(num)
				}
			}
		} else {
			t[0] = int(low[0] - '0')
		}

		if !unicode.IsDigit(rune(high[0])) {
			for _, num := range numbersString {
				if strings.HasPrefix(high, num) {
					t[1] = convertNum(num)
				}
			}
		} else {
			t[1] = int(high[0] - '0')
		}

		fmt.Printf("For linenumber %d, low %d, high %d\n", lineNumber, t[0], t[1])
		total2 = total2 + t[0]*10 + t[1]
	}
	fmt.Println("Answer 1 : ", total)
	fmt.Println("Answer 2 : ", total2)
}

func convertNumStr(s string) string {
	switch s {
	case "1":
		return "one"
	case "2":
		return "two"
	case "3":
		return "three"
	case "4":
		return "four"
	case "5":
		return "five"
	case "6":
		return "six"
	case "7":
		return "seven"
	case "8":
		return "eight"
	case "9":
		return "nine"
	default:
		return "this should not reach"
	}
}

func convertNum(s string) int {
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
