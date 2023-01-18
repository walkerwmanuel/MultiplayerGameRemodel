package cardLogic

var (
	suites = []string{"Heart", "Club", "Diamond", "Spade"}
	values = []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
)

// Card - the type to represent a card
// Suite - suite of a card - Clubs, Spades, Diamonds, Hearts
// Value - value of a card - Ace, King, Queen, Jack, Ten...Two
type Card struct {
	Suite string
	Value string
}

// Deck - the type for a deck
type Deck struct {
	Cards []Card
}
