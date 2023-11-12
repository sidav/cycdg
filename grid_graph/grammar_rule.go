package graph

type ReplacementRule struct {
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

func (r *ReplacementRule) IsApplicableForGraph(g *Graph) bool {
	if !r.vectorable {
		for x := range g.nodes {
			for y := range g.nodes[0] {
				if r.IsApplicableAt(g, x, y, 0, 0) {
					return true
				}
			}
		}
	} else {
		if !(r.cardinalVectorsAllowed || r.diagonalVectorsAllowed) {
			debugPanic("Rule %s is vectorable, but has no directions allowed!", r.Name)
		}
		for x := range g.nodes {
			for y := range g.nodes[0] {
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

func (r *ReplacementRule) GetRandomApplicableCoordsForGraph(g *Graph) (int, int, int, int) {
	var candidates [][4]int
	for x := range g.nodes {
		for y := range g.nodes[0] {
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
