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
	// Edge tags
	TagLockedEdge
	TagSecretEdge
)

type Tag struct {
	Kind TagKind
	Id   int // should be unique in TagKind scope
}
