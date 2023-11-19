package replacement

import (
	"cycdg/graph_replacement/grammar"
	graph "cycdg/graph_replacement/grid_graph"
	"cycdg/lib/random"
)

var rnd random.PRNG

type GraphReplacementApplier struct {
	graph                *graph.Graph
	MinCycles, MaxCycles int
	DesiredFeatures      int
}

func (gra *GraphReplacementApplier) GetGraph() *graph.Graph {
	return gra.graph
}

func (gra *GraphReplacementApplier) Init(r random.PRNG, width, height int) {
	if gra.MinCycles == 0 {
		gra.MinCycles = 1
	}
	if gra.MaxCycles == 0 {
		gra.MaxCycles = 4
	}
	if gra.DesiredFeatures == 0 {
		gra.DesiredFeatures = 5
	}
	rnd = r
	grammar.SetRandom(rnd)
	gra.graph = &graph.Graph{}
	gra.graph.Init(width, height)
	gra.ApplyRandomInitialRule()
	// gra.graph.ap
}
