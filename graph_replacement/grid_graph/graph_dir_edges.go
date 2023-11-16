package graph

import "cycdg/graph_replacement/grid_graph/geometry"

func (g *Graph) IsEdgeByVectorDirectional(x, y, vx, vy int) bool {
	n := g.GetEdgeByVector(x, y, vx, vy)
	return n.IsDirectional()
}

func (g *Graph) CountDirEdgesAt(x, y int, countIn, countOut bool) int {
	if !countIn && !countOut {
		panic("And wat should I count?")
	}
	count := 0
	for _, dir := range cardinalDirections {
		vx, vy := unwrapCoords(dir)
		otherx, othery := x+vx, y+vy
		if !g.areCoordsInBounds(otherx, othery) {
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

func (g *Graph) isEdgeByVectorDirectionalAndActive(x, y, vx, vy int) bool {
	n := g.GetEdgeByVector(x, y, vx, vy)
	return n.IsActive() && n.IsDirectional()
}

func (g *Graph) IsEdgeDirectedByVector(x, y, vx, vy int) bool {
	n := g.GetEdgeByVector(x, y, vx, vy)
	return n.IsDirectional() && (n.IsReverse() == (vx < 0 || vy < 0))
}

func (g *Graph) IsEdgeDirectedBetweenCoords(x, y, tox, toy int) bool {
	vx, vy := tox-x, toy-y
	n := g.GetEdgeByVector(x, y, vx, vy)
	return n.IsDirectional() && (n.IsReverse() == (vx < 0 || vy < 0))
}

func (g *Graph) doCoordsHaveIngoingLinksOnly(x, y int) bool {
	for i := range cardinalDirections {
		vx, vy := unwrapCoords(cardinalDirections[i])
		if g.areCoordsInBounds(x+vx, y+vy) {
			if g.isEdgeByVectorDirectionalAndActive(x, y, vx, vy) {
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
	g.NodeAt(x, y).SetLinkByVector(vx, vy, true, true, reverse)
}

func (g *Graph) EnableDirectionalLinkBetweenCoords(from, to geometry.Coords) {
	vx, vy := from.VectorTo(to)
	g.EnableDirLinkByVector(from[0], from[1], vx, vy)
}

func (g *Graph) DisableDirectionalLinkBetweenCoords(from, to geometry.Coords) {
	vx, vy := from.VectorTo(to)
	g.setLinkByVector(from[0], from[1], vx, vy, false)
}
