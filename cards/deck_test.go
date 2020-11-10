package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	sut := newDeck()
	if len(sut) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(sut))
	}
	// 本当は全要素を比較した方が良い(reflect.DeepEqualを使って比較する)
	// 参考:https://cipepser.hatenablog.com/entry/go-deepequal-slice-map
	if sut[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", sut[0])
	}
	if sut[len(sut) - 1] != "Four of Hearts" {
		t.Errorf("Expected Four of Hearts, but got %v", sut[len(sut) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesging"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)
	if len(loadedDeck) != 15 {
		t.Errorf("expected 16, but got %v", len(loadedDeck))
	}

	// 上のassertでテストが失敗しても、ファイルクリーンアップは実行される
	os.Remove(fileName)
}