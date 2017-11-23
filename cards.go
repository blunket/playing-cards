package deck

type Card struct {
	Name, Suit string
	Value      int
}

type Deck []Card

func NewDeck() Deck {
	var d []Card
	for _, s := range suits {
		for v := 0; v < 13; v++ {
			c := Card{
				Name:  names[v],
				Suit:  s,
				Value: v + 1,
			}
			d = append(d, c)
		}
	}
	return d
}
