//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker //special case to watch for
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (card Card) String() string {
	if card.Suit == Joker {
		return card.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", card.Rank, card.Suit)
}

func NewDeck(options ...func([]Card) []Card) []Card {
	var deck []Card
	for _, suitTBA := range suits {
		for rankTBA := minRank; rankTBA <= maxRank; rankTBA++ {
			deck = append(deck, Card{Suit: suitTBA, Rank: rankTBA})
		}
	}
	for _, option := range options {
		deck = option(deck)
	}
	return deck
}

func AbsoluteRank(c Card) int {
	//          0-3           1-13           1-13
	return int(c.Suit)*int(maxRank) + int(c.Rank)
	//each rank has 13 unique values
	//Spades   1-13
	//diamonds 14-27 and so on
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return AbsoluteRank(cards[i]) < AbsoluteRank(cards[j])
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(deck []Card) []Card {
	shuffled := make([]Card, len(deck)) //New shuffled deck with empty cards
	//create a random variable for use, seeded with deck passed in
	random := rand.New(rand.NewSource(int64(AbsoluteRank(deck[0]))))

	//Use random variable to create a slice containing numbers up to the length
	//of the deck in a random order
	newOrder := random.Perm(len(deck))

	//reassign each card to a new position based on above random slice
	for i, j := range newOrder {
		shuffled[i] = deck[j]
	}
	return shuffled
}

func AddJokers(numberWanted int) func([]Card) []Card {
	//NewDeck(AddJokers(3)) instead of NewDeck(AddJokers, AddJokers, AddJokers)
	return (func(deck []Card) []Card {
		for i := 0; i < numberWanted; i++ {
			deck = append(deck, Card{Rank: Rank(i + 14), Suit: Joker})
		} //Rank for joker exists to differentiate them for games that require it
		return deck
	})

}

func FilterCards(f func(deck Card) bool) func([]Card) []Card {
	return func(deck []Card) []Card {
		var filtered []Card
		for _, card := range deck {
			if !f(card) {
				filtered = append(filtered, card)
			}
		}
		return filtered
	}
}

// function duplicates all cards in current deck
// and adds them to the deck
func AddDeck(n int) func([]Card) []Card {
	return func(deck []Card) []Card {
		var toAdd []Card
		for i := 0; i < n; i++ {
			//duplicates every card n times
			toAdd = append(toAdd, deck...)
		}
		return toAdd
	}
}
