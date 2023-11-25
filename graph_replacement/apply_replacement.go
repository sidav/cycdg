package replacement

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grammar"
	"strings"
)

func (ra *GraphReplacementApplier) SelectRandomRuleToApply() *ReplacementRule {
	index := rnd.SelectRandomIndexFromWeighted(len(AllReplacementRules),
		func(i int) int {
			r := AllReplacementRules[i]
			if r.AddsCycle {
				if ra.MinCycles > ra.graph.CyclesCount {
					return 2 // ra.graph.AppliedRulesCount
				}
				if ra.MaxCycles <= ra.graph.CyclesCount {
					return 0
				}
			}
			return 1
		})
	return AllReplacementRules[index]
}

func (ra *GraphReplacementApplier) shouldFeatureBeAdded() bool {
	// TODO: rework :(
	if ra.DesiredFeatures <= ra.graph.AppliedFeaturesCount {
		return false
	}
	featuresPerc := (100*ra.graph.AppliedFeaturesCount + ra.DesiredFeatures/2) / ra.DesiredFeatures
	return rnd.Rand(120) > featuresPerc
}

func (ra *GraphReplacementApplier) ApplyRandomReplacementRuleToTheGraph() {
	var rule *ReplacementRule
	var applicableCoords [][]Coords
	try := 0
	for {
		rule = ra.SelectRandomRuleToApply()
		applicableCoords = rule.FindAllApplicableCoordVariantsRecursively(ra.graph)
		if len(applicableCoords) > 0 {
			break
		}
		try++
		if try > 10000 {
			panic("No applicable coords even after 10000 tries!")
		}
	}
	ra.applyReplacementRule(rule, applicableCoords)
}

func (ra *GraphReplacementApplier) applyReplacementRule(rule *ReplacementRule, applicableCoords [][]Coords) {
	crds := applicableCoords[rnd.Rand(len(applicableCoords))]

	// Set random optional feature to be added if needed
	selectedOptionalFeature := ra.SelectRandomOptionalFeatureToApply(rule)
	if selectedOptionalFeature != nil && selectedOptionalFeature.PrepareFeature != nil {
		selectedOptionalFeature.PrepareFeature(ra.graph, crds...)
	}

	rule.ApplyToGraph(ra.graph, crds...)

	if selectedOptionalFeature != nil && selectedOptionalFeature.ApplyFeature != nil {
		selectedOptionalFeature.ApplyFeature(ra.graph, crds...)
	}

	// update stats
	if rule.AddsCycle {
		ra.graph.CyclesCount++
	}
	ra.graph.AppliedRulesCount++
	ra.graph.AppliedRules = append(ra.graph.AppliedRules, ra.makeAppliedRuleString(rule, nil, selectedOptionalFeature, crds))

	// checking graph sanity (are there any bad graph patterns after the rule?)
	sane, errs := ra.graph.TestSanity()
	if !sane {
		panic(sprintf("Rule %s has caused the graph to have following problems:\n%v\nCoords: %v",
			rule.Name, strings.Join(errs, ";\n"), crds))
	}
}

func (ra *GraphReplacementApplier) SelectRandomOptionalFeatureToApply(rule *ReplacementRule) *FeatureAdder {
	var selectedOptionalFeature *FeatureAdder
	if ra.shouldFeatureBeAdded() && len(rule.OptionalFeatures) > 0 {
		selectedOptionalFeature = rule.OptionalFeatures[rnd.Rand(len(rule.OptionalFeatures))]
	}
	return selectedOptionalFeature
}

func (ra *GraphReplacementApplier) SelectRandomMandatoryFeatureToApply(rule *ReplacementRule) *FeatureAdder {
	var mandatoryFeature *FeatureAdder
	if len(rule.MandatoryFeatures) > 0 {
		mandatoryFeature = rule.MandatoryFeatures[rnd.Rand(len(rule.MandatoryFeatures))]
	}
	return mandatoryFeature
}

func (ra *GraphReplacementApplier) makeAppliedRuleString(
	rule *ReplacementRule, mandatory, optional *FeatureAdder, coords []Coords) string {
	ruleString := rule.Name
	if optional != nil {
		ruleString += "+" + optional.Name
	}
	return sprintf("%-15s at %v", ruleString, coords)
}
