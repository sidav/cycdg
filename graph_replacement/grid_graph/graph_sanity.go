package graph

func (g *Graph) TestSanity() (sane bool, problems []string) {
	sane = true
	// test if any disabled node has tags
	for x := range g.nodes {
		for y := range g.nodes[x] {
			n := g.NodeAt(x, y)
			if !n.IsActive() && n.HasAnyTags() {
				sane = false
				problems = append(problems, sprintf("Inactive node at %d,%d has tags", x, y))
			}
		}
	}
	// test if any disabled node has active links
	for x := range g.nodes {
		for y := range g.nodes[x] {
			n := g.NodeAt(x, y)
			if !n.IsActive() {
				for _, d := range cardinalDirections {
					if g.IsEdgeByVectorActive(x, y, d[0], d[1]) {
						sane = false
						problems = append(problems, sprintf("Inactive node at %d,%d has active link vector %v", x, y, d))
					}
				}
			}
		}
	}

	// test if any enabled node has no links
	for x := range g.nodes {
		for y := range g.nodes[x] {
			n := g.NodeAt(x, y)
			if n.IsActive() {
				if g.CountEdgesAt(x, y) == 0 {
					sane = false
					problems = append(problems, sprintf("Active node at %d,%d has no active links!", x, y))
				}
			}
		}
	}

	return sane, problems
}
