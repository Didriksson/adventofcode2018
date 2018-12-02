package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadFile("day2.txt")
	part1(input)
	part2(input)
}

func part1(input []string) {
	var twos, threes int
	for _, line := range input {
		two, three := checkNoOfCharactersPerRow(line)
		if two > 0 {
			twos++
		}
		if three > 0 {
			threes++
		}
	}

	fmt.Println("Answer part 1: ", twos, "*", threes, twos*threes)
}

func part2(input []string) {
	for i, cLine := range input {
		for _, compareLine := range input[i+1:] {
			common, nonMatching := getCommonString(cLine, compareLine)
			if nonMatching == 1 {
				fmt.Println("Answer part 2: ", common)
				os.Exit(0)
			}
		}
	}
}

func getCommonString(a, b string) (string, int) {
	var common string
	var noOfNonCommon int
	for i, c := range a {
		if c == rune(b[i]) {
			common += string(c)
		} else {
			noOfNonCommon++
		}
	}
	return common, noOfNonCommon
}

func checkNoOfCharactersPerRow(row string) (int, int) {
	var twos, threes int
	wordcount := make(map[string]int)
	for _, c := range row {
		wordcount[string(c)] = wordcount[string(c)] + 1
		if wordcount[string(c)] == 2 {
			twos++
		}
		if wordcount[string(c)] == 3 {
			twos--
			threes++
		}
		if wordcount[string(c)] == 4 {
			threes--
		}
	}
	return twos, threes
}

func ReadFile(path string) []string {
	file, err := os.Open(path)
	var lines []string
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
