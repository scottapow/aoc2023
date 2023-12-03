package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maxes map[string]int8 = map[string]int8{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("data.txt")
	var total int = 0

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parseAndAddToTotal(text, &total)
	}

	file.Close()
	fmt.Println(total)
}

func parseAndAddToTotal(text string, total *int) {
	gameAndValue := strings.Split(text, ": ")
	gameNumber := getGameNumber(gameAndValue[0])
	rounds := strings.Split(gameAndValue[1], ";")
	roundAreValid := areRoundsValid(rounds)

	// Game 1: 1 green, 4 blue; 1 blue, 2 green, 1 red; 1 red, 1 green, 2 blue; 1 green, 1 red; 1 green; 1 green, 1 blue, 1 red
	if roundAreValid {
		*total = *total + gameNumber
	}
}

func getGameNumber(game string) int {
	r, err := regexp.Compile(`Game (\d+)`)

	if err != nil {
		panic(err)
	}

	matches := r.FindStringSubmatch(game)
	gameNum, err := strconv.Atoi(matches[1])

	if err != nil {
		panic(err)
	}

	return gameNum
}

func areRoundsValid(rounds []string) bool {
	isValid := true
	for i := 0; i < len(rounds); i++ {
		isRedValid := isRoundColorValid("red", rounds[i])
		isGreenValid := isRoundColorValid("green", rounds[i])
		isBludValid := isRoundColorValid("blue", rounds[i])
		if !isRedValid || !isGreenValid || !isBludValid {
			isValid = false
		}
	}

	return isValid
}

func isRoundColorValid(color string, round string) bool {
	roundIsValid := true

	exp := "(?m)(?P<" + color + `>\d+) ` + color
	r := regexp.MustCompile(exp)
	matches := r.FindAllStringSubmatch(round, -1)

	if len(matches) != 0 {
		if len(matches[0]) != 0 {
			num, _ := strconv.Atoi(matches[0][1])
			if int8(num) > maxes[color] {
				roundIsValid = false
			}
		}
	}

	return roundIsValid
}
