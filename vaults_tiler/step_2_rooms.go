package vaultstiler

import (
	"cycdg/vaults_tiler/vaults_dict"
)

func (t *VaultsTiler) SetRoomVaults() {
	w, h := t.graph.GetSize()
	// fill the room tiles on the map
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			node := t.graph.NodeAt(x, y)
			if node.IsActive() {
				vlt := t.getParcelForNodeAt(x, y)
				t.applyVaultAt(vlt, x, y)
			}
		}
	}
}

func (t *VaultsTiler) getParcelForNodeAt(nodeX, nodeY int) *vaultsdict.Vault {
	doorFlags := 0
	if t.graph.IsEdgeByVectorActive(nodeX, nodeY, 0, -1) {
		doorFlags |= vaultsdict.DIR_NORTH
	}
	if t.graph.IsEdgeByVectorActive(nodeX, nodeY, 1, 0) {
		doorFlags |= vaultsdict.DIR_EAST
	}
	if t.graph.IsEdgeByVectorActive(nodeX, nodeY, 0, 1) {
		doorFlags |= vaultsdict.DIR_SOUTH
	}
	if t.graph.IsEdgeByVectorActive(nodeX, nodeY, -1, 0) {
		doorFlags |= vaultsdict.DIR_WEST
	}
	return t.dict.GetRandomVaultByParams(doorFlags)
}
