package replacement

import (
	graph "cycdg/graph_replacement/grid_graph"
	"cycdg/graph_replacement/grid_graph/geometry"
	"cycdg/lib/random"
	"strings"
)

var rnd random.PRNG

type GraphReplacementApplier struct {
	graph                *graph.Graph
	MinCycles, MaxCycles int
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
	rnd = r
	graph.SetRandom(r)
	gra.graph = &graph.Graph{}
	gra.graph.InitWithConnectedNodes(width, height)
	gra.graph.ApplyRandomInitialRule()
	// gra.graph.ap
}

func (gra *GraphReplacementApplier) ApplyRandomRuleToTheGraph() {
	var rule *indifferentRule
	var applicableCoords [][]geometry.Coords
	try := 0
	for {
		rule = allReplacementRules[rnd.Rand(len(allReplacementRules))]
		if rule.AddsCycle && gra.graph.CyclesCount >= gra.MaxCycles {
			continue
		}
		applicableCoords = rule.FindAllApplicableCoordVariantsRecursively(gra.graph)
		if len(applicableCoords) > 0 {
			break
		}
		try++
		if try > 10000 {
			panic("No applicable coords even after 10000 tries!")
		}
	}
	gra.applyIndifferentRule(rule, applicableCoords)
}

func (gra *GraphReplacementApplier) applyIndifferentRule(rule *indifferentRule, applicableCoords [][]geometry.Coords) {
	crds := applicableCoords[rnd.Rand(len(applicableCoords))]
	rule.applyToCoords(gra.graph, crds...)
	if rule.AddsCycle {
		gra.graph.CyclesCount++
	}
	gra.graph.AppliedRulesCount++
	gra.graph.AppliedRules = append(gra.graph.AppliedRules, sprintf("%-10s at %v", rule.Name, crds))
	sane, errs := gra.graph.TestSanity()
	if !sane {
		panic(sprintf("Rule %s has caused the graph to have following problems:\n%v\nCoords: %v",
			rule.Name, strings.Join(errs, ";\n"), crds))
	}
}
