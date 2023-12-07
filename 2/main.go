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

type RGB struct {
	red   int8
	green int8
	blue  int8
}

func main() {
	file, err := os.Open("data.txt")
	gameNumberTotal := 0
	powerTotal := 0

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parseLine(text, &gameNumberTotal, &powerTotal)
	}

	file.Close()
	fmt.Println("total", gameNumberTotal)
	fmt.Println("power", powerTotal)
}

func parseLine(text string, gameNumberTotal *int, powerTotal *int) {
	gameAndValue := strings.Split(text, ": ")
	gameNumber := getGameNumber(gameAndValue[0])
	rounds := strings.Split(gameAndValue[1], ";")

	linePower := getLinePowerFromRounds(rounds)
	*powerTotal = *powerTotal + linePower

	roundAreValid := areRoundsValid(rounds)
	if roundAreValid {
		*gameNumberTotal = *gameNumberTotal + gameNumber
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
	value := getCubeValue(color, round)

	if value > maxes[color] {
		roundIsValid = false
	}

	return roundIsValid
}

func getCubeValue(color string, round string) int8 {
	exp := "(?m)(?P<" + color + `>\d+) ` + color
	r := regexp.MustCompile(exp)
	matches := r.FindAllStringSubmatch(round, -1)
	var value int8 = 0

	if len(matches) != 0 {
		if len(matches[0]) != 0 {
			num, _ := strconv.Atoi(matches[0][1])
			value = int8(num)
		}
	}

	return value
}

func getLinePowerFromRounds(rounds []string) int {
	rgb := RGB{
		red:   0,
		green: 0,
		blue:  0,
	}

	for _, round := range rounds {
		red := getCubeValue("red", round)
		if red > rgb.red {
			rgb.red = red
		}
		green := getCubeValue("green", round)
		if green > rgb.green {
			rgb.green = green
		}
		blue := getCubeValue("blue", round)
		if blue > rgb.blue {
			rgb.blue = blue
		}
	}

	total := int(rgb.blue) * int(rgb.red) * int(rgb.green)
	return total
}
