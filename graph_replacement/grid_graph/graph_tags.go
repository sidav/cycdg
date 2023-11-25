package graph

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

func (g *Graph) AddNodeTag(x, y int, tag TagKind) {
	id := g.AppliedTags[tag]
	g.AppliedTags[tag] = id + 1
	g.nodes[x][y].AddTag(tag, id)
}

func (g *Graph) AddNodeTagByCoords(c Coords, tag TagKind) {
	x, y := c.Unwrap()
	g.AddNodeTag(x, y, tag)
}

func (g *Graph) AddNodeTagByCoordsPreserveLastId(c Coords, tag TagKind) {
	g.AppliedTags[tag]--
	x, y := c.Unwrap()
	g.AddNodeTag(x, y, tag)
}

func (g *Graph) AddEdgeTagByVector(x, y, vx, vy int, tag TagKind) {
	id := g.AppliedTags[tag]
	g.AppliedTags[tag] = id + 1
	g.GetEdgeByVector(x, y, vx, vy).AddTag(tag, id)
}

func (g *Graph) AddEdgeTagByCoords(c1, c2 Coords, tag TagKind) {
	id := g.AppliedTags[tag]
	g.AppliedTags[tag] = id + 1
	edge := g.GetEdgeBetweenCoords(c1, c2)
	edge.AddTag(tag, id)
}

func (g *Graph) AddEdgeTagByCoordsPreserveLastId(c1, c2 Coords, tag TagKind) {
	g.AppliedTags[tag]--
	g.AddEdgeTagByCoords(c1, c2, tag)
}

func (g *Graph) AddTagToAllActiveEdgesAtCoords(t TagKind, crds Coords) {
	x, y := crds.Unwrap()
	for _, dir := range cardinalDirections {
		if g.AreCoordsInBounds(x+dir[0], y+dir[1]) && g.isEdgeByVectorDirectionalAndActive(x, y, dir[0], dir[1]) {
			g.AddEdgeTagByVector(x, y, dir[0], dir[1], t)
		}
	}
}

func (g *Graph) SwapNodeTags(c1, c2 Coords) {
	n1 := g.NodeAt(c1.Unwrap())
	n2 := g.NodeAt(c2.Unwrap())
	n1.SwapTagsWith(n2)
}

func (g *Graph) DoesNodeHaveAnyTags(x, y int) bool {
	return g.NodeAt(x, y).HasAnyTags()
}

func (g *Graph) DoesNodeHaveTag(x, y int, t TagKind) bool {
	return g.NodeAt(x, y).HasTag(t)
}

func (g *Graph) DoesNodeByCoordsHaveTag(c Coords, t TagKind) bool {
	return g.NodeAt(c.Unwrap()).HasTag(t)
}

func (g *Graph) CountNodeTags(c Coords) int {
	x, y := c.Unwrap()
	return len(g.NodeAt(x, y).GetTags())
}

func (g *Graph) DoesEdgeHaveZeroTags(c1, c2 Coords) bool {
	edge := g.GetEdgeBetweenCoords(c1, c2)
	return len(edge.GetTags()) == 0
}

func (g *Graph) SwapEdgeTags(cfrom1, cfrom2, cto1, cto2 Coords) {
	edge1 := g.GetEdgeBetweenCoords(cfrom1, cfrom2)
	edge2 := g.GetEdgeBetweenCoords(cto1, cto2)
	edge1.SwapTagsWith(edge2)
}

// WARNING: COPIES IDS TOO!
func (g *Graph) CopyEdgeTagsPreservingIds(cfrom1, cfrom2, cto1, cto2 Coords) {
	edge1 := g.GetEdgeBetweenCoords(cfrom1, cfrom2)
	edge2 := g.GetEdgeBetweenCoords(cto1, cto2)
	edge1.CopyTagsFrom(edge2)
}
