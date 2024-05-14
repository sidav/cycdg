package tiler

type StructTile struct {
	TileType     uint8
	nextTileType uint8 // for cellular automatas
}

func (st *StructTile) isAnyDoor() bool {
	return st.TileType == TileTypeDoor || st.TileType == TileTypeLockedDoor || st.TileType == TileTypeSecretDoor
}

const (
	TileTypeUnset uint8 = iota
	TileTypeRoomFloor
	TileTypeCaveFloor
	TileTypeWall
	TileTypeBarrier

	// Door tiles
	TileTypeDoor
	TileTypeLockedDoor
	TileTypeSecretDoor
)
