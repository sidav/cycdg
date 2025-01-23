package vaultstiler

import (
	"cycdg/graph_replacement/grid_graph/graph_element"
	vaultsdict "cycdg/vaults_tiler/vaults_dict"
)

func (t *VaultsTiler) SetRoomDoors() {
	w, h := t.graph.GetSize()
	for nx := 0; nx < w; nx++ {
		for ny := 0; ny < h; ny++ {
			if nx < w-1 {
				t.addEastDoor(nx, ny)
			}
			if ny < h-1 {
				t.addSouthDoor(nx, ny)
			}
		}
	}
}

func (t *VaultsTiler) addEastDoor(nx, ny int) {
	edge := t.graph.GetEdgeByVector(nx, ny, 1, 0)
	if !edge.IsActive() {
		return
	}
	wallX := (t.nodeSize+1)*(nx+1) - 1
	wallY := (t.nodeSize + 1) * ny
	applicables := t.getCompatibleWallParcelsForEdgeWallAt(edge, wallX, wallY, false)
	if len(applicables) == 0 {
		// panicf("Vertical wall lookup failed, tags %d", edge.GetTags()[0].Kind)
		t.applyWallParcelForEdgeAt(nil, nil, wallX, wallY, false)
		return
	}
	t.applyWallParcelForEdgeAt(applicables[rnd(len(applicables))], edge, wallX, wallY, false)
}

func (t *VaultsTiler) addSouthDoor(nx, ny int) {
	edge := t.graph.GetEdgeByVector(nx, ny, 0, 1)
	if !edge.IsActive() {
		return
	}
	wallX := (t.nodeSize + 1) * nx
	wallY := (t.nodeSize+1)*(ny+1) - 1
	applicables := t.getCompatibleWallParcelsForEdgeWallAt(edge, wallX, wallY, true)
	if len(applicables) == 0 {
		// panicf("Horizontal wall lookup failed, tags %d", edge.GetTags()[0].Kind)
		t.applyWallParcelForEdgeAt(nil, nil, wallX, wallY, true)
		return
	}
	t.applyWallParcelForEdgeAt(applicables[rnd(len(applicables))], edge, wallX, wallY, true)
}

func (t *VaultsTiler) getCompatibleWallParcelsForEdgeWallAt(edge *graph_element.Edge, wallX, wallY int, horizontal bool) []*vaultsdict.Wall {
	applicables := make([]*vaultsdict.Wall, 0)
	for _, wParc := range t.dict.Walls {
		// Check if the tag is good
		if len(edge.GetTags()) > 0 && !wParc.ImplementsEdgeTag(edge.GetTags()[0].Kind) {
			continue
		} else if len(edge.GetTags()) == 0 && !wParc.ImplementsNullTag() {
			continue
		}
		good := true
		for parcelCoord := 0; parcelCoord < t.nodeSize; parcelCoord++ {
			walkable := t.areAdjacent2TilesWalkable(wallX, wallY+parcelCoord, false)
			if horizontal {
				walkable = t.areAdjacent2TilesWalkable(wallX+parcelCoord, wallY, true)
			}
			if wParc.DoesCoordRequireAdjacentFloor(parcelCoord) && !walkable {
				good = false
				break
			}
		}
		if good {
			applicables = append(applicables, wParc)
		}
	}
	return applicables
}
