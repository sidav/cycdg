package graph

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

func (g *Graph) AreCoordsLinked(c1, c2 Coords) bool {
	return g.AreCoordsLinkedXY(c1.X, c1.Y, c2.X, c2.Y)
}

func (g *Graph) AreCoordsLinkedXY(x1, y1, x2, y2 int) bool {
	if x1 < 0 || y1 < 0 {
		return false
	}
	if x2 < 0 || y2 < 0 {
		return false
	}
	return g.GetEdgeBetweenCoordsXY(x1, y1, x2, y2).IsActive()
}

func (g *Graph) IsEdgeByVectorActive(x, y, vx, vy int) bool {
	return g.AreCoordsLinkedXY(x, y, x+vx, y+vy)
}

func (g *Graph) CountEdgesAtXY(x, y int) int {
	count := 0
	for _, dir := range cardinalDirections {
		vx, vy := unwrapCoords(dir)
		otherx, othery := x+vx, y+vy
		if !g.AreCoordsInBounds(otherx, othery) {
			continue
		}
		if g.GetEdgeBetweenCoordsXY(otherx, othery, x, y).IsActive() {
			count++
		}
	}
	return count
}

func (g *Graph) GetEdgeByVector(x, y, vx, vy int) *Edge {
	return g.GetEdgeBetweenCoordsXY(x, y, x+vx, y+vy)
}

func (g *Graph) GetEdgeBetweenCoordsXY(fromx, fromy, x, y int) *Edge {
	vx, vy := x-fromx, y-fromy
	if vx*vy != 0 {
		debugPanic("Diagonal connection?.. %d,%d -> %d,%d", fromx, fromy, x, y)
	}
	if vx == 0 && vy == 0 {
		debugPanic("Zero vector?.. %d,%d -> %d,%d", fromx, fromy, x, y)
	}
	if vx == -1 {
		fromx--
		vx = 1
	}
	if vy == -1 {
		fromy--
		vy = 1
	}
	return g.NodeAtXY(fromx, fromy).GetEdgeByVector(vx, vy)
}

func (g *Graph) GetEdgeBetweenCoords(from, to Coords) *Edge {
	fx, fy := from.Unwrap()
	tx, ty := to.Unwrap()
	return g.GetEdgeBetweenCoordsXY(fx, fy, tx, ty)
}

func (g *Graph) setLinkByVector(x, y, vx, vy int, link bool) {
	if vx*vy != 0 {
		debugPanic("Diagonal connection?.. %d,%d -> %d,%d", x, y, vx, vy)
	}
	if vx == -1 {
		x--
		vx = 1
	}
	if vy == -1 {
		y--
		vy = 1
	}
	g.NodeAtXY(x, y).SetLinkByVector(vx, vy, link, false)
}
