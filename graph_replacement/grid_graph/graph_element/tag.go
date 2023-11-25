package graph_element

type TagKind uint8

const (
	// Node tags
	TagStart = iota
	TagGoal
	TagKeyForEdge
	TagHalfkey // for TagBilockedEdge
	TagBoss
	TagTreasure
	TagHazard
	TagTrap
	TagTeleportBidirectional
	// Edge tags
	TagLockedEdge
	TagBilockedEdge // requires two Halfkeys of same id
	TagWindowEdge   // can be seen through, but not passable
	TagOnetimeEdge  // one-time passage
	TagSecretEdge
)

type Tag struct {
	Kind TagKind
	Id   int // should be unique in TagKind scope
}
