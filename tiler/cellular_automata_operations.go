package tiler

import "fmt"

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

// Tile types should be propagated after call!
func (t *Tiler) randomlySwapTileWithNeighbour(x, y int, typeToSwapWith uint8) {
	allowedArr := [][2]int{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 || i*j != 0 || !t.areCoordsValid(x+i, y+j) {
				continue
			}
			if t.tiledMap[x+i][y+j].TileType == typeToSwapWith {
				allowedArr = append(allowedArr, [2]int{x + i, y + j})
			}
		}
	}
	if len(allowedArr) > 0 {
		index := rnd(len(allowedArr))
		if false {
			panic(fmt.Sprintf("Length is %d, index is %d", len(allowedArr), index))
		}
		sx, sy := allowedArr[index][0], allowedArr[index][1]
		// panic(fmt.Sprintf("Swapping %d,%d with %d,%d", x, y, sx, sy))
		t.tiledMap[sx][sy].nextTileType = t.tiledMap[x][y].TileType
		t.tiledMap[x][y].nextTileType = typeToSwapWith
	}
}
