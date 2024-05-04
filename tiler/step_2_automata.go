package tiler

func (t *Tiler) doCellularAutomatae() {
	t.resetCellularAutomataNextState()

	// Thin all the doors
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		roomFloors := t.countTileTypeInPlusAround(TileTypeRoomFloor, x, y)
		caveFloors := t.countTileTypeInPlusAround(TileTypeCaveFloor, x, y)
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
			t.countTileTypeInPlusAround(TileTypeDoor, x, y) == 1 {
			t.tiledMap[x][y].nextTileType = TileTypeWall
		}
	})

	// Grow the rooms for the square look (remove corners for rooms)
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		floors8 := t.countTileTypeIn8Around(TileTypeRoomFloor, x, y)
		floorsPlus := t.countTileTypeInPlusAround(TileTypeRoomFloor, x, y)
		wallsPlus := t.countTileTypeInPlusAround(TileTypeWall, x, y)

		if t.tiledMap[x][y].TileType == TileTypeUnset &&
			floorsPlus == 2 && (floors8 == 5 || floors8 == 3) &&
			wallsPlus == 0 {
			t.tiledMap[x][y].nextTileType = TileTypeRoomFloor
		}
	})

	// // Grow the rooms (increase size) - TODO: how to do it?!
	// t.execFuncAsCAStep(1, func(x, y int) {
	// 	floors4 := t.countTileTypeInPlusAround(TileTypeRoomFloor, x, y)
	// 	floors8 := t.countTileTypeIn8Around(TileTypeRoomFloor, x, y)
	// 	unsets := t.countTileTypeIn8Around(TileTypeUnset, x, y)
	// 	if t.tiledMap[x][y].TileType == TileTypeUnset && floors4 == 1 {
	// 		if unsets == 5 || unsets == 6 && floors8 == 2 {
	// 			t.tiledMap[x][y].nextTileType = TileTypeRoomFloor
	// 		}
	// 	}
	// })

	// Grow the walls near the room floors
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		walls := t.countTileTypeInPlusAround(TileTypeWall, x, y)
		unsets := t.countTileTypeInPlusAround(TileTypeUnset, x, y)
		floors8 := t.countTileTypeIn8Around(TileTypeRoomFloor, x, y)
		floorsPlus := t.countTileTypeInPlusAround(TileTypeRoomFloor, x, y)
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
	// TODO: except the secret doors here!
	t.repeatedlyExecFuncAsCAStep(func(x, y int) {
		floorsPlus := t.countTileTypeInPlusAround(TileTypeCaveFloor, x, y)
		if t.tiledMap[x][y].TileType == TileTypeDoor {
			if floorsPlus == 2 {
				t.tiledMap[x][y].nextTileType = TileTypeCaveFloor
			}
		}
	})

	// Erode the caves' walls
	t.execFuncAsCAStep(4, func(x, y int) {
		rfloors := t.countTileTypeIn8Around(TileTypeRoomFloor, x, y)
		cfloors8 := t.countTileTypeIn8Around(TileTypeCaveFloor, x, y)
		cfloors4 := t.countTileTypeInPlusAround(TileTypeCaveFloor, x, y)
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
		cfloors8 := t.countTileTypeIn8Around(TileTypeCaveFloor, x, y)
		if t.tiledMap[x][y].TileType == TileTypeCaveFloor {
			if cfloors8 > 7 && rndChancePercent(40) {
				t.tiledMap[x][y].nextTileType = TileTypeWall
			}
		}
	})

	// Erode/dilate the caves' walls
	t.execFuncAsCAStep(3, func(x, y int) {
		cfloors8 := t.countTileTypeIn8Around(TileTypeCaveFloor, x, y)
		rfloors8 := t.countTileTypeIn8Around(TileTypeRoomFloor, x, y)
		// cfloors4 := t.countTileTypeInPlusAround(TileTypeCaveFloor, x, y)
		doors4 := t.countTileTypeInPlusAround(TileTypeDoor, x, y)
		walls4 := t.countTileTypeInPlusAround(TileTypeWall, x, y)
		walls8 := t.countTileTypeIn8Around(TileTypeWall, x, y)
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
