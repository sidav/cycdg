package graph_element

import (
	"fmt"
)

type Node struct {
	active    bool
	finalized bool    // if true, restrict changing this node
	edges     [2]Edge // to the right and to the down
	tags      []*Tag
}

func (gn *Node) Init() {
	// gn.finalized = false
	gn.ResetActiveAndLinks()
	gn.tags = nil
}

func (gn *Node) Finalize() {
	gn.finalized = true
}

func (gn *Node) UnsafeUnfinalize() {
	gn.finalized = false
}

func (gn *Node) SwapTagsWith(gn2 *Node) {
	gn.tags, gn2.tags = gn2.tags, gn.tags
}

func (gn *Node) ResetActiveAndLinks() {
	if gn.finalized {
		panic("Node is finalized!")
	}
	gn.active = false
	for i := range gn.edges {
		gn.edges[i].enabled = false
	}
}

func (gn *Node) AddTag(tagType TagKind, tagId int) {
	if gn.finalized {
		panic("AddTag: Node is finalized!")
	}
	gn.tags = append(gn.tags, &Tag{Kind: tagType, Id: tagId})
}

func (gn *Node) RemoveTagByIndex(i int) {
	if gn.finalized {
		panic("AddTag: Node is finalized!")
	}
	gn.tags = append(gn.tags[:i], gn.tags[i+1:]...)
}

func (gn *Node) HasAnyTags() bool {
	return len(gn.tags) > 0
}

func (gn *Node) HasTag(t TagKind) bool {
	for _, tg := range gn.tags {
		if tg.Kind == t {
			return true
		}
	}
	return false
}

func (gn *Node) ResetTags() {
	gn.tags = nil
}

func (gn *Node) SetLinkByVector(vx, vy int, enabled, reverse bool) {
	e := gn.GetEdgeByVector(vx, vy)
	e.enabled = enabled
	e.dirReversed = reverse
}

func (gn *Node) HasLinkToVector(vx, vy int) bool {
	return gn.GetEdgeByVector(vx, vy).enabled
}

func (gn *Node) GetEdgeByVector(vx, vy int) *Edge {
	if gn == nil {
		panic(fmt.Sprintf("NIL NODE FOR GET EDGE BY VECTOR %d,%d", vx, vy))
	}
	if vx == 1 && vy == 0 {
		return &gn.edges[0]
	} else if vx == 0 && vy == 1 {
		return &gn.edges[1]
	}
	panic(fmt.Sprintf("Oh noes, bad vector %d,%d", vx, vy))
}

func (gn *Node) GetTags() []*Tag {
	return gn.tags
}

func (gn *Node) SetActive(lock bool) {
	gn.active = lock
}

func (gn *Node) IsActive() bool {
	return gn.active
}

func (gn *Node) IsFinalized() bool {
	return gn.finalized
}
