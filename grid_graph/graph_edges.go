package graph

import . "cycdg/grid_graph/graph_element"

func (g *Graph) AreCoordsInterlinked(x1, y1, x2, y2 int) bool {
	if x1 < 0 || y1 < 0 {
		return false
	}
	if x2 < 0 || y2 < 0 {
		return false
	}
	return g.GetEdgeBetweenCoords(x1, y1, x2, y2).IsActive()
}

func (g *Graph) IsEdgeByVectorActive(x, y, vx, vy int) bool {
	return g.AreCoordsInterlinked(x, y, x+vx, y+vy)
}

func (g *Graph) GetEdgeByVector(x, y, vx, vy int) *Edge {
	return g.GetEdgeBetweenCoords(x, y, x+vx, y+vy)
}

func (g *Graph) GetEdgeBetweenCoords(fromx, fromy, x, y int) *Edge {
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
	return g.NodeAt(fromx, fromy).GetEdgeByVector(vx, vy)
}

func (g *Graph) addEdgeTagFromCoordsByVector(x, y, vx, vy int, tag string) {
	g.GetEdgeByVector(x, y, vx, vy).AddTag(tag)
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
	g.NodeAt(x, y).SetLinkByVector(vx, vy, link, false, false)
}

func (g *Graph) setLinkBetweenCoords(fromx, fromy, x, y int, link bool) {
	vx, vy := x-fromx, y-fromy
	g.setLinkByVector(fromx, fromy, vx, vy, link)
}
