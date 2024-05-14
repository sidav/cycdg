package tiler

func (t *Tiler) GetTileMap() [][]StructTile {
	// create the map itself
	t.setInitialTileMap()
	t.doCellularAutomatae()

	// TODO: remove this
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		wallsOrUnsets := t.countTileTypesIn8Around(x, y, true, TileTypeWall, TileTypeUnset)
		if x == 0 || y == 0 || x == len(t.tiledMap)-1 || y == len(t.tiledMap[x])-1 {
			t.tiledMap[x][y].nextTileType = TileTypeWall
		}
		if wallsOrUnsets == 8 {
			t.tiledMap[x][y].nextTileType = TileTypeUnset
		}
	})

	return t.tiledMap
}
