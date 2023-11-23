package graph_element

type TagKind uint8

const (
	// Node tags
	TagStart = iota
	TagGoal
	TagKeyForEdge
	TagBoss
	TagTreasure
	TagHazard
	TagTrap
	TagTeleportBidirectional
	// Edge tags
	TagLockedEdge
	TagBilockedEdge // requires two keys of same id
	TagWindowEdge   // can be seen through, but not passable
	TagSecretEdge
)

type Tag struct {
	Kind TagKind
	Id   int // should be unique in TagKind scope
}
