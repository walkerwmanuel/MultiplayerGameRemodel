package cardLogic

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// NewDeck - creates a deck with 52 cards
func NewDeck() *Deck {
	d := Deck{}
	//Makes an empty map of 52 type card
	d.Cards = make([]Card, 52)
	index := 0
	//Itteration through all 52 types of cards and stores it
	for i := range suites {
		for j := range values {
			d.Cards[index] = NewCard(suites[i], values[j])
			index++
		}
	}
	return &d
}

// Shuffle - Shuffles through deck and places the cards in a random order
func (d *Deck) Shuffle() {
	if len(d.Cards) == 0 {
		return
	}

	// iterate through each card in Cards of deck
	for i := range d.Cards {
		// pull out card of current iteration index 0...len(d.Cards)
		card := d.Cards[i]
		// creates a new (random) position from 0...len(d.Cards)
		newPos, _ := rand.Int(rand.Reader, big.NewInt(int64(len(d.Cards))))
		// convert newPos to int
		newPosInt := newPos.Uint64()
		// pull out card in new position
		otherCard := d.Cards[newPosInt]
		// swap them
		d.Cards[i] = otherCard
		d.Cards[newPosInt] = card
	}
}

// PrintDeck prints the contents of a deck
func (d *Deck) PrintDeck() {
	fmt.Println()
	for i := range d.Cards {
		fmt.Printf("[%s:%s] ", d.Cards[i].Suite, d.Cards[i].Value)
	}
	fmt.Println()
}

// NewCard - makes a new card
func NewCard(v, s string) Card {
	c := Card{
		Suite:  s,
		Filler: "of",
		Value:  v,
	}
	return c
}
