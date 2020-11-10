package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	var cards deck
	cardsSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(d); i++ {
		randIndex := rand.Intn(len(d))

		d[i], d[randIndex] = d[randIndex], d[i]
	}

	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
