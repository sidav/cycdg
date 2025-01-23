package vaultstiler

import (
	"cycdg/graph_replacement/grid_graph/graph_element"
	vaultsdict "cycdg/vaults_tiler/vaults_dict"
)

func (t *VaultsTiler) areCoordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(t.tiledMap) && y < len(t.tiledMap[x])
}

func (t *VaultsTiler) areAdjacent2TilesWalkable(x, y int, horiz bool) bool {
	x1, y1 := x-1, y
	x2, y2 := x+1, y
	if horiz {
		x1, y1 = x, y-1
		x2, y2 = x, y+1
	}
	if !(t.areCoordsValid(x1, y1) && t.areCoordsValid(x2, y2)) {
		return false
	}
	return t.tiledMap[x1][y1].IsAnyWalkable() && t.tiledMap[x2][y2].IsAnyWalkable()
}

func (t *VaultsTiler) countTileTypesIn8Around(x, y int, countOOB bool, types ...uint8) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i != x || j != y {
				if t.areCoordsValid(i, j) {
					for _, typ := range types {
						if t.tiledMap[i][j].TileType == typ {
							count++
							break
						}
					}
				} else if countOOB {
					count++
				}
			}
		}
	}
	return count
}

func (t *VaultsTiler) countTileTypesInPlusAround(x, y int, countOOB bool, types ...uint8) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i*j == 0) && (i != 0 || j != 0) {
				if t.areCoordsValid(x+i, y+j) {
					for _, typ := range types {
						if t.tiledMap[x+i][y+j].TileType == typ {
							count++
							break
						}
					}
				} else if countOOB {
					count++
				}
			}
		}
	}
	return count
}

func (t *VaultsTiler) countTileTypesInRadiusAround(x, y, r int, types ...uint8) int {
	count := 0
	for _, typ := range types {
		for i := x - r; i <= x+r; i++ {
			for j := y - r; j <= y+r; j++ {
				if (i != x || j != y) && t.areCoordsValid(i, j) && t.tiledMap[i][j].TileType == typ {
					count++
				}
			}
		}
	}
	return count
}

func (t *VaultsTiler) execFuncAtEachTile(execFunc func(x, y int)) {
	for x := range t.tiledMap {
		for y := range t.tiledMap[x] {
			execFunc(x, y)
		}
	}
}

func (t *VaultsTiler) applyVaultAt(v *vaultsdict.Vault, graphX, graphY int) {
	for vx := 0; vx < t.nodeSize; vx++ {
		for vy := 0; vy < t.nodeSize; vy++ {
			chr := v.GetCharAt(vx, vy)
			tileTypeToSet := TileTypeUnset
			switch chr {
			case '.':
				tileTypeToSet = TileTypeRoomFloor
			case '#':
				tileTypeToSet = TileTypeWall
			}
			t.tiledMap[vx+graphX*(t.nodeSize+1)][vy+graphY*(t.nodeSize+1)].TileType = tileTypeToSet
		}
	}
}

// If "horiz" is false, will be applied as a vertical column
func (t *VaultsTiler) applyWallParcelForEdgeAt(w *vaultsdict.Wall, edge *graph_element.Edge, tx, ty int, horiz bool) {
	if w == nil { // Needed for debug, remove
		for z := 0; z < t.nodeSize; z++ {
			if horiz {
				t.tiledMap[tx+z][ty].TileType = TileTypeUnset
			} else {
				t.tiledMap[tx][ty+z].TileType = TileTypeUnset
			}
		}
		return
	}

	for z := 0; z < t.nodeSize; z++ {
		tileTypeToSet := TileTypeUnset
		chr := w.GetCharAt(z)
		switch chr {
		case '.':
			tileTypeToSet = TileTypeRoomFloor
		case '#':
			tileTypeToSet = TileTypeWall
		case '+':
			tileTypeToSet = t.getTileTypeForEdgeDoor(edge)
		case '?':
			tileTypeToSet = TileTypeSecretDoor
		case '\'':
			tileTypeToSet = TileTypeWindow
		default:
			panicf("Unknown char in parcel: [ %s ]", string(chr))
		}
		if horiz {
			t.tiledMap[tx+z][ty].TileType = tileTypeToSet
		} else {
			t.tiledMap[tx][ty+z].TileType = tileTypeToSet
		}
	}
}

func (t *VaultsTiler) getTileTypeForEdgeDoor(e *graph_element.Edge) uint8 {
	tags := e.GetTags()
	if len(tags) > 0 {
		tag := tags[0]
		switch tag.Kind {
		case graph_element.TagLockedEdge, graph_element.TagBilockedEdge, graph_element.TagMasterLockedEdge:
			return TileTypeLockedDoor
		case graph_element.TagSecretEdge:
			return TileTypeSecretDoor
		}
	}
	return TileTypeDoor
}
