package grammar

func CreateExampleGrammarObject() Grammar {
	eg := &exampleGrammar{}
	eg.initialRules = allInitialRules
	eg.replacementRules = make([]ReplacementRule, 0)
	eg.replacementRules = append(eg.replacementRules, allCycleRules...)
	eg.replacementRules = append(eg.replacementRules, allGrowingRules...)
	eg.replacementRules = append(eg.replacementRules, allNonGrowingRules...)
	return eg
}

type exampleGrammar struct {
	initialRules     []InitialRule
	replacementRules []ReplacementRule
}

func (eg *exampleGrammar) GetAllInitialRules() []InitialRule {
	return eg.initialRules
}

func (eg *exampleGrammar) GetAllReplacementRules() []ReplacementRule {
	return eg.replacementRules
}

// func (eg *exampleGrammar) GetAllCycleRules() []ReplacementRule {
// 	return allCycleRules
// }

// func (eg *exampleGrammar) GetAllGrowingRules() []ReplacementRule {
// 	return allGrowingRules
// }

// func (eg *exampleGrammar) GetAllNongrowingRules() []ReplacementRule {
// 	return allNonGrowingRules
// }
