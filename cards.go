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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// New creates a new Deck of unshuffled cards.
func New() Deck {
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

func (dest *Deck) DrawFrom(src *Deck, n int) {
	dest.Cards, src.Cards = src.Cards[0:n], src.Cards[n:]
}

func (d Deck) String() string {
	var g []string
	for _, c := range d.Cards {
		g = append(g, c.Glyph)
	}
	return strings.Join(g, " ")
}
