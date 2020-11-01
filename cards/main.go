package main

import "fmt"

func main() {
	card := cardContent()
	fmt.Println(card)
}

func cardContent() string {
	return "Ace of Spades"
}