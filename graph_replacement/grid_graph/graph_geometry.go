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

func (g *Graph) GetRandomCoordsByFunc(appropriate func(int, int) bool) (int, int) {
	var candidates [][2]int
	for x := range g.nodes {
		for y := range g.nodes[x] {
			if appropriate(x, y) {
				candidates = append(candidates, [2]int{x, y})
			}
		}
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind][0], candidates[ind][1]
}

func (g *Graph) GetRandomAdjacentCoordsByFunc(fx, fy int, appropriate func(int, int) bool) (int, int) {
	x, y := g.GetRandomCoordsByFunc(
		func(x, y int) bool {
			return areCoordsAdjacent(x, y, fx, fy) && appropriate(x, y)
		},
	)
	return x, y
}
