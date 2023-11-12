package graph

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
		g.applyRule(rule)
	} else {
		debugPanic("Initial rule %s failed!", rule.Name)
	}
}

func (g *Graph) applyRandomReplacementRule(minCycles, maxCycles int) {
	index := rnd.SelectRandomIndexFromWeighted(len(atomaryRules), func(i int) int {
		if atomaryRules[i].IsApplicableForGraph(g) {
			if atomaryRules[i].AddsCycle {
				if g.CyclesCount >= maxCycles {
					return 0
				} else if g.CyclesCount < minCycles {
					return 2
				}
			}
			return 1
		}
		return 0
	})
	rule := &atomaryRules[index]
	if rule.IsApplicableForGraph(g) {
		g.applyRule(rule)
	}
	sane, errs := g.TestSanity()
	if !sane {
		panic(sprintf("Rule %s has caused the graph to have following problems: %v", rule.Name, errs))
	}
}

func (g *Graph) applyRule(rule *ReplacementRule) {
	x, y, vx, vy := rule.GetRandomApplicableCoordsForGraph(g)
	rule.ApplyOnGraphAt(g, x, y, vx, vy)
	if rule.AddsCycle {
		g.CyclesCount++
	}
	g.AppliedRulesCount++
	g.AppliedRules = append(g.AppliedRules, sprintf("%-10s at %d,%d vector %d,%d", rule.Name, x, y, vx, vy))
}
