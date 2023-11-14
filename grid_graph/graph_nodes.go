package graph

import (
	"cycdg/grid_graph/geometry"
	. "cycdg/grid_graph/graph_element"
)

func (g *Graph) resetNodeAndConnections(x, y int) {
	g.NodeAt(x, y).ResetActiveAndLinks()
	for _, d := range cardinalDirections {
		if g.areCoordsInBounds(x+d[0], y+d[1]) {
			e := g.GetEdgeByVector(x, y, d[0], d[1])
			e.Reset()
		}
	}
}

func (g *Graph) AreNodesBetweenCoordsEditable(x1, y1, x2, y2 int) bool {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if !g.areCoordsInBounds(x, y) {
				return false
			}
			if g.IsNodeFinalized(x, y) {
				return false
			}
		}
	}
	return true
}

func (g *Graph) enableNode(x, y int) {
	g.nodes[x][y].SetActive(true)
}

func (g *Graph) FinalizeNode(c geometry.Coords) {
	g.NodeAt(c.Unwrap()).Finalize()
}

func (g *Graph) IsNodeEditable(x, y int) bool {
	return !g.IsNodeFinalized(x, y)
}

func (g *Graph) IsNodeFinalized(x, y int) bool {
	return g.NodeAt(x, y).IsFinalized()
}

func (g *Graph) IsNodeActive(x, y int) bool {
	return g.NodeAt(x, y).IsActive()
}

func (g *Graph) NodeAt(x, y int) *Node {
	if !g.areCoordsInBounds(x, y) {
		return nil
	}
	return g.nodes[x][y]
}

func (g *Graph) GetRandomNodeByFunc(appropriate func(int, int) bool) *Node {
	return g.NodeAt(g.GetRandomCoordsByFunc(appropriate))
}

func (g *Graph) addNodeTag(x, y int, tag string) {
	g.nodes[x][y].AddTag(tag)
}

func (g *Graph) SwapTagsAtCoords(x1, y1, x2, y2 int) {
	n1 := g.NodeAt(x1, y1)
	t1 := n1.GetTags()
	n2 := g.NodeAt(x2, y2)
	t2 := n2.GetTags()
	n1.ResetTags()
	n1.SetAllTags(t2)
	n2.ResetTags()
	n2.SetAllTags(t1)
}

func (g *Graph) GetEnabledNodesCount() int {
	total := 0
	for x := range g.nodes {
		for y := range g.nodes[x] {
			if g.IsNodeActive(x, y) {
				total++
			}
		}
	}
	return total
}

func (g *Graph) GetFilledNodesPercentage() int {
	count := 0
	for x := range g.nodes {
		for y := range g.nodes[x] {
			if g.IsNodeActive(x, y) || g.IsNodeFinalized(x, y) {
				count++
			}
		}
	}
	w, h := g.GetSize()
	totalNodes := w * h
	return (100*count + totalNodes/2) / totalNodes
}

func (g *Graph) DoesNodeHaveAnyTags(x, y int) bool {
	return g.NodeAt(x, y).HasAnyTags()
}
