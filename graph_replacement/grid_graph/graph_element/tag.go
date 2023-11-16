package graph_element

type TagKind uint8

const (
	TagStart = iota
	TagGoal
	TagLockedEdge
	TagKeyForEdge
	TagBoss
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
	default:
		panic("No such idiom!")
	}
}
