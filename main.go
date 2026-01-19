package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")


	//Reading from a file 
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards:=newDeck()
	cards.shuffle()
	cards.print()

}
