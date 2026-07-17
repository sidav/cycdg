package graph

import "cycdg/graph_replacement/geometry"

func (g *Graph) CountDirEdgesAt(x, y int, countIn, countOut bool) int {
	if !countIn && !countOut {
		panic("And wat should I count?")
	}
	count := 0
	for _, dir := range cardinalDirections {
		vx, vy := unwrapCoords(dir)
		otherx, othery := x+vx, y+vy
		if !g.AreCoordsInBounds(otherx, othery) {
			continue
		}
		if countIn {
			if g.IsEdgeDirectedBetweenCoords(otherx, othery, x, y) {
				count++
			}
		}
		if countOut {
			if g.IsEdgeDirectedBetweenCoords(x, y, otherx, othery) {
				count++
			}
		}
	}
	return count
}

func (g *Graph) IsEdgeDirectedByVector(x, y, vx, vy int) bool {
	n := g.GetEdgeByVector(x, y, vx, vy)
	return n.IsActive() && n.IsReverse() == (vx < 0 || vy < 0)
}

func (g *Graph) IsEdgeDirectedBetweenCoords(x, y, tox, toy int) bool {
	vx, vy := tox-x, toy-y
	n := g.GetEdgeByVector(x, y, vx, vy)
	return n.IsActive() && n.IsReverse() == (vx < 0 || vy < 0)
}

func (g *Graph) IsEdgeDirectedFromCoordsToPair(from geometry.Coords, tox, toy int) bool {
	return g.IsEdgeDirectedBetweenCoords(from.X, from.Y, tox, toy)
}

func (g *Graph) IsEdgeDirectedFromCoords(from, to geometry.Coords) bool {
	return g.IsEdgeDirectedBetweenCoords(from.X, from.Y, to.X, to.Y)
}

func (g *Graph) doCoordsHaveIngoingLinksOnly(x, y int) bool {
	for i := range cardinalDirections {
		vx, vy := unwrapCoords(cardinalDirections[i])
		if g.AreCoordsInBounds(x+vx, y+vy) {
			if g.IsEdgeByVectorActive(x, y, vx, vy) {
				if g.IsEdgeDirectedByVector(x, y, vx, vy) {
					return false
				}
			}
		}
	}
	return true
}

func (g *Graph) EnableDirLinkByVector(x, y, vx, vy int) {
	if vx*vy != 0 {
		debugPanic("Diagonal connection?.. %d,%d -> %d,%d", x, y, vx, vy)
	}
	reverse := false
	if vx == -1 {
		x--
		vx = 1
		reverse = true
	} else if vy == -1 {
		y--
		vy = 1
		reverse = true
	}
	g.NodeAtXY(x, y).SetLinkByVector(vx, vy, true, reverse)
}

func (g *Graph) EnableDirLinkByCoords(from, to geometry.Coords) {
	vx, vy := from.VectorTo(to)
	g.EnableDirLinkByVector(from.X, from.Y, vx, vy)
}

func (g *Graph) DisableDirLinkByCoords(from, to geometry.Coords) {
	if !g.IsEdgeDirectedBetweenCoords(from.X, from.Y, to.X, to.Y) {
		debugPanic("Direction unsatisfied")
	}
	vx, vy := from.VectorTo(to)
	g.setLinkByVector(from.X, from.Y, vx, vy, false)
}
