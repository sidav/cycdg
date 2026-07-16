package grammar

type Grammar interface {
	GetAllInitialRules() []*InitialRule
	GetAllReplacementRulesForStep(step int) []*ReplacementRule
}
