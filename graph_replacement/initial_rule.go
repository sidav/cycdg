package replacement

import . "cycdg/graph_replacement/grid_graph"

type InitialRule struct {
	Name string
	// params restrictions:
	vectorable             bool
	cardinalVectorsAllowed bool
	diagonalVectorsAllowed bool

	// metadata:
	AddsCycle bool

	// funcs:
	IsApplicableAt func(g *Graph, x, y, vx, vy int) bool
	ApplyOnGraphAt func(g *Graph, x, y, vx, vy int)
}

func (r *InitialRule) IsApplicableForGraph(g *Graph) bool {
	w, h := g.GetSize()
	if !r.vectorable {
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if r.IsApplicableAt(g, x, y, 0, 0) {
					return true
				}
			}
		}
	} else {
		if !(r.cardinalVectorsAllowed || r.diagonalVectorsAllowed) {
			debugPanic("Rule %s is vectorable, but has no directions allowed!", r.Name)
		}
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if r.cardinalVectorsAllowed {
					for _, dir := range cardinalDirections {
						if r.IsApplicableAt(g, x, y, dir[0], dir[1]) {
							return true
						}
					}
				}
				if r.diagonalVectorsAllowed {
					for _, dir := range diagonalDirections {
						if r.IsApplicableAt(g, x, y, dir[0], dir[1]) {
							return true
						}
					}
				}
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
			if !r.vectorable {
				if r.IsApplicableAt(g, x, y, 0, 0) {
					candidates = append(candidates, [4]int{x, y, 0, 0})
				}
			}
			if r.cardinalVectorsAllowed {
				for _, dir := range cardinalDirections {
					if r.IsApplicableAt(g, x, y, dir[0], dir[1]) {
						candidates = append(candidates, [4]int{x, y, dir[0], dir[1]})
					}
				}
			}
			if r.diagonalVectorsAllowed {
				for _, dir := range diagonalDirections {
					if r.IsApplicableAt(g, x, y, dir[0], dir[1]) {
						candidates = append(candidates, [4]int{x, y, dir[0], dir[1]})
					}
				}
			}
		}
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind][0], candidates[ind][1], candidates[ind][2], candidates[ind][3]
}
