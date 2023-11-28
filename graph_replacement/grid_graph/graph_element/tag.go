package graph_element

type TagKind uint8

const (
	// Node tags
	TagStart = iota
	TagGoal
	TagKey
	TagHalfkey   // for TagBilockedEdge
	TagMasterkey // for TagMasterLockedEdge, only one per map
	TagBoss
	TagTreasure
	TagHazard
	TagTrap
	TagTeleportBidirectional
	// Edge tags
	TagLockedEdge
	TagBilockedEdge     // requires two Halfkeys of same id
	TagMasterLockedEdge // requires the "master key"
	TagWindowEdge       // can be seen through, but not passable
	TagOneTimeEdge      // one-time passage
	TagOneWayEdge       // can be passed only in one direction (or maybe a door that opens from only one side?)
	TagSecretEdge
)

type Tag struct {
	Kind TagKind
	Id   int // should be unique in TagKind scope
}
