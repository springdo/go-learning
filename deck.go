package main

import "fmt"

// create a new type of DECK
// of type slice of STRING

// this is like  a type extends slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	// var suits []string = []string {}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "2", "3", "4"}

	//  use _ as a special throw away variable
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(numb int) {
	for _, card := range d {
		fmt.Println(card)
	}
}
