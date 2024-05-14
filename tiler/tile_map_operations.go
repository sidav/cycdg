package tiler

func (t *Tiler) areCoordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(t.tiledMap) && y < len(t.tiledMap[x])
}

func (t *Tiler) countTileTypesIn8Around(x, y int, countOOB bool, types ...uint8) int {
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

func (t *Tiler) countTileTypesInPlusAround(x, y int, countOOB bool, types ...uint8) int {
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

func (t *Tiler) countTileTypesInRadiusAround(x, y, r int, types ...uint8) int {
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

func (t *Tiler) execFuncAtEachTile(execFunc func(x, y int)) {
	for x := range t.tiledMap {
		for y := range t.tiledMap[x] {
			execFunc(x, y)
		}
	}
}
