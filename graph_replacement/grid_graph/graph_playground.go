package graph

func (g *Graph) AlterSomething() {
	if g.GetFilledNodesPercentage() == 0 {
		g.ApplyRandomInitialRule()
		return
	}
}

func (g *Graph) ApplyRandomInitialRule() {
	rule := &initialRules[rnd.Rand(len(initialRules))]
	if rule.IsApplicableForGraph(g) {
		g.applyInitialRule(rule)
	} else {
		debugPanic("Initial rule %s failed!", rule.Name)
	}
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
