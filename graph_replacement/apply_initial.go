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
	appliedFeature := rule.MandatoryFeatures[rnd.Rand(len(rule.MandatoryFeatures))]
	appliedFeature.ApplyFeature(gra.graph)
	if rule.AddsCycle {
		gra.CyclesCount++
	}
	gra.AppliedRulesCount++
	gra.AppliedRules = append(gra.AppliedRules, newAppliedRuleInfoInitial(rule, appliedFeature, x, y))
}
