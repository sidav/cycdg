package vaultsdict

import "cycdg/graph_replacement/grid_graph/graph_element"

// Vault is a MxN predefined "map piece" which is used to fill the map
type Wall struct {
	charmap        string // Length SHOULD be equal to dictionary vaultSize!
	implementsTags []graph_element.TagKind
	resrictNullTag bool // false if the parcel can be applied as an edge with no tag, true otherwise

	// inner:
	mirror bool // How many times the vault is rotated CW (used in order not to rotate the charmap)
}

func (w *Wall) Mirror() {
	w.mirror = !w.mirror
}

func (w *Wall) GetCharAt(x int) rune {
	// Consider the rotation
	if w.mirror {
		x = len(w.charmap) - x - 1
	}
	return rune(w.charmap[x])
}

func (w *Wall) DoesCoordRequireAdjacentFloor(x int) bool {
	// Consider the rotation
	if w.mirror {
		x = len(w.charmap) - x - 1
	}
	run := rune(w.charmap[x])
	return run == '+' || run == '?'
}

func (w *Wall) ImplementsNullTag() bool {
	return !w.resrictNullTag
}

func (w *Wall) ImplementsEdgeTag(tag graph_element.TagKind) bool {
	for i := range w.implementsTags {
		if w.implementsTags[i] == tag {
			return true
		}
	}
	return false
}
