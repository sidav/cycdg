package tiler

// All funcs there perform operations using tiler.tileMap as cellular automata state

func (t *Tiler) resetCellularAutomataNextState() {
	t.execFuncAtEachTile(func(x, y int) {
		t.tiledMap[x][y].nextTileType = t.tiledMap[x][y].TileType
	})
}

func (t *Tiler) propagateTileTypes() bool {
	changed := false
	for x := range t.tiledMap {
		for y := range t.tiledMap[x] {
			if t.tiledMap[x][y].TileType != t.tiledMap[x][y].nextTileType {
				changed = true
				t.tiledMap[x][y].TileType = t.tiledMap[x][y].nextTileType
			}
		}
	}
	return changed
}

// Execs the func for EACH tile of the map, repeatedly, until changes stop occuring
func (t *Tiler) repeatedlyExecFuncAsCAStep(step func(x, y int)) {
	changed := true
	for changed {
		t.execFuncAtEachTile(step)
		changed = t.propagateTileTypes()
	}
}

// Execs the func for EACH tile of the map, repeatedly, until changes stop occuring
func (t *Tiler) execFuncAsCAStep(times int, step func(x, y int)) {
	changed := true
	for times > 0 && changed {
		t.execFuncAtEachTile(step)
		changed = t.propagateTileTypes()
		times--
	}
}
