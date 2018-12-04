package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	input := ReadFile("day4.txt")
	sort.Strings(input)
	const layout = "2006-01-02 15:04"

	start, _ := time.Parse(layout, "1518-11-01 23:58")
	end, _ := time.Parse(layout, "1518-11-02 00:40")
	fmt.Println(end.Sub(start))
	/* 	for _, l := range input {
		const layout = "2006-01-02 15:04"
		fmt.Println(time.Parse(layout, "1518-11-05 00:03"))
	} */
}

type guard struct {
	id int
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
