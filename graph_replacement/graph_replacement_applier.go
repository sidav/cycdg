package replacement

import (
	"cycdg/graph_replacement/grammar"
	graph "cycdg/graph_replacement/grid_graph"
	"cycdg/lib/random"
)

var rnd random.PRNG

type GraphReplacementApplier struct {
	graph                                    *graph.Graph
	MinCycles, MaxCycles                     int
	DesiredFeatures                          int
	MinFilledPercentage, MaxFilledPercentage int
	desiredFillPercentage                    int // The real resulting value will most likely be bigger than this
	MaxTeleports                             int

	// Some meta-info on applied rules
	CyclesCount          int
	AppliedRulesCount    int
	AppliedFeaturesCount int
	TeleportsCount       int
	AppliedRules         []*AppliedRuleInfo
}

func (gra *GraphReplacementApplier) GetGraph() *graph.Graph {
	return gra.graph
}

func (gra *GraphReplacementApplier) Init(r random.PRNG, width, height int) {
	rnd = r
	grammar.SetRandom(rnd)

	if width < 4 || height < 4 {
		debugPanic("Minimum allowed size violation: at least 4x4 is allowed.")
	}
	if gra.MinCycles == 0 {
		gra.MinCycles = 1
	}
	if gra.MaxCycles == 0 {
		gra.MaxCycles = 4
	}
	if gra.MaxCycles < gra.MinCycles {
		gra.MaxCycles = gra.MinCycles
	}
	if gra.DesiredFeatures == 0 {
		gra.DesiredFeatures = 5
	}
	gra.desiredFillPercentage = rnd.RandInRange(gra.MinFilledPercentage, gra.MaxFilledPercentage)

	gra.AppliedRules = nil
	gra.AppliedRulesCount = 0
	gra.AppliedFeaturesCount = 0
	gra.TeleportsCount = 0
	gra.CyclesCount = 0

	gra.graph = &graph.Graph{}
	gra.graph.Init(width, height)

	gra.ApplyRandomInitialRule()
}

func (gra *GraphReplacementApplier) Reset() {
	gra.desiredFillPercentage = rnd.RandInRange(gra.MinFilledPercentage, gra.MaxFilledPercentage)
	gra.AppliedRules = nil
	gra.AppliedRulesCount = 0
	gra.AppliedFeaturesCount = 0
	gra.TeleportsCount = 0
	gra.CyclesCount = 0

	width, height := gra.graph.GetSize()
	gra.graph = &graph.Graph{}
	gra.graph.Init(width, height)

	gra.ApplyRandomInitialRule()
}

func (gra *GraphReplacementApplier) FilledEnough() bool {
	if gra.desiredFillPercentage == 0 {
		debugPanic("Zero DesiredFillPercentage!")
	}
	return gra.graph.GetFilledNodesPercentage() >= gra.desiredFillPercentage
}
