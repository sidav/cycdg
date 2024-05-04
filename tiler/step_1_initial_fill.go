package tiler

// Build map of 5x5 (or other, set in NodeSize) tiles of appropriate type
func (t *Tiler) setInitialTileMap() {
	// create the map itself
	w, h := t.graph.GetSize()
	t.tiledMap = make([][]StructTile, w*2*t.nodeSize)
	for i := range t.tiledMap {
		t.tiledMap[i] = make([]StructTile, h*2*t.nodeSize)
		for k := range t.tiledMap[i] {
			t.tiledMap[i][k].TileType = TileTypeUnset
		}
	}
	// fill the room tiles on the map
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if t.graph.NodeAt(x, y).IsActive() {
				tag := TileTypeCaveFloor
				if t.graph.NodeAt(x, y).HasAnyTags() {
					tag = TileTypeRoomFloor
				}
				t.fillSquare(x*2, y*2, tag)
			}
		}
	}
	// fill the doors and/or barriers on the map
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if !t.graph.NodeAt(x, y).IsActive() {
				continue
			}
			// check right edge
			if x < w-1 && t.graph.NodeAt(x+1, y).IsActive() {
				if t.graph.GetEdgeBetweenIntCoords(x, y, x+1, y).IsActive() {
					t.fillSquare(x*2+1, y*2, TileTypeDoor)
				} else {
					t.fillSquare(x*2+1, y*2, TileTypeBarrier)
				}
			}
			// check bottom edge
			if y < h-1 && t.graph.NodeAt(x, y+1).IsActive() {
				if t.graph.GetEdgeBetweenIntCoords(x, y, x, y+1).IsActive() {
					t.fillSquare(x*2, y*2+1, TileTypeDoor)
				} else {
					t.fillSquare(x*2, y*2+1, TileTypeBarrier)
				}
			}
		}
	}
}

func (t *Tiler) fillSquare(x, y int, tileType uint8) {
	startX := t.nodeSize * x
	startY := t.nodeSize * y
	for i := 0; i < t.nodeSize; i++ {
		for j := 0; j < t.nodeSize; j++ {
			t.tiledMap[startX+i][startY+j].TileType = tileType
		}
	}
}
