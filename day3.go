package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	input := ReadFile("day3.txt")
	var claims = []claim{}

	for _, line := range input {
		claims = append(claims, parseClaim(line))
	}
	fmt.Println(claims)
//	fabric := make(map[string]string)
}

func parseClaim(line string) claim{ReadFile("day3.txt")
	s := strings.Split(line, " ")
	id := parseToInt(string(s[0][1:]))
	
	cords := strings.Split(s[2], ",")
	x := parseToInt(cords[0])
	y := parseToInt(strings.TrimSuffix(cords[1], ":"))
	
	w := parseToInt(strings.Split(s[3], "x")[0])
	h := parseToInt(strings.Split(s[3], "x")[1])

	return claim{id, x ,y ,w, h}
}

func parseToInt(value string) int {
	parsedValue, _ := strconv.Atoi(value)
	return parsedValue
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

type claim struct{
	id int
	x int
	y int
	w int
	h int
}
