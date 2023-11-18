package graph

func (g *Graph) GetSize() (w, h int) {
	return len(g.nodes), len(g.nodes[0])
}

func (g *Graph) AreCoordsInBounds(x, y int) bool {
	return x >= 0 && x < len(g.nodes) && y >= 0 && y < len(g.nodes[0])
}

func (g *Graph) areCoordsOnBorder(x, y int) bool {
	return x == 0 || x == len(g.nodes)-1 || y == 0 || y == len(g.nodes[0])-1
}

func (g *Graph) areCoordsInCorner(x, y int) bool {
	return (x == 0 || x == len(g.nodes)-1) && (y == 0 || y == len(g.nodes[0])-1)
}

func (g *Graph) checkNearCoords(x, y int, check func(x, y int) bool) bool {
	for _, d := range cardinalDirections {
		n := g.NodeAt(x+d[0], y+d[1])
		if n != nil && check(x+d[0], y+d[1]) {
			return true
		}
	}
	return false
}
