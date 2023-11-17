package replacement

import (
	"fmt"
)

var (
	cardinalDirections = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	diagonalDirections = [4][2]int{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
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

func sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func intabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
