// Deck of cards for Cartographers game
package deck

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Type uint8

const (
	_ Type = iota
	Water
	Farm
	Village
	Forest
	Monster
	Ruins
)

type Shape [][]bool
type Shapes []Shape

type Card struct {
	Name      string `json:"name"`
	Time      uint8  `json:"time,omitempty"`
	Coin      bool   `json:"coin,omitempty"`
	Clockwise bool   `json:"clockwise,omitempty"`
	Shape     Shapes `json:"shape,omitempty"`
	Types     []Type `json:"types,omitempty"`
}

type Deck []Card

const ruinsAsciiArt string = `
 |_,.
   |||
   ||||
   ||||
 _/____\_
|________|`

func (t Type) String() string {
	switch t {
	case Water:
		return "💧"
	case Farm:
		return "🌱"
	case Village:
		return "🏠"
	case Forest:
		return "🌲"
	case Monster:
		return "👾"
	case Ruins:
		return "🏛"
	default:
		return "??"
	}
}

func (s Shape) String() string {
	return strings.Join(s.Lines(), "\n")
}

func (s Shape) Lines() []string {
	var lines []string

	for _, row := range s {
		var line strings.Builder

		for _, col := range row {
			if col {
				line.WriteString("[]")
			} else {
				line.WriteString(padding(2))
			}
		}

		lines = append(lines, line.String())
	}

	return lines
}

// String stringfies an array of shapes. Using the following rule:
// - When a card has a single shape, then just return its string value
// - When a card has two distinct shapes, joins each line
// into a single one, while adding a coin to the first shape
// - Anything else, return an empty string
func (s Shapes) String() string {
	if len(s) == 1 {
		return s[0].String()
	}

	var sb strings.Builder

	if len(s) == 2 {
		lines1 := s[0].Lines()
		lines2 := s[1].Lines()

		size1 := len(lines1)
		size2 := len(lines2)

		highest := max(size1, size2)

		for i := 0; i < highest; i++ {
			if i < size1 {
				sb.WriteString(lines1[i])
			} else {
				l := len(lines1[0])
				pad := padding(l)
				sb.WriteString(pad)
			}

			sb.WriteString(padding(2))

			if i == 1 {
				sb.WriteString("💰")
			} else {
				sb.WriteString(padding(2))
			}

			sb.WriteString(padding(2))
			sb.WriteString("|")
			sb.WriteString(padding(2))

			if i < size2 {
				sb.WriteString(lines2[i])
			} else {
				l := len(lines2[0])
				pad := padding(l)
				sb.WriteString(pad)
			}

			if i < highest-1 {
				sb.WriteString("\n")
			}
		}
	}

	return sb.String()
}

func (c Card) IsMonster() bool {
	return len(c.Types) == 1 && c.Types[0] == Monster
}

func (c Card) IsRuins() bool {
	return len(c.Types) == 1 && c.Types[0] == Ruins
}

func (c Card) String() string {
	var sb strings.Builder

	sb.WriteString("CARD: ")
	sb.WriteString(c.Name)

	if c.IsRuins() {
		sb.WriteString("\n")
		sb.WriteString(ruinsAsciiArt)

		return sb.String()
	}

	types := make([]string, len(c.Types))
	for i, t := range c.Types {
		types[i] = t.String()
	}

	sb.WriteString("\nTYPE: ")
	sb.WriteString(strings.Join(types, " | "))

	if c.IsMonster() {
		sb.WriteString("\nORIENTATION: ")
		if c.Clockwise {
			sb.WriteString("🔃")
		} else {
			sb.WriteString("🔄")
		}
	} else {
		sb.WriteString("\nTIME: ")
		sb.WriteString(fmt.Sprint(c.Time))
		sb.WriteString("⏳")
	}

	sb.WriteString("\n\n")

	sb.WriteString(c.Shape.String())

	return sb.String()
}

type data struct {
	Monsters Deck `json:"monsters,omitempty"`
	Explore  Deck `json:"explore,omitempty"`
}

func NewDeck() (Deck, error) {
	f, err := os.Open("cards.json")
	if err != nil {
		return nil, fmt.Errorf("opening cards.json file: %v", err)
	}

	var d data
	err = json.NewDecoder(f).Decode(&d)
	if err != nil {
		return nil, fmt.Errorf("parsing json: %v", err)
	}

	var out Deck
	out = append(out, d.Monsters...)
	out = append(out, d.Explore...)

	return out, nil
}

func padding(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		sb.WriteRune(' ')
	}

	return sb.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
