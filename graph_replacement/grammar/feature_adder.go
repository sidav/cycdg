package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

// Features are anything that add something tag-related to the graph. Key/door pairs, bosses etc.
type FeatureAdder struct {
	Name             string
	AdditionalWeight int
	// applied before the rule itself, can be nil
	PrepareFeature func(g *Graph, crds ...Coords)
	// applied after the rule itself, can be nil too (but why?)
	ApplyFeature func(g *Graph, crds ...Coords)
}

// In case of keys: it should be called BEFORE the nodes/edges with lock are added!
// It may end up being behind the lock otherwise.
// So it must be used with rules which add locked edge, not change an existing to locked
func addTagAtRandomActiveNode(g *Graph, tag TagKind) {
	crd := getRandomGraphCoordsByScore(g, func(x, y int) int {
		crd := NewCoords(x, y)
		if !g.IsNodeActive(crd) {
			return 0
		}
		if g.DoesNodeHaveAnyTags(crd) {
			return 1
		}
		return 100 // prefer non-tagged nodes
	})
	g.AddNodeTagByCoords(crd, tag)
}
