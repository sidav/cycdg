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

func (c *Coords) IsAdjacentToXY(x, y int) bool {
	return c.ManhattanDistToXY(x, y) == 1
}

func (c *Coords) IsCardinalToPair(x, y int) bool {
	return c[0] == x || c[1] == y
}

func (c *Coords) ManhattanDistToXY(x, y int) int {
	return intAbs(x-c[0]) + intAbs(y-c[1])
}

func (c *Coords) VectorTo(c2 Coords) (int, int) {
	return c2[0] - c[0], c2[1] - c[1]
}

func NewCoords(x, y int) Coords {
	var a Coords = [2]int{x, y}
	return a
}

func AreCoords2DArraysEqual(a1, a2 [][]Coords) bool {
	for i := range a1 {
		found := false
	nextCoordInA2:
		for j := range a2 {
			for k := range a2[j] {
				if !a2[j][k].Equals(a1[i][k]) {
					continue nextCoordInA2
				}
			}
			found = true
			break
		}
		if !found {
			fmt.Printf("WTF %v is not in %v: iteration %d\n", a1[i], a2, i)
			return false
		}
	}
	return true
}

func AreXYCoordsInCoordsArray(x, y int, coords []Coords) bool {
	for i := range coords {
		if coords[i].EqualsPair(x, y) {
			return true
		}
	}
	return false
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

func (c *Coords) GetRectangleForAnotherCornerCoords(corner Coords) (x, y, w, h int) {
	x, y = c.Unwrap()
	x2, y2 := corner.Unwrap()
	w = intAbs(x2-x) + 1 // +1 because the map is tiled
	h = intAbs(y2-y) + 1
	if x2 < x {
		x = x2
	}
	if y2 < y {
		y = y2
	}
	return x, y, w, h
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
