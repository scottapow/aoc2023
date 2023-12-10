package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
)

type card struct {
	Winning []int `json:"winning"`
	Pool    []int `json:"pool"`
}

func main() {
	file, err := os.Open("data.json")
	var cards []card
	winTotal := 0
	cardsInstancesTotal := 0

	if err != nil {
		panic(err)
	}

	fileBytes, _ := io.ReadAll(file)
	json.Unmarshal(fileBytes, &cards)

	calculateWinValues(&cards, &winTotal)
	calculateCardWinsCount(&cards, &cardsInstancesTotal)
	fmt.Println(winTotal)
	fmt.Println(cardsInstancesTotal)

	file.Close()
}

func calculateWinValues(cards *[]card, total *int) {
	for _, card := range *cards {
		cardTotal := 0
		for _, n := range card.Pool {
			wins := slices.Contains(card.Winning, n)
			if wins {
				if cardTotal != 0 {
					cardTotal = cardTotal * 2
				} else {
					cardTotal = 1
				}
			}
		}
		*total += cardTotal
	}
}

func calculateCardWinsCount(cards *[]card, total *int) {
	m := make(map[int]int)
	for cardIndex, card := range *cards {
		var cardNumber int = cardIndex + 1
		m[cardNumber]++
		for i := 0; i < m[cardNumber]; i++ {
			cardWinTotal := 0
			for _, n := range card.Pool {
				wins := slices.Contains(card.Winning, n)
				if wins {
					cardWinTotal++
					updateCardNumner := cardNumber + cardWinTotal
					m[updateCardNumner]++
				}
			}
		}
	}

	for _, v := range m {
		*total += v
	}
}
