package main

import "fmt"

func main() {
	deck := newDeck()	// 同一パッケージ内では import 不要で別ファイルの関数を読み込める
	hand, _ := deal(deck, 4)
	fmt.Println(hand.shuffle())

	//fileName := "test_deck"
	//hand.saveToFile(fileName)

	//read := readFile(fileName)
	//deck(strings.Split(string(read), ","))
	//fmt.Println(deck)

	//fmt.Println(deckFromReadFile("de"))
}