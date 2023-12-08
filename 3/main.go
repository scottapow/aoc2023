package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Symbol struct {
	value string
	index int
}
type Number struct {
	value      int
	indexStart int
	indexEnd   int
}

func main() {
	file, err := os.Open("data.txt")

	var symbols [][]Symbol
	var numbers [][]Number

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// map symbols
		r := regexp.MustCompile(`(?m)&|%|\/|@|\+|\*|=|\$|#|-`)
		matches := r.FindAllStringSubmatchIndex(text, -1)
		if len(matches) != 0 {
			var rowSymbols []Symbol
			for _, symbolIndex := range matches {
				if len(symbolIndex) != 0 {
					rowSymbols = append(rowSymbols, Symbol{
						value: string([]rune(text)[symbolIndex[0]]),
						index: symbolIndex[0],
					})
				}
			}
			symbols = append(symbols, rowSymbols)
		} else {
			symbols = append(symbols, []Symbol{})
		}

		// map numbers
		r = regexp.MustCompile(`(?s)\d+`)
		matches = r.FindAllStringSubmatchIndex(text, -1)
		if len(matches) != 0 {
			var rowNumbers []Number
			for _, symbolIndex := range matches {
				if len(symbolIndex) != 0 {
					value, _ := strconv.Atoi(text[symbolIndex[0]:symbolIndex[1]])
					rowNumbers = append(rowNumbers, Number{
						value:      value,
						indexStart: symbolIndex[0],
						indexEnd:   symbolIndex[1],
					})
				}
			}
			numbers = append(numbers, rowNumbers)
		} else {
			numbers = append(numbers, []Number{})
		}
	}

	fmt.Print(numbers[0])

	file.Close()
}
