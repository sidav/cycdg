package replacement

import . "cycdg/graph_replacement/grammar"

func (gra *GraphReplacementApplier) ApplyRandomInitialRule() {
	rule := &AllInitialRules[rnd.Rand(len(AllInitialRules))]
	if rule.IsApplicableForGraph(gra.graph) {
		gra.applyInitialRule(rule)
	} else {
		debugPanic("Initial rule %s failed!", rule.Name)
	}
}

func (gra *GraphReplacementApplier) applyInitialRule(rule *InitialRule) {
	x, y := rule.GetRandomApplicableCoordsForGraph(gra.graph)
	rule.ApplyOnGraphAt(gra.graph, x, y)
	appliedFeature := rule.Features[rnd.Rand(len(rule.Features))]
	appliedFeature.ApplyFeature(gra.graph)
	if rule.AddsCycle {
		gra.graph.CyclesCount++
	}
	gra.graph.AppliedRulesCount++
	gra.graph.AppliedRules = append(gra.graph.AppliedRules, sprintf("%-10s at %d,%d", rule.Name+"+"+appliedFeature.Name, x, y))
}
