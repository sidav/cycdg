package graph

import (
	. "cycdg/graph_replacement/grid_graph/graph_element"
	"cycdg/lib/random"
)

var (
	rnd random.PRNG
)

// graph with nodes placed at 2D grid
type Graph struct {
	nodes             [][]*Node
	CyclesCount       int
	AppliedRulesCount int
	AppliedRules      []string
	AppliedTags       map[TagKind]int
}

func (g *Graph) Init(r random.PRNG, w, h int) {
	rnd = r
	g.AppliedRules = nil
	g.AppliedRulesCount = 0
	g.AppliedTags = make(map[TagKind]int)
	g.CyclesCount = 0
	g.nodes = make([][]*Node, w)
	for x := range g.nodes {
		g.nodes[x] = make([]*Node, h)
	}

	for x := range g.nodes {
		for y := range g.nodes[x] {
			g.nodes[x][y] = &Node{}
			g.nodes[x][y].Init()
		}
	}

	// removing links for border nodes
	for x := range g.nodes {
		// g.nodes[x][0].setLinkByVector(false, 0, -1)
		g.nodes[x][h-1].SetLinkByVector(0, 1, false, false, false)
	}
	for y := range g.nodes[0] {
		// g.nodes[0][y].setLinkByVector(false, -1, 0)
		g.nodes[w-1][y].SetLinkByVector(1, 0, false, false, false)
	}
}