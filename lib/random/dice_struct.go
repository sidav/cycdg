package random

import "fmt"

type Dice struct {
	// 2d6 + 5 - Number is 2, 6 Sides, Modifier is 5
	Number, Sides, Modifier int
}

func NewDice(dnum, dval, dmod int) *Dice {
	return &Dice{Number: dnum, Sides: dval, Modifier: dmod}
}

func (d *Dice) Alter(num, sides, mod int) {
	d.Number = num
	d.Sides = sides
	d.Modifier = mod
}

func (d *Dice) Roll(rnd PRNG) int {
	return rnd.RollDice(d.Number, d.Sides, d.Modifier)
}

func (d *Dice) GetMinimumPossible() int {
	return d.Number + d.Modifier
}

func (d *Dice) GetMaximumPossible() int {
	return d.Number*d.Sides + d.Modifier
}

func (d *Dice) GetShortDescriptionString() string {
	if d.Modifier < 0 {
		return fmt.Sprintf("%dd%d%d", d.Number, d.Sides, d.Modifier)
	}
	if d.Modifier > 0 {
		return fmt.Sprintf("%dd%d+%d", d.Number, d.Sides, d.Modifier)
	}
	return fmt.Sprintf("%dd%d", d.Number, d.Sides)
}

func (d *Dice) GetDescriptionString() string {
	return fmt.Sprintf("%s %d-%d", d.GetShortDescriptionString(), d.GetMinimumPossible(), d.GetMaximumPossible())
}
