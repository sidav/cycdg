package grammar

type ruleMetadata struct {
	AddsCycle                bool
	AddsTeleport             bool
	AdditionalWeight         int
	EnablesNodes             int
	EnablesNodesUnknown      bool // request full enabled nodes recalculation on apply
	FinalizesDisabledNodes   int
	UnfinalizesDisabledNodes int
}
