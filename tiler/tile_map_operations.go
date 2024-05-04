package tiler

func (t *Tiler) areCoordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(t.tiledMap) && y < len(t.tiledMap[x])
}

func (t *Tiler) countTileTypeIn8Around(tileType uint8, x, y int) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if (i != x || j != y) && t.areCoordsValid(i, j) && t.tiledMap[i][j].TileType == tileType {
				count++
			}
		}
	}
	return count
}

func (t *Tiler) countTileTypeInPlusAround(tileType uint8, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i*j == 0) && (i != 0 || j != 0) &&
				t.areCoordsValid(x+i, y+j) && t.tiledMap[x+i][y+j].TileType == tileType {
				count++
			}
		}
	}
	return count
}

func (t *Tiler) countAllTileTypesIn8Around(x, y int, types ...uint8) int {
	count := 0
	for _, typ := range types {
		count += t.countTileTypeIn8Around(typ, x, y)
	}
	return count
}

func (t *Tiler) countAllTileTypesInPlusAround(x, y int, types ...uint8) int {
	count := 0
	for _, typ := range types {
		count += t.countTileTypeInPlusAround(typ, x, y)
	}
	return count
}

func (t *Tiler) execFuncAtEachTile(execFunc func(x, y int)) {
	for x := range t.tiledMap {
		for y := range t.tiledMap[x] {
			execFunc(x, y)
		}
	}
}
