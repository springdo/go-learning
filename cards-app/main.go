package main

func main() {

	// card := newCard()
	// or
	// var card = "Ace of Spades"

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)

	// cards.saveToFile("my-cards")

	// otherCards := newDeckFromFile("my-card")
	// otherCards.print()

}
