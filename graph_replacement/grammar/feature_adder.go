package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

// Features are anything that add something tag-related to the graph. Key/door pairs, bosses etc.
type FeatureAdder struct {
	Name string
	// applied before the rule itself, can be nil
	PrepareFeature func(g *Graph, crds ...Coords)
	// applied after the rule itself, can be nil too (but why?)
	ApplyFeature func(g *Graph, crds ...Coords)
}

// In case of keys: it should be called BEFORE the nodes/edges with lock are added!
// It may end up being behind the lock otherwise.
// So it must be used with rules which add locked edge, not change an existing to locked
func addTagAtRandomActiveNode(g *Graph, tag TagKind) {
	// crd := getRandomGraphCoordsByFunc(g, func(x, y int) bool {
	// 	return g.IsNodeActive(x, y)
	// })
	// prefer non-tagged nodes
	crd := getRandomGraphCoordsByScore(g, func(x, y int) int {
		if !g.IsNodeActive(x, y) {
			return 0
		}
		if g.DoesNodeHaveAnyTags(x, y) {
			return 1
		}
		return 5
	})
	g.AddNodeTagByCoords(crd, tag)
}
