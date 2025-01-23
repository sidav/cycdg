package vaultstiler

import (
	graph "cycdg/graph_replacement/grid_graph"
	vaultsdict "cycdg/vaults_tiler/vaults_dict"
)

// Transforms graph replacement result to a tiled map
type VaultsTiler struct {
	graph *graph.Graph
	dict  *vaultsdict.VaultsDictionary

	tiledMap [][]StructTile

	nodeSize int
}

func (t *VaultsTiler) Init(g *graph.Graph, dict *vaultsdict.VaultsDictionary) {
	t.graph = g
	t.dict = dict

	t.nodeSize = dict.VaultSize
}

func (t *VaultsTiler) GetTileMap() [][]StructTile {
	t.setInitialTileMap()
	t.SetRoomVaults()
	t.SetRoomDoors()
	// Fill everything that remains unset with walls
	// t.execFuncAtEachTile(
	// 	func(x, y int) {
	// 		if t.tiledMap[x][y].TileType == TileTypeUnset {
	// 			t.tiledMap[x][y].TileType = TileTypeWall
	// 		}
	// 	},
	// )

	// t.execFuncAtEachTile(
	// 	func(x, y int) {
	// 		if t.tiledMap[x][y].TileType == TileTypeWall && t.areAdjacent2TilesWalkable(x, y, false) {
	// 			t.tiledMap[x][y].TileType = TileTypeUnset
	// 		}
	// 	},
	// )
	return t.tiledMap
}
