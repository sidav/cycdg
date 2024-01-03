package replacement

import (
	"cycdg/graph_replacement/grammar"
	graph "cycdg/graph_replacement/grid_graph"
	"cycdg/lib/random"
	"fmt"
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
	CyclesCount                 int
	AppliedRulesCount           int
	AppliedFeaturesCount        int
	TeleportsCount              int
	EnabledNodesCount           int
	FinalizedDisabledNodesCount int
	AppliedRules                []*AppliedRuleInfo
}

func (gra *GraphReplacementApplier) GetGraph() *graph.Graph {
	return gra.graph
}

func (gra *GraphReplacementApplier) Init(r random.PRNG, width, height int) {
	rnd = r
	grammar.SetRandom(rnd)

	if width < 4 || height < 4 {
		gra.debugPanic("Minimum allowed size violation: at least 4x4 is allowed.")
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

	gra.graph = &graph.Graph{}
	gra.graph.Init(width, height)

	gra.Reset()
}

func (gra *GraphReplacementApplier) Reset() {
	gra.desiredFillPercentage = rnd.RandInRange(gra.MinFilledPercentage, gra.MaxFilledPercentage)
	gra.AppliedRules = nil
	gra.AppliedRulesCount = 0
	gra.AppliedFeaturesCount = 0
	gra.TeleportsCount = 0
	gra.CyclesCount = 0
	gra.EnabledNodesCount = 0
	gra.FinalizedDisabledNodesCount = 0

	width, height := gra.graph.GetSize()
	gra.graph = &graph.Graph{}
	gra.graph.Init(width, height)

	gra.ApplyRandomInitialRule()
}

func (gra *GraphReplacementApplier) FilledEnough() bool {
	if gra.desiredFillPercentage == 0 {
		gra.debugPanic("Zero DesiredFillPercentage!")
	}
	if (gra.EnabledNodesCount + gra.FinalizedDisabledNodesCount) != gra.graph.GetFilledNodesCount() {
		gra.debugPanic("Debug error: gra.EnabledNodesCount (%d) != gra.graph.GetFilledNodesCount() (%d)",
			(gra.EnabledNodesCount + gra.FinalizedDisabledNodesCount), gra.graph.GetFilledNodesCount())
	}
	return getIntPercentage(gra.EnabledNodesCount+gra.FinalizedDisabledNodesCount, gra.graph.GetTotalNodesCount()) >= gra.desiredFillPercentage
}

func (gra *GraphReplacementApplier) StringifyGenerationMetadata() string {
	return fmt.Sprintf("%d rules %d cycles %d forced-empty, fill %d%%", gra.AppliedRulesCount, gra.CyclesCount,
		gra.FinalizedDisabledNodesCount, gra.graph.GetFilledNodesPercentage())
}

func (gra *GraphReplacementApplier) debugPanic(msg string, args ...interface{}) {
	fmt.Println()
	message := fmt.Sprintf(msg+"\n", args...)
	if gra.graph != nil {
		message += " Applied rules:\n"
		for i, rul := range gra.AppliedRules {
			message += fmt.Sprintf(" %-2d: %s\n", i, rul.StringifyRule())
		}
		message += fmt.Sprintf("Fill percentage: %d\n", gra.graph.GetFilledNodesPercentage())
		message += fmt.Sprintf("Empty-fin percentage: %d", gra.graph.GetFinalizedEmptyNodesPercentage())
	}
	panic(message)
}
