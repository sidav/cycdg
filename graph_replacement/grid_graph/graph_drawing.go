package graph

import . "cycdg/graph_replacement/geometry"

func (g *Graph) drawCardinalConnectedLine(x1, y1, x2, y2 int, directed bool) {
	vx := 0
	if x2 != x1 {
		vx = (x2 - x1) / intabs(x2-x1)
	}
	vy := 0
	if y2 != y1 {
		vy = (y2 - y1) / intabs(y2-y1)
	}
	g.EnableNode(x1, y1)
	x := x1
	y := y1
	for vx != 0 && x != x2 {
		g.EnableNodeByVector(x, y, vx, vy)
		g.EnableDirLinkByVector(x, y, vx, vy)
		x += vx
	}
	for vy != 0 && y != y2 {
		g.EnableNodeByVector(x, y, vx, vy)
		g.EnableDirLinkByVector(x, y, vx, vy)
		y += vy
	}
}

// Should only add links/nodes, not remove!
func (g *Graph) DrawConnectedDirectionalRect(x, y, w, h int, ccw bool) {
	rghX, botY := x+w-1, y+h-1
	if ccw {
		g.drawCardinalConnectedLine(x, y, x, botY, true)
		g.drawCardinalConnectedLine(x, botY, rghX, botY, true)
		g.drawCardinalConnectedLine(rghX, botY, rghX, y, true)
		g.drawCardinalConnectedLine(rghX, y, x, y, true)
	} else {
		g.drawCardinalConnectedLine(x, y, rghX, y, true)
		g.drawCardinalConnectedLine(rghX, y, rghX, botY, true)
		g.drawCardinalConnectedLine(rghX, botY, x, botY, true)
		g.drawCardinalConnectedLine(x, botY, x, y, true)
	}
}

// Draws two paths from source to sink alongside the rect.
// Result example (O is source, S is sink):
// N > S < N
// ^       ^
// N       N
// ^       ^
// O > N > N
func (g *Graph) drawBiсonnectedDirectionalRect(x, y, w, h, sourceX, sourceY, sinkX, sinkY int) {
	allCoords := GetAllRectCoordsClockwise(x, y, w, h)
	sourceIndex := findInCoordsArray(sourceX, sourceY, allCoords)
	sinkIndex := findInCoordsArray(sinkX, sinkY, allCoords)
	g.EnableNode(sourceX, sourceY)
	// first path: clockwise
	index := sourceIndex
	for index != sinkIndex {
		nextIndex := (index + 1) % len(allCoords)
		currX, currY := unwrapCoords(allCoords[index])
		nextX, nextY := unwrapCoords(allCoords[nextIndex])
		vx, vy := nextX-currX, nextY-currY
		g.EnableNodeByVector(currX, currY, vx, vy)
		g.EnableDirLinkByVector(currX, currY, vx, vy)
		index = nextIndex
	}
	// second path: counter-clockwise
	index = sourceIndex
	for index != sinkIndex {
		nextIndex := index - 1
		if nextIndex < 0 {
			nextIndex += len(allCoords)
		}
		currX, currY := unwrapCoords(allCoords[index])
		nextX, nextY := unwrapCoords(allCoords[nextIndex])
		vx, vy := nextX-currX, nextY-currY
		g.EnableNodeByVector(currX, currY, vx, vy)
		g.EnableDirLinkByVector(currX, currY, vx, vy)
		index = nextIndex
	}
}

func (g *Graph) DrawBiсonnectedDirectionalRect(x, y, w, h int, source, sink Coords) {
	g.drawBiсonnectedDirectionalRect(x, y, w, h, source[0], source[1], sink[0], sink[1])
}

func (g *Graph) DrawEnabledConnectedCardinalLine(from, to Coords) {
	if !from.IsCardinalToPair(to.Unwrap()) {
		debugPanic("%v is not cardinal to %v!", from, to)
	}
	fx, fy := from.Unwrap()
	tx, ty := to.Unwrap()
	g.drawCardinalConnectedLine(fx, fy, tx, ty, true)
}
