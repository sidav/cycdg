package graph_element

type TagKind uint8

const (
	// Node tags
	TagStart = iota
	TagGoal
	TagKeyForEdge
	TagBoss
	TagTreasure
	// Edge tags
	TagLockedEdge
	TagSecretEdge
)

type Tag struct {
	Kind TagKind
	Id   int // should be unique in TagKind scope
}

func (t *Tag) GetStringIdiom() string {
	switch t.Kind {
	case TagStart:
		return "STRT"
	case TagGoal:
		return "GOAL"
	case TagKeyForEdge:
		return "KEY"
	case TagBoss:
		return "BOSS"
	case TagTreasure:
		return "TRSR"
	default:
		return "????"
	}
}
