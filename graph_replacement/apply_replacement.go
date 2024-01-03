package replacement

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grammar"
	"strings"
)

const baseRuleWeight = 10

func (ra *GraphReplacementApplier) SelectRandomRuleToApply() *ReplacementRule {
	index := rnd.SelectRandomIndexFromWeighted(len(AllReplacementRules),
		func(i int) int {
			r := AllReplacementRules[i]
			if r.Metadata.AddsCycle {
				if ra.MinCycles > ra.CyclesCount {
					return 2 * baseRuleWeight // ra.graph.AppliedRulesCount
				}
				if ra.MaxCycles <= ra.CyclesCount {
					return 0
				}
			}
			if r.Metadata.AddsTeleport && ra.TeleportsCount >= ra.MaxTeleports {
				return 0
			}
			return r.Metadata.AdditionalWeight + baseRuleWeight
		})
	return AllReplacementRules[index]
}

func (ra *GraphReplacementApplier) shouldFeatureBeAdded() bool {
	// TODO: rework :(
	if ra.DesiredFeatures <= ra.AppliedFeaturesCount {
		return false
	}
	featuresPerc := (100*ra.AppliedFeaturesCount + ra.DesiredFeatures/2) / ra.DesiredFeatures
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

	selectedMandatoryFeature := ra.SelectRandomMandatoryFeatureToApply(rule)
	if selectedMandatoryFeature != nil && selectedMandatoryFeature.PrepareFeature != nil {
		selectedMandatoryFeature.PrepareFeature(ra.graph, crds...)
	}

	// Set random optional feature to be added if needed
	selectedOptionalFeature := ra.SelectRandomOptionalFeatureToApply(rule)
	if selectedOptionalFeature != nil && selectedOptionalFeature.PrepareFeature != nil {
		selectedOptionalFeature.PrepareFeature(ra.graph, crds...)
	}

	rule.ApplyToGraph(ra.graph, crds...)

	if selectedMandatoryFeature != nil && selectedMandatoryFeature.ApplyFeature != nil {
		selectedMandatoryFeature.ApplyFeature(ra.graph, crds...)
	}
	if selectedOptionalFeature != nil {
		ra.AppliedFeaturesCount++
		if selectedOptionalFeature.ApplyFeature != nil {
			selectedOptionalFeature.ApplyFeature(ra.graph, crds...)
		}
	}

	// update stats
	if rule.Metadata.AddsCycle {
		ra.CyclesCount++
	}
	if rule.Metadata.AddsTeleport {
		ra.TeleportsCount++
	}
	ra.AppliedRulesCount++
	ra.AppliedRules = append(ra.AppliedRules, newAppliedRuleInfo(
		rule, selectedMandatoryFeature, selectedOptionalFeature, crds))

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
		index := rnd.SelectRandomIndexFromWeighted(len(rule.OptionalFeatures), func(x int) int {
			return baseRuleWeight + rule.OptionalFeatures[x].AdditionalWeight
		})
		selectedOptionalFeature = rule.OptionalFeatures[index]
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
