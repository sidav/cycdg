package tiler

type StructTile struct {
	TileType     uint8
	nextTileType uint8 // for cellular automatas
}

const (
	TileTypeUnset uint8 = iota
	TileTypeRoomFloor
	TileTypeCaveFloor
	TileTypeWall
	TileTypeBarrier
	TileTypeDoor
)
