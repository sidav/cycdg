package tiler

type Tile struct {
	TileType uint8
}

const (
	TileTypeUnset uint8 = iota
	TileTypeFloor
	TileTypeWall
	TileTypeBarrier
	TileTypeDoor
)

func (t *Tiler) GetTileMap() [][]Tile {
	// create the map itself
	t.setInitialTileMap()
	return t.tiledMap
}
