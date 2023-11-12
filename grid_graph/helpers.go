package graph

import (
	"cycdg/lib/random"
	"fmt"
)

type direction uint8

const (
	dirE direction = iota
	dirS
	dirW
	dirN
)

var (
	rnd                random.PRNG
	cardinalDirections = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	diagonalDirections = [4][2]int{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
)

type Coords [2]int

func (c *Coords) unwrap() (int, int) {
	return c[0], c[1]
}

func (c *Coords) equals(c2 Coords) bool {
	return c[0] == c2[0] && c[1] == c2[1]
}

func newCoords(x, y int) Coords {
	var a Coords = [2]int{x, y}
	return a
}

func printCoordsArray(a [][]Coords) {
	for i := range a {
		for j := range a[i] {
			fmt.Printf("%d,%d  ", a[i][j][0], a[i][j][1])
		}
		fmt.Printf(" |  ")
	}
	fmt.Printf("\n")
}

func SetRandom(r random.PRNG) {
	rnd = r
}

func VectorToDirection(vx, vy int) direction {
	if vx == 0 && vy == -1 {
		return dirN
	}
	if vx == 0 && vy == 1 {
		return dirS
	}
	if vx == -1 && vy == 0 {
		return dirW
	}
	if vx == 1 && vy == 0 {
		return dirE
	}
	debugPanic("No such direction: %d,%d", vx, vy)
	panic(nil)
}

func getRandomVectorByFunc(appropriate func(vx, vy int) bool) (int, int) {
	var candidates [][2]int
	for _, d := range cardinalDirections {
		if appropriate(d[0], d[1]) {
			candidates = append(candidates, d)
		}
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind][0], candidates[ind][1]
}

// note: it's not IN rectangle!
func areCoordsOnRectangle(x, y, rx, ry, w, h int) bool {
	if x < rx || x >= rx+w || y < ry || y >= ry+h {
		return false
	}
	return x == rx || x == rx+w-1 || y == ry || y == ry+h-1
}

func areCoordsAdjacent(x1, y1, x2, y2 int) bool {
	dx := x2 - x1
	dy := y2 - y1
	return dx*dy == 0 && (dx == -1 || dx == 1 || dy == -1 || dy == 1)
}

func rotateCoordsCW(x, y int) (int, int) {
	return y, -x
}

func rotateCoordsCCW(x, y int) (int, int) {
	return -y, x
}

func randomUnitVector() (int, int) {
	return rnd.RandomUnitVectorInt(false)
}

func sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func intabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetAllRectCoordsClockwise(x, y, w, h int) [][2]int {
	rightX, bottomY := x+w-1, y+h-1
	totalCoords := 2*w + 2*(h-2)
	coords := make([][2]int, totalCoords)
	currCoord := 0
	vx, vy := 1, 0
	currX, currY := x, y
	for currCoord < totalCoords {
		coords[currCoord][0], coords[currCoord][1] = currX, currY
		currCoord++
		currX += vx
		currY += vy
		if currX == x && currY == y || currX == x && currY == bottomY ||
			currX == rightX && currY == y || currX == rightX && currY == bottomY {
			vx, vy = -vy, vx
		}
	}
	return coords
}

func findInCoordsArray(x, y int, coords [][2]int) int {
	for i := range coords {
		if coords[i][0] == x && coords[i][1] == y {
			return i
		}
	}
	debugPanic("No coords found!\nRequested %d,%d from %v", x, y, coords)
	panic(nil)
}

func unwrapCoords(coords [2]int) (int, int) {
	return coords[0], coords[1]
}
