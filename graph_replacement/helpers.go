package replacement

import (
	"fmt"
)

type direction uint8

const (
	dirE direction = iota
	dirS
	dirW
	dirN
)

func debugPanic(msg string, args ...interface{}) {
	panic(sprintf(msg, args...))
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

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minint(a, b int) int {
	if a < b {
		return a
	}
	return b
}
