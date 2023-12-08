package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const rowMax int = 140

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
	r := regexp.MustCompile(`(?mi)[^.\d\n]+`)
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

func (e *engine) calculateAdjacentTotals() int {
	adjacentTotals := 0
	for rowIndex, row := range e.numbers {
		isFirstRow := rowIndex == 0
		isLastRow := rowIndex == len(e.numbers)-1
		for _, n := range row {
			partNumberFound := false
			startIndex := int(math.Max(float64(n.indexStart-1), 0))
			endIndex := int(math.Min(float64(n.indexEnd), float64(rowMax)))

			// check row above
			if !isFirstRow {
				prevRowSymbols := e.symbols[rowIndex-1]
				for _, s := range prevRowSymbols {
					if s.index >= startIndex && s.index <= endIndex {
						adjacentTotals += n.value
						partNumberFound = true
						break
					}
				}
				if partNumberFound {
					continue
				}
			}

			rowSymbols := e.symbols[rowIndex]
			// check left and right
			for _, s := range rowSymbols {
				if s.index == startIndex || s.index == endIndex {
					adjacentTotals += n.value
					partNumberFound = true
					break
				}
			}
			if partNumberFound {
				continue
			}

			// check row below
			if !isLastRow {
				nextRowSymbols := e.symbols[rowIndex+1]
				for _, s := range nextRowSymbols {
					if s.index >= startIndex && s.index <= endIndex {
						adjacentTotals += n.value
					}
				}
			}
		}
	}

	return adjacentTotals
}

func (e *engine) calculateGearRatios() int {
	gearRatioTotal := 0
	for rowIndex, row := range e.symbols {
		isFirstRow := rowIndex == 0
		isLastRow := rowIndex == len(e.numbers)-1

		for _, s := range row {
			symbolIndexStart := s.index
			symbolIndexEnd := int(math.Min(float64(s.index+1), float64(rowMax)))
			var adjacentNumbers []int

			// check row above
			if !isFirstRow {
				prevRowNumbers := e.numbers[rowIndex-1]
				for _, n := range prevRowNumbers {
					if n.indexStart <= symbolIndexEnd && n.indexEnd >= symbolIndexStart {
						adjacentNumbers = append(adjacentNumbers, n.value)
					}
				}
			}

			rowNumbers := e.numbers[rowIndex]
			// check left and right
			for _, n := range rowNumbers {
				if n.indexEnd == symbolIndexStart || n.indexStart == symbolIndexEnd {
					adjacentNumbers = append(adjacentNumbers, n.value)
				}
			}

			// check row below
			if !isLastRow {
				nextRowNumbers := e.numbers[rowIndex+1]
				for _, n := range nextRowNumbers {
					if n.indexStart <= symbolIndexEnd && n.indexEnd >= symbolIndexStart {
						adjacentNumbers = append(adjacentNumbers, n.value)
					}
				}
			}

			if len(adjacentNumbers) == 2 {
				gearRatioTotal += (adjacentNumbers[0] * adjacentNumbers[1])
			}
		}
	}

	return gearRatioTotal
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

	adjacentTotals := e.calculateAdjacentTotals()
	gearRatioTotals := e.calculateGearRatios()
	fmt.Println("adjacentTotals", adjacentTotals)
	fmt.Println("gearRatioTotals", gearRatioTotals)

	file.Close()
}
