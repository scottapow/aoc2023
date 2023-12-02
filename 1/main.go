package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	var total int = 0

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		r := regexp.MustCompile("[0-9]")
		matches := r.FindAllString(text, -1)
		if matches != nil {
			total += addMatches(matches[0], matches[len(matches)-1])
		}
	}

	file.Close()
	fmt.Println(total)
}

func addMatches(first string, last string) int {
	combined, err := strconv.Atoi(first + last)

	if err != nil {
		panic(err)
	}

	return combined
}
