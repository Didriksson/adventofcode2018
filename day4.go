package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := ReadFile("day4.txt")
	sort.Strings(input)
	guards := make(map[string][]entry)

	var currentGuard string
	for _, l := range input {
		var idRe = regexp.MustCompile(`(?m)#(\d+)`)
		var id = strings.TrimPrefix(string(idRe.FindString(l)), "#")
		if id != "" {
			currentGuard = id
		} else {
			guards[currentGuard] = append(guards[currentGuard], parseEntry(l))
		}
	}
	avidSleeper := getIDOfGuardSleepingTheMost(guards)
	avidSleepersMostAvidSleepingMinute, _ := getMinuteGuardIsMostFrequentlyAsleep(guards[avidSleeper])
	id, _ := strconv.Atoi(avidSleeper)

	fmt.Println("Part 1:", id*avidSleepersMostAvidSleepingMinute)

	id, min := getGuardIDAndMostAvidSleepingMinute(guards)
	fmt.Println("Part 2:", id*min)
}

func getGuardIDAndMostAvidSleepingMinute(guards map[string][]entry) (int, int) {
	var highestMin int
	var highestNoOfTimes int
	var highestId string

	for id, guard := range guards {
		min, times := getMinuteGuardIsMostFrequentlyAsleep(guard)
		if times > highestNoOfTimes {
			highestMin = min
			highestId = id
			highestNoOfTimes = times
		}
	}
	id, _ := strconv.Atoi(highestId)
	return id, highestMin
}

func getIDOfGuardSleepingTheMost(guards map[string][]entry) string {
	var highestMin float64
	var highestID string
	for id, guard := range guards {
		currentMin := timeAsleepForGuard(guard)
		if currentMin > highestMin {
			highestID = id
			highestMin = currentMin
		}
	}
	return highestID
}

func getMinuteGuardIsMostFrequentlyAsleep(e []entry) (int, int) {
	asleep := make(map[int]int)
	sort.Slice(e, func(i, j int) bool { return e[i].time.Before(e[j].time) })

	if len(e) == 1 {
		return -1, -1
	}

	for i := 0; i < len(e)-1; i = i + 2 {
		durationofSleep := e[i+1].time.Sub(e[i].time).Minutes() - 1
		for deltamin := 0; deltamin <= int(durationofSleep); deltamin++ {
			addMinutes, _ := time.ParseDuration(strconv.Itoa(deltamin) + "m")
			min := e[i].time.Add(addMinutes).Minute()
			asleep[min] = asleep[min] + 1
		}
	}
	var highest int
	for k, v := range asleep {
		if v > asleep[highest] {
			highest = k
		}
	}
	return highest, asleep[highest]
}

func timeAsleepForGuard(e []entry) float64 {

	sort.Slice(e, func(i, j int) bool { return e[i].time.Before(e[j].time) })

	if len(e) == 1 {
		return 0
	}

	var totalTimeAsleep float64

	for i := 0; i < len(e)-1; i = i + 2 {
		totalTimeAsleep += e[i+1].time.Sub(e[i].time).Minutes()
	}
	return totalTimeAsleep
}

func parseEntry(line string) entry {
	var re = regexp.MustCompile(`(?m)\[(.+)].+(begins shift|falls asleep|wakes up)`)
	const layout = "2006-01-02 15:04"
	regexMatches := re.FindStringSubmatch(line)
	time := regexMatches[1]
	action := regexMatches[2]
	return entry{parseTime(time), action}
}

func parseTime(timestring string) time.Time {
	const layout = "2006-01-02 15:04"
	time, _ := time.Parse(layout, timestring)
	return time
}

type entry struct {
	time   time.Time
	action string
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
