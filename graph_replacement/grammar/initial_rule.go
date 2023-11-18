package grammar

import . "cycdg/graph_replacement/grid_graph"

type InitialRule struct {
	Name string
	// metadata:
	AddsCycle bool

	// funcs:
	IsApplicableAt func(g *Graph, x, y, vx, vy int) bool
	ApplyOnGraphAt func(g *Graph, x, y, vx, vy int)
}

func (r *InitialRule) IsApplicableForGraph(g *Graph) bool {
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if r.IsApplicableAt(g, x, y, 0, 0) {
				return true
			}
		}
	}
	return false
}

func (r *InitialRule) GetRandomApplicableCoordsForGraph(g *Graph) (int, int, int, int) {
	w, h := g.GetSize()
	var candidates [][4]int
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if r.IsApplicableAt(g, x, y, 0, 0) {
				candidates = append(candidates, [4]int{x, y, 0, 0})
			}
		}
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind][0], candidates[ind][1], candidates[ind][2], candidates[ind][3]
}
