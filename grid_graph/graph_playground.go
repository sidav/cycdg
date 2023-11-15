package graph

import (
	"cycdg/grid_graph/geometry"
	"strings"
)

func (g *Graph) AlterSomething() {
	if g.GetFilledNodesPercentage() == 0 {
		g.applyRandomInitialRule()
		return
	}
	g.applyRandomReplacementRule(2, 3)
}

func (g *Graph) applyRandomInitialRule() {
	rule := &initialRules[rnd.Rand(len(initialRules))]
	if rule.IsApplicableForGraph(g) {
		g.applyInitialRule(rule)
	} else {
		debugPanic("Initial rule %s failed!", rule.Name)
	}
}

func (g *Graph) applyRandomReplacementRule(minCycles, maxCycles int) {
	var rule *indifferentRule
	var applicableCoords [][]geometry.Coords
	for {
		rule = allReplacementRules[rnd.Rand(len(allReplacementRules))]
		if rule.AddsCycle && g.CyclesCount >= maxCycles {
			continue
		}
		applicableCoords = rule.FindAllApplicableCoordVariantsRecursively(g)
		if len(applicableCoords) > 0 {
			break
		}
	}
	g.applyIndifferentRule(rule, applicableCoords)
}

func (g *Graph) applyInitialRule(rule *ReplacementRule) {
	x, y, vx, vy := rule.GetRandomApplicableCoordsForGraph(g)
	rule.ApplyOnGraphAt(g, x, y, vx, vy)
	if rule.AddsCycle {
		g.CyclesCount++
	}
	g.AppliedRulesCount++
	g.AppliedRules = append(g.AppliedRules, sprintf("%-10s at %d,%d vector %d,%d", rule.Name, x, y, vx, vy))
}

func (g *Graph) applyIndifferentRule(rule *indifferentRule, applicableCoords [][]geometry.Coords) {
	crds := applicableCoords[rnd.Rand(len(applicableCoords))]
	rule.applyToCoords(g, crds...)
	if rule.AddsCycle {
		g.CyclesCount++
	}
	g.AppliedRulesCount++
	g.AppliedRules = append(g.AppliedRules, sprintf("%-10s at %v", rule.Name, crds))
	sane, errs := g.TestSanity()
	if !sane {
		panic(sprintf("Rule %s has caused the graph to have following problems:\n%v\nCoords: %v",
			rule.Name, strings.Join(errs, ";\n"), crds))
	}
}
