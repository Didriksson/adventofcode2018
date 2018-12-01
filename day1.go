package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	currentFreq := 0
	seenFreqs := make(map[int]bool)
	seenFreqs[0] = true
	for {
		currentFreq = iterateList(currentFreq, seenFreqs)
	}
}

func iterateList(currentFreq int, seenFreqs map[int]bool) int {
	file, err := os.Open("day1.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		currentFreq += value
		_, exist := seenFreqs[currentFreq]
		if exist {
			fmt.Println("Hittat upprepad frekvens: ", currentFreq)
			os.Exit(0)
		}
		seenFreqs[currentFreq] = true
	}
	return currentFreq
}
