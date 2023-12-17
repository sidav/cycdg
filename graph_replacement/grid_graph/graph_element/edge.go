package graph_element

type Edge struct {
	enabled     bool
	directional bool
	dirReversed bool
	tags        []*Tag
}

func (e *Edge) Reset() {
	e.enabled = false
	e.directional = false
	e.dirReversed = false
	// if len(e.tags) > 0 {
	// 	panic("Tagged node being reset!")
	// }
}

func (e *Edge) IsDirectional() bool {
	return e.directional
}

func (e *Edge) IsReverse() bool {
	return e.dirReversed
}

func (e *Edge) IsActive() bool {
	return e.enabled
}

func (e *Edge) GetTags() []*Tag {
	return e.tags
}

func (e *Edge) SwapTagsWith(e2 *Edge) {
	e.tags, e2.tags = e2.tags, e.tags
}

func (e *Edge) CopyTagsFrom(from *Edge) {
	for _, t := range from.tags {
		e.AddTag(t.Kind, t.Id)
	}
}

func (e *Edge) AddTag(kind TagKind, id int) {
	e.tags = append(e.tags, &Tag{Kind: kind, Id: id})
}
