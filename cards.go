package deck

import (
	"math/rand"
	"strings"
	"time"
)

type Card struct {
	Name, Face, Suit, Glyph string
	Value                   int
}

type Deck struct {
	Cards []Card
}

type Hand struct {
	Cards []Card
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// NewDeck creates a deck of unshuffled cards.
func NewDeck() Deck {
	// d := make([]Card, 52)
	var d []Card
	for _, s := range suits {
		for v := 0; v < 13; v++ {
			f := names[v]
			n := f + " of " + s
			g := glyphs[n]

			c := Card{
				Name:  n,
				Face:  f,
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

func NewHand() Hand {
	return Hand{}
}

func (h *Hand) Draw(d *Deck, n int) {
	h.Cards, d.Cards = d.Cards[0:n], d.Cards[n+1:]
}

func (h Hand) String() string {
	var g []string
	for _, c := range h.Cards {
		g = append(g, c.Glyph)
	}
	return strings.Join(g, " ")
}

func (d Deck) String() string {
	var g []string
	for _, c := range d.Cards {
		g = append(g, c.Glyph)
	}
	return strings.Join(g, " ")
}
