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
	rule.ApplyToGraph(ra.graph, crds...)

	if rule.AddsCycle {
		ra.graph.CyclesCount++
	}
	ra.graph.AppliedRulesCount++
	ra.graph.AppliedRules = append(ra.graph.AppliedRules, sprintf("%-10s at %v", rule.Name, crds))
	sane, errs := ra.graph.TestSanity()
	if !sane {
		panic(sprintf("Rule %s has caused the graph to have following problems:\n%v\nCoords: %v",
			rule.Name, strings.Join(errs, ";\n"), crds))
	}
}
