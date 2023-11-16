package graph

import "cycdg/grid_graph/graph_element"

func (g *Graph) addNodeTag(x, y int, tag graph_element.TagKind) {
	id := g.AppliedTags[tag]
	g.AppliedTags[tag] = id + 1
	g.nodes[x][y].AddTag(tag, id)
}

func (g *Graph) SwapTagsAtCoords(x1, y1, x2, y2 int) {
	n1 := g.NodeAt(x1, y1)
	n2 := g.NodeAt(x2, y2)
	n1.SwapTagsWith(n2)
}

func (g *Graph) addEdgeTagByVector(x, y, vx, vy int, tag graph_element.TagKind) {
	id := g.AppliedTags[tag]
	g.AppliedTags[tag] = id + 1
	g.GetEdgeByVector(x, y, vx, vy).AddTag(tag, id)
}
