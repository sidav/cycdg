package graph

import (
	"cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

func (g *Graph) ResetNodeAndConnections(c geometry.Coords) {
	x, y := c.Unwrap()
	g.NodeAt(x, y).ResetActiveAndLinks()
	for _, d := range cardinalDirections {
		if g.AreCoordsInBounds(x+d[0], y+d[1]) {
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
			if !g.AreCoordsInBounds(x, y) {
				return false
			}
			if g.IsNodeFinalized(x, y) {
				return false
			}
		}
	}
	return true
}

func (g *Graph) EnableNode(x, y int) {
	g.nodes[x][y].SetActive(true)
}

func (g *Graph) EnableNodeByVector(x, y, vx, vy int) {
	g.nodes[x+vx][y+vy].SetActive(true)
}

func (g *Graph) EnableNodeByCoords(c geometry.Coords) {
	g.EnableNode(c.Unwrap())
}

func (g *Graph) FinalizeNode(c geometry.Coords) {
	g.NodeAt(c.Unwrap()).Finalize()
}

// Used for workarounds... Use cautiously
func (g *Graph) UnsafeUnfinalizeNode(c geometry.Coords) {
	g.NodeAt(c.Unwrap()).UnsafeUnfinalize()
}

func (g *Graph) IsNodeEditable(x, y int) bool {
	return !g.IsNodeFinalized(x, y)
}

func (g *Graph) IsNodeFinalized(x, y int) bool {
	return g.NodeAt(x, y).IsFinalized()
}

func (g *Graph) HasNoFinalizedNodesNearXY(x, y int, allowDiagonal bool) bool {
	if !allowDiagonal {
		debugPanic("Not implemented!")
	}
	for vx := -1; vx <= 1; vx++ {
		for vy := -1; vy <= 1; vy++ {
			if vx == vy && vx == 0 {
				continue
			}
			nodeHere := g.NodeAt(x+vx, y+vy)
			if nodeHere != nil && nodeHere.IsFinalized() {
				return false
			}
		}
	}
	return true
}

func (g *Graph) IsNodeActive(x, y int) bool {
	return g.NodeAt(x, y).IsActive()
}

func (g *Graph) NodeAt(x, y int) *Node {
	if !g.AreCoordsInBounds(x, y) {
		return nil
	}
	return g.nodes[x][y]
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

func (g *Graph) GetEnabledNodesPercentage() int {
	return getIntPercentage(g.GetEnabledNodesCount(), g.GetTotalNodesCount())
}

// func (g *Graph) GetFilledNodesCount() int {
// 	count := 0
// 	for x := range g.nodes {
// 		for y := range g.nodes[x] {
// 			// TODO: remove g.IsNodeFinalized(x, y) from here
// 			if g.IsNodeActive(x, y) || g.IsNodeFinalized(x, y) {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
// func (g *Graph) GetFilledNodesPercentage() int {
// 	return getIntPercentage(g.GetFilledNodesCount(), g.GetTotalNodesCount())
// }

func (g *Graph) GetFinalizedEmptyNodesCount() int {
	emptyFinsCount := 0
	for x := range g.nodes {
		for y := range g.nodes[x] {
			if !g.IsNodeActive(x, y) && g.IsNodeFinalized(x, y) {
				emptyFinsCount++
			}
		}
	}
	return emptyFinsCount
}

func (g *Graph) GetFinalizedEmptyNodesPercentage() int {
	return getIntPercentage(g.GetFinalizedEmptyNodesCount(), g.GetTotalNodesCount())
}

func (g *Graph) CountEmptyEditableNodesNearEnabledOnes() int {
	count := 0
	for x := 0; x < len(g.nodes); x++ {
		for y := 0; y < len(g.nodes[x]); y++ {
			currNode := g.NodeAt(x, y)
			if currNode != nil && !currNode.IsActive() && !currNode.IsFinalized() {
				// check if any neighbour is active
				for i := range cardinalDirections {
					neighbour := g.NodeAt(x+cardinalDirections[i][0], y+cardinalDirections[i][1])
					if neighbour != nil && neighbour.IsActive() {
						count++
						break
					}
				}
			}
		}
	}
	return count
}
