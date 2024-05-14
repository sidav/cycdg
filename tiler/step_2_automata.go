package tiler

func (t *Tiler) doCellularAutomatae() {
	t.resetCellularAutomataNextState()

	// Thin all the doors
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		roomFloors := t.countTileTypesInPlusAround(x, y, false, TileTypeRoomFloor)
		caveFloors := t.countTileTypesInPlusAround(x, y, false, TileTypeCaveFloor)
		if t.tiledMap[x][y].TileType == TileTypeDoor {
			if roomFloors == 1 && caveFloors == 0 {
				t.tiledMap[x][y].nextTileType = TileTypeRoomFloor
			} else if roomFloors == 0 && caveFloors == 1 {
				t.tiledMap[x][y].nextTileType = TileTypeCaveFloor
			}
		}
	})

	// Leave only 1 tile for doors, wall everything else
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		if t.tiledMap[x][y].TileType == TileTypeDoor &&
			t.countTileTypesInPlusAround(x, y, false, TileTypeDoor) == 1 {
			t.tiledMap[x][y].nextTileType = TileTypeWall
		}
	})

	// Randomly displace doors along walls
	t.execFuncAsCAStep(2, func(x, y int) {
		if t.tiledMap[x][y].TileType == TileTypeDoor && rndChancePercent(25) {
			t.randomlySwapTileWithNeighbour(x, y, TileTypeWall)
		}
	})

	// Grow the rooms (increase size) - may give undesired results
	t.execFuncAsCAStep(rnd(3), func(x, y int) {
		floors4 := t.countTileTypesInPlusAround(x, y, false, TileTypeRoomFloor)
		unsets := t.countTileTypesInPlusAround(x, y, false, TileTypeUnset)
		if t.tiledMap[x][y].TileType == TileTypeUnset && floors4+unsets == 4 {
			if floors4 != 2 && unsets != 4 {
				t.tiledMap[x][y].nextTileType = TileTypeRoomFloor
			}
		}
	})

	// Grow the rooms for the square look (remove corners for rooms)
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		floors8 := t.countTileTypesIn8Around(x, y, true, TileTypeRoomFloor)
		floorsPlus := t.countTileTypesInPlusAround(x, y, false, TileTypeRoomFloor)
		wallsPlus := t.countTileTypesInPlusAround(x, y, false, TileTypeWall)

		if t.tiledMap[x][y].TileType == TileTypeUnset &&
			floorsPlus == 2 && (floors8 == 5 || floors8 == 3) &&
			wallsPlus == 0 {
			t.tiledMap[x][y].nextTileType = TileTypeRoomFloor
		}
	})

	// Grow the walls near the room floors
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		walls := t.countTileTypesInPlusAround(x, y, false, TileTypeWall)
		unsets := t.countTileTypesInPlusAround(x, y, false, TileTypeUnset)
		floors8 := t.countTileTypesIn8Around(x, y, true, TileTypeRoomFloor)
		floorsPlus := t.countTileTypesInPlusAround(x, y, false, TileTypeRoomFloor)
		if (t.tiledMap[x][y].TileType == TileTypeUnset || t.tiledMap[x][y].TileType == TileTypeBarrier) &&
			(floors8 > 0 && (walls == 1 || walls == 4 || walls == 2 && floorsPlus > 0) ||
				floors8 > 0 && walls == 2 && unsets == 2) {
			t.tiledMap[x][y].nextTileType = TileTypeWall
		}
	})

	// Replace all unsets and barriers with walls
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		if t.tiledMap[x][y].TileType == TileTypeUnset || t.tiledMap[x][y].TileType == TileTypeBarrier {
			t.tiledMap[x][y].nextTileType = TileTypeWall
		}
	})

	// CAVES

	// Remove cave-to-cave doors
	// TODO: except secret and keyed doors here!
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		floorsPlus := t.countTileTypesInPlusAround(x, y, false, TileTypeCaveFloor)
		if t.tiledMap[x][y].TileType == TileTypeDoor {
			if floorsPlus == 2 {
				t.tiledMap[x][y].nextTileType = TileTypeCaveFloor
			}
		}
	})

	// Erode the caves' walls
	t.execFuncAsCAStep(4, func(x, y int) {
		rfloors := t.countTileTypesIn8Around(x, y, true, TileTypeRoomFloor)
		cfloors8 := t.countTileTypesIn8Around(x, y, true, TileTypeCaveFloor)
		cfloors4 := t.countTileTypesInPlusAround(x, y, false, TileTypeCaveFloor)
		if t.tiledMap[x][y].TileType == TileTypeWall && rndChancePercent(50) {
			if rfloors == 0 && (cfloors8 > 2 || cfloors4 > 2) {
				t.tiledMap[x][y].nextTileType = TileTypeCaveFloor
			}
		}
	})

	// Place "wall seeds" here and there
	t.execFuncAsCAStep(1, func(x, y int) {
		// cfloorsRadius2 := t.countTileTypesInRadiusAround(x, y, 2, TileTypeCaveFloor)
		// cWallsRadius2 := t.countTileTypesInRadiusAround(x, y, 2, TileTypeWall)
		cfloors8 := t.countTileTypesIn8Around(x, y, true, TileTypeCaveFloor)
		if t.tiledMap[x][y].TileType == TileTypeCaveFloor {
			if cfloors8 > 7 && rndChancePercent(40) {
				t.tiledMap[x][y].nextTileType = TileTypeWall
			}
		}
	})

	// Erode/dilate the caves' walls
	t.execFuncAsCAStep(3, func(x, y int) {
		cfloors8 := t.countTileTypesIn8Around(x, y, true, TileTypeCaveFloor)
		rfloors8 := t.countTileTypesIn8Around(x, y, true, TileTypeRoomFloor)
		// cfloors4 := t.countAllTileTypesInPlusAround(x, y, false, TileTypeCaveFloor)
		doors4 := t.countTileTypesInPlusAround(x, y, false, TileTypeDoor)
		walls4 := t.countTileTypesInPlusAround(x, y, true, TileTypeWall)
		walls8 := t.countTileTypesIn8Around(x, y, true, TileTypeWall)
		wallsR2 := t.countTileTypesInRadiusAround(x, y, 2, TileTypeWall)
		if doors4 == 0 && rfloors8 == 0 {
			if t.tiledMap[x][y].TileType == TileTypeCaveFloor {
				if walls4 != 2 && (walls8 >= 6 || wallsR2 <= 2) {
					t.tiledMap[x][y].nextTileType = TileTypeWall
				}
			}
			if t.tiledMap[x][y].TileType == TileTypeWall {
				if cfloors8 > 0 && (walls8 < 4 && wallsR2 > 3) {
					t.tiledMap[x][y].nextTileType = TileTypeCaveFloor
				}
			}
		}
	})
}
