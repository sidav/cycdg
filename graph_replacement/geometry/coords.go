package geometry

import "fmt"

type Coords struct {
	X, Y int
}

func NewCoords(x, y int) Coords {
	return Coords{X: x, Y: y}
}

func (c Coords) Unwrap() (int, int) {
	return c.X, c.Y
}

func (c Coords) ToString() string {
	return fmt.Sprintf("%d, %d", c.X, c.Y)
}

func (c Coords) Equals(c2 Coords) bool {
	return c.X == c2.X && c.Y == c2.Y
}

func (c Coords) EqualsPair(x, y int) bool {
	return c.X == x && c.Y == y
}

func (c Coords) IsAdjacentTo(c2 Coords) bool {
	return c.ManhattanDistTo(c2) == 1
}

func (c Coords) IsAdjacentToXY(x, y int) bool {
	return c.ManhattanDistToXY(x, y) == 1
}

func (c Coords) IsCardinalToPair(x, y int) bool {
	return c.X == x || c.Y == y
}

func (c Coords) ManhattanDistToXY(x, y int) int {
	return intAbs(x-c.X) + intAbs(y-c.Y)
}

func (c Coords) ManhattanDistTo(c2 Coords) int {
	return intAbs(c2.X-c.X) + intAbs(c2.Y-c.Y)
}

func (c Coords) VectorTo(c2 Coords) (int, int) {
	return c2.X - c.X, c2.Y - c.Y
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
			fmt.Printf("%d,%d  ", a[i][j].X, a[i][j].Y)
		}
		fmt.Printf(" |  ")
	}
	fmt.Printf("\n")
}

func (c Coords) IsAdjacentToRectangleCorner(rx, ry, w, h int) bool {
	x, y := c.Unwrap()
	return (x == rx || x == rx+w-1) && (y == ry+1 || y == ry+h-2) ||
		(x == rx+1 || x == rx+w-2) && (y == ry || y == ry+h-1)
}

// note: it's not IN rectangle!
func (c Coords) IsOnRectangle(rx, ry, w, h int) bool {
	x, y := c.Unwrap()
	if x < rx || x >= rx+w || y < ry || y >= ry+h {
		return false
	}
	return x == rx || x == rx+w-1 || y == ry || y == ry+h-1

}

func (c Coords) GetRectangleForAnotherCornerCoords(corner Coords) (x, y, w, h int) {
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
