package grammar

func CreateExampleGrammarObject() Grammar {
	eg := &exampleGrammar{}
	eg.initialRules = allInitialRules
	eg.replacementRules = make([]*ReplacementRule, 0)
	eg.replacementRules = append(eg.replacementRules, allCycleRules...)
	eg.replacementRules = append(eg.replacementRules, allGrowingRules...)
	eg.replacementRules = append(eg.replacementRules, allNonGrowingRules...)
	return eg
}

type exampleGrammar struct {
	initialRules     []*InitialRule
	replacementRules []*ReplacementRule
}

func (eg *exampleGrammar) GetAllInitialRules() []*InitialRule {
	return eg.initialRules
}

func (eg *exampleGrammar) GetAllReplacementRulesForStep(step int) []*ReplacementRule {
	var applicableRules []*ReplacementRule
	for _, v := range eg.replacementRules {
		if v.isApplicableAtStep(step) {
			applicableRules = append(applicableRules, v)
		}
	}
	return applicableRules
}
