package replacement

import . "cycdg/graph_replacement/grammar"

func (gra *GraphReplacementApplier) ApplyRandomInitialRule() {
	totalRules := len(gra.grammar.GetAllInitialRules())
	rule := &(gra.grammar.GetAllInitialRules()[rnd.Rand(totalRules)])
	if rule.IsApplicableForGraph(gra.graph) {
		gra.applyInitialRule(rule)
	} else {
		gra.debugPanic("Initial rule %s failed!", rule.Name)
	}
	gra.EnabledNodesCount = gra.graph.GetEnabledNodesCount()
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
