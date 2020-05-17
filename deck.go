package main

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"1", "2", "3", "4"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		println(card)
	}
}
