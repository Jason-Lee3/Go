package main

func main() {
	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// temp := cards.toString()
	// fmt.Println([]byte(temp))
	// cards.saveToFile("file_1")
	// cards := newDeckFromFile("file_11")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))
}
