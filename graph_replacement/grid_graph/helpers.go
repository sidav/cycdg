package graph

import (
	"cycdg/lib/random"
	"fmt"
)

var (
	rnd                random.PRNG
	cardinalDirections = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	diagonalDirections = [4][2]int{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
)

func debugPanic(msg string, args ...interface{}) {
	panic(sprintf(msg, args...))
}

func SetRandom(r random.PRNG) {
	rnd = r
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
