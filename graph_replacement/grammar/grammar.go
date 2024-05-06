package grammar

type Grammar interface {
	GetAllInitialRules() []*InitialRule
	GetAllReplacementRules() []*ReplacementRule
}
