package grammar

import . "cycdg/graph_replacement/grid_graph"

type InitialRule struct {
	Name string
	// metadata:
	AddsCycle bool

	// funcs:
	IsApplicableAt func(g *Graph, x, y int) bool
	ApplyOnGraphAt func(g *Graph, x, y int)
	Features       []*FeatureAdder // They're mandatory to apply!
}

func (r *InitialRule) IsApplicableForGraph(g *Graph) bool {
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if r.IsApplicableAt(g, x, y) {
				return true
			}
		}
	}
	return false
}

func (r *InitialRule) GetRandomApplicableCoordsForGraph(g *Graph) (int, int) {
	w, h := g.GetSize()
	var candidates [][2]int
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if r.IsApplicableAt(g, x, y) {
				candidates = append(candidates, [2]int{x, y})
			}
		}
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind][0], candidates[ind][1]
}
