package replacement

import (
	"cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grammar"
	graph "cycdg/graph_replacement/grid_graph"
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
	SetRandom(rnd)
	gra.graph = &graph.Graph{}
	gra.graph.Init(rnd, width, height)
	gra.ApplyRandomInitialRule()
	// gra.graph.ap
}

func (gra *GraphReplacementApplier) ApplyRandomReplacementRuleToTheGraph() {
	var rule *ReplacementRule
	var applicableCoords [][]geometry.Coords
	try := 0
	for {
		rule = AllReplacementRules[rnd.Rand(len(AllReplacementRules))]
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
	gra.applyReplacementRule(rule, applicableCoords)
}

func (gra *GraphReplacementApplier) applyReplacementRule(rule *ReplacementRule, applicableCoords [][]geometry.Coords) {
	crds := applicableCoords[rnd.Rand(len(applicableCoords))]
	rule.ApplyToGraph(gra.graph, crds...)
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

func (gra *GraphReplacementApplier) ApplyRandomInitialRule() {
	rule := &AllInitialRules[rnd.Rand(len(AllInitialRules))]
	if rule.IsApplicableForGraph(gra.graph) {
		gra.applyInitialRule(rule)
	} else {
		debugPanic("Initial rule %s failed!", rule.Name)
	}
}

func (gra *GraphReplacementApplier) applyInitialRule(rule *InitialRule) {
	x, y, vx, vy := rule.GetRandomApplicableCoordsForGraph(gra.graph)
	rule.ApplyOnGraphAt(gra.graph, x, y, vx, vy)
	if rule.AddsCycle {
		gra.graph.CyclesCount++
	}
	gra.graph.AppliedRulesCount++
	gra.graph.AppliedRules = append(gra.graph.AppliedRules, sprintf("%-10s at %d,%d vector %d,%d", rule.Name, x, y, vx, vy))
}
