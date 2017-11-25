package deck

import (
	"math/rand"
	"time"
)

type Card struct {
	Name, Suit, Glyph string
	Value             int
}

type Deck struct {
	Cards []Card
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NewDeck() Deck {
	var d []Card
	for _, s := range suits {
		for v := 0; v < 13; v++ {
			n := names[v]
			full := n + " of " + s
			g := glyphs[full]

			c := Card{
				Name:  n,
				Suit:  s,
				Value: v + 1,
				Glyph: g,
			}

			d = append(d, c)
		}
	}
	return Deck{Cards: d}
}

func (d *Deck) Shuffle() {
	c := d.Cards
	for i := range c {
		j := rand.Intn(i + 1)
		c[i], c[j] = c[j], c[i]
	}
}
