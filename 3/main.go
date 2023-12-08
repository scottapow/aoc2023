package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type symbol struct {
	value string
	index int
}
type number struct {
	value      int
	indexStart int
	indexEnd   int
}
type engine struct {
	symbols [][]symbol
	numbers [][]number
}

func (e *engine) mapSymbols(line string) {
	// map symbols
	r := regexp.MustCompile(`(?m)&|%|\/|@|\+|\*|=|\$|#|-`)
	matches := r.FindAllStringSubmatchIndex(line, -1)
	if len(matches) != 0 {
		var rowSymbols []symbol
		for _, symbolIndex := range matches {
			if len(symbolIndex) != 0 {
				rowSymbols = append(rowSymbols, symbol{
					value: string([]rune(line)[symbolIndex[0]]),
					index: symbolIndex[0],
				})
			}
		}
		e.symbols = append(e.symbols, rowSymbols)
	} else {
		e.symbols = append(e.symbols, []symbol{})
	}
}
func (e *engine) mapNumbers(line string) {
	// map numbers
	r := regexp.MustCompile(`(?s)\d+`)
	matches := r.FindAllStringSubmatchIndex(line, -1)
	if len(matches) != 0 {
		var rowNumbers []number
		for _, symbolIndex := range matches {
			if len(symbolIndex) != 0 {
				value, _ := strconv.Atoi(line[symbolIndex[0]:symbolIndex[1]])
				rowNumbers = append(rowNumbers, number{
					value:      value,
					indexStart: symbolIndex[0],
					indexEnd:   symbolIndex[1],
				})
			}
		}
		e.numbers = append(e.numbers, rowNumbers)
	} else {
		e.numbers = append(e.numbers, []number{})
	}
}

func main() {
	file, err := os.Open("data.txt")
	e := &engine{
		symbols: [][]symbol{},
		numbers: [][]number{},
	}

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		e.mapNumbers(line)
		e.mapSymbols(line)
	}

	fmt.Print(e.numbers[0])

	file.Close()
}
