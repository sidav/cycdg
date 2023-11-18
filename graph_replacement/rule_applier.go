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

func (gra *GraphReplacementApplier) ApplyRandomInitialRule() {
	rule := &grammar.AllInitialRules[rnd.Rand(len(grammar.AllInitialRules))]
	if rule.IsApplicableForGraph(gra.graph) {
		gra.applyInitialRule(rule)
	} else {
		debugPanic("Initial rule %s failed!", rule.Name)
	}
}

func (gra *GraphReplacementApplier) applyInitialRule(rule *grammar.InitialRule) {
	x, y, vx, vy := rule.GetRandomApplicableCoordsForGraph(gra.graph)
	rule.ApplyOnGraphAt(gra.graph, x, y, vx, vy)
	if rule.AddsCycle {
		gra.graph.CyclesCount++
	}
	gra.graph.AppliedRulesCount++
	gra.graph.AppliedRules = append(gra.graph.AppliedRules, sprintf("%-10s at %d,%d vector %d,%d", rule.Name, x, y, vx, vy))
}
