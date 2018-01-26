package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of DECK
// of type slice of STRING

// this is like  a type extends slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	// var suits []string = []string {}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7"}

	//  use _ as a special throw away variable
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function attaches the print func to to the type deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}

}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	// for type conversion, add the type you want and open bracket with the value
	// eg to cast to bite slice then []byte(THING TO CONVERT)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// double return type
func deal(d deck, handSize int) (deck, deck) {

	// array[startIndexInclusive : upToNotIncluding]
	// eg [0: 2] is index 0, 1.

	// return the the drawn cards , and remaining deck
	return d[:handSize], d[handSize:]
}

func newDeckFromFile(filename string) deck {
	biteslice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR - ", err)
		os.Exit(1)
	}
	s := strings.Split(string(biteslice), ", ")
	return deck(s)
}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().Unix())
	randoCalrisian := rand.New(source)
	for i := range d {
		newPosition := randoCalrisian.Intn(len(d) - 1)

		// the non go way
		current := d[i]
		d[i] = d[newPosition]
		d[newPosition] = current
		// the go way ....
		// d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
