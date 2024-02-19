package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	var ex1 Card
	var ex2 Card
	var ex3 Card
	var ex4 Card
	var ex5 Card
	ex1.Rank, ex1.Suit = Ace, Club
	ex2.Rank, ex2.Suit = Five, Diamond
	ex3.Rank, ex3.Suit = Eight, Heart
	ex4.Rank, ex4.Suit = Queen, Spade
	ex5.Suit = Suit(Joker)
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)
	fmt.Println(ex4)
	fmt.Println(ex5)

	// Output:
	// Ace of Clubs
	// Five of Diamonds
	// Eight of Hearts
	// Queen of Spades
	// Joker
}

func TestDefaultSort(t *testing.T) {
	deck := NewDeck(DefaultSort)
	expect := Card{Rank: Ace, Suit: Spade}
	if deck[0] != expect {
		t.Error("Expected Ace of Spades as first card...\n Received: ", deck[0])
	}
}

func TestAddJokers(t *testing.T) {
	//create a deck with jokers and shuffle
	deck := NewDeck(AddJokers(15))
	jCount := 0
	for _, card := range deck {
		//fmt.Println(card.Rank)
		if card.Suit == Joker {
			jCount++
		}
	}
	if jCount != 15 {
		t.Error("\nExpected 15 Jokers in deck!\nFound: ", jCount)
	}
}

func TestFilterCards(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Ace || card.Rank == Four
	}
	deck := NewDeck(AddJokers(2), FilterCards(filter))
	for _, card := range deck {
		if card.Rank == Ace || card.Rank == Four {
			t.Error("Expected to find no Aces or 4s")
		}
	}
}

func TestAddDeck(t *testing.T) {
	deck := NewDeck(AddDeck(4))
	if len(deck) != 13*4*3 {
		t.Errorf("\nDeck length Received: %d\n Deck length Expected: %d", len(deck), 13*4*3)
	}
}
