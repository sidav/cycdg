package graph_element

type Edge struct {
	enabled     bool
	directional bool
	dirReversed bool
	tags        []string
}

func (e *Edge) Reset() {
	e.enabled = false
	e.directional = false
	e.dirReversed = false
	// if len(e.tags) > 0 {
	// 	panic("Tagged node being reset!")
	// }
}

func (e *Edge) GetTags() []string {
	return e.tags
}

func (e *Edge) IsDirectional() bool {
	return e.directional
}

func (e *Edge) IsReverse() bool {
	return e.dirReversed
}

func (e *Edge) SetDirection(r bool) {
	e.directional = true
	e.dirReversed = r
}

func (e *Edge) IsActive() bool {
	return e.enabled
}

func (e *Edge) AddTag(t string) {
	e.tags = append(e.tags, t)
}

func (e *Edge) HasTag(t string) bool {
	for _, s := range e.tags {
		if s == t {
			return true
		}
	}
	return false
}
