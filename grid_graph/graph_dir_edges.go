package graph

func (g *Graph) IsEdgeByVectorDirectional(x, y, vx, vy int) bool {
	n := g.GetEdgeByVector(x, y, vx, vy)
	return n.IsDirectional()
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

func (g *Graph) enableDirLinkByVector(x, y, vx, vy int) {
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

func (g *Graph) enableDirectionalLinkBetweenCoords(fromx, fromy, x, y int) {
	vx, vy := x-fromx, y-fromy
	g.enableDirLinkByVector(fromx, fromy, vx, vy)
}
