package deck

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCardStringfy(t *testing.T) {
	cases := []struct {
		card Card
		want string
	}{
		{
			card: Card{
				Name:      "Goblin Attack",
				Clockwise: false,
				Types:     []Type{Monster},
				Shape: []Shape{
					{
						{true, false, false},
						{false, true, false},
						{false, false, true},
					},
				},
			},
			want: `CARD: Goblin Attack
TYPE: 👾
ORIENTATION: 🔄

[]    
  []  
    []`,
		},
		{
			card: Card{
				Name:      "Bugbear Assault",
				Clockwise: true,
				Types:     []Type{Monster},
				Shape: []Shape{
					{
						{true, false, true},
						{true, false, true},
					},
				},
			},
			want: `CARD: Bugbear Assault
TYPE: 👾
ORIENTATION: 🔃

[]  []
[]  []`,
		},
		{
			card: Card{
				Name:  "Temple Ruins",
				Types: []Type{Ruins},
			},
			want: `CARD: Temple Ruins

 |_,.
   |||
   ||||
   ||||
 _/____\_
|________|`,
		},
		{
			card: Card{
				Name:  "Great River",
				Time:  1,
				Coin:  true,
				Types: []Type{Water},
				Shape: []Shape{
					{
						{true},
						{true},
						{true},
					},
					{
						{false, false, true},
						{false, true, true},
						{true, true, false},
					},
				},
			},
			want: `CARD: Great River
TYPE: 💧
TIME: 1⏳

[]      |      []
[]  💰  |    [][]
[]      |  [][]  `,
		},
		{
			card: Card{
				Name:  "Farmland",
				Time:  1,
				Coin:  true,
				Types: []Type{Farm},
				Shape: []Shape{
					{
						{true},
						{true},
					},
					{
						{false, true, false},
						{true, true, true},
						{false, true, false},
					},
				},
			},
			want: `CARD: Farmland
TYPE: 🌱
TIME: 1⏳

[]      |    []  
[]  💰  |  [][][]
        |    []  `,
		},
		{
			card: Card{
				Name:  "Forgotten Forest",
				Time:  1,
				Coin:  true,
				Types: []Type{Forest},
				Shape: []Shape{
					{
						{true, false},
						{false, true},
					},
					{
						{true, false},
						{true, true},
						{false, true},
					},
				},
			},
			want: `CARD: Forgotten Forest
TYPE: 🌲
TIME: 1⏳

[]        |  []  
  []  💰  |  [][]
          |    []`,
		},
		{
			card: Card{
				Name:  "Treetop Village",
				Time:  2,
				Coin:  false,
				Types: []Type{Forest, Village},
				Shape: []Shape{
					{
						{false, false, true, true},
						{true, true, true, false},
					},
				},
			},
			want: `CARD: Treetop Village
TYPE: 🌲 | 🏠
TIME: 2⏳

    [][]
[][][]  `,
		},
		{
			card: Card{
				Name:  "Rift Land",
				Time:  0,
				Coin:  false,
				Types: []Type{Forest, Village, Farm, Water, Monster},
				Shape: []Shape{
					{
						{true},
					},
				},
			},
			want: `CARD: Rift Land
TYPE: 🌲 | 🏠 | 🌱 | 💧 | 👾
TIME: 0⏳

[]`,
		},
	}

	for _, c := range cases {
		t.Run(c.card.Name, func(t *testing.T) {
			out := c.card.String()

			diff := cmp.Diff(out, c.want)

			if out != c.want {
				t.Errorf("want:\n%s\n\ngot:\n%s\n", c.want, diff)
			}
		})
	}
}
