package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile("day3.txt")
	var claims = []claim{}
	fabric := make(map[string][]int)

	for _, line := range input {
		claims = append(claims, parseClaim(line))
	}

	for _, c := range claims {
		fabric = applyClaimToFabric(c, fabric)
	}

	fmt.Println("Answer to part 1:", lookForOverlapping(fabric))

	noOfUniqueCellsPerID := removeDuplicates(fabric)
	fmt.Println("Answer to part 2: Claim with id - ", lookForMatchingClaim(claims, noOfUniqueCellsPerID).id)

}

func lookForMatchingClaim(claimsList []claim, uniquePerIds map[int]int) claim {
	for id, numberOf := range uniquePerIds {
		for _, claim := range claimsList {
			if claim.id == id {
				if claim.w*claim.h == numberOf {
					return claim
				}
				break
			}
		}
	}
	fmt.Println("Could not find the mathing claim")
	return claim{}
}

func removeDuplicates(f map[string][]int) map[int]int {
	var uniqueCells = make(map[int]int)
	for _, v := range f {
		if len(v) == 1 {
			uniqueCells[v[0]] = uniqueCells[v[0]] + 1
		}
	}
	return uniqueCells
}

func lookForOverlapping(f map[string][]int) int {
	var overlapping int
	for _, id := range f {
		if len(id) > 1 {
			overlapping++
		}
	}
	return overlapping
}

func applyClaimToFabric(c claim, f map[string][]int) map[string][]int {
	for y := c.y; y < c.y+c.h; y++ {
		for x := c.x; x < c.x+c.w; x++ {
			cord := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, ",")
			f[cord] = append(f[cord], c.id)
		}
	}
	return f
}

func parseClaim(line string) claim {
	ReadFile("day3.txt")
	s := strings.Split(line, " ")
	id := parseToInt(string(s[0][1:]))

	cords := strings.Split(s[2], ",")
	x := parseToInt(cords[0])
	y := parseToInt(strings.TrimSuffix(cords[1], ":"))

	w := parseToInt(strings.Split(s[3], "x")[0])
	h := parseToInt(strings.Split(s[3], "x")[1])

	return claim{id, x, y, w, h}
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

type claim struct {
	id int
	x  int
	y  int
	w  int
	h  int
}
