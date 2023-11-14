package geometry

import "fmt"

type Coords [2]int

func (c *Coords) Unwrap() (int, int) {
	return c[0], c[1]
}

func (c *Coords) Equals(c2 Coords) bool {
	return c[0] == c2[0] && c[1] == c2[1]
}

func (c *Coords) EqualsPair(x, y int) bool {
	return c[0] == x && c[1] == y
}

func (c *Coords) VectorTo(c2 Coords) (int, int) {
	return c2[0] - c[0], c2[1] - c[1]
}

func NewCoords(x, y int) Coords {
	var a Coords = [2]int{x, y}
	return a
}

func PrintCoordsArray(a [][]Coords) {
	for i := range a {
		for j := range a[i] {
			fmt.Printf("%d,%d  ", a[i][j][0], a[i][j][1])
		}
		fmt.Printf(" |  ")
	}
	fmt.Printf("\n")
}
