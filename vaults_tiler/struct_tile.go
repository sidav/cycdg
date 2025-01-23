package vaultstiler

type StructTile struct {
	TileType     uint8
	nextTileType uint8 // for cellular automatas
}

func (st *StructTile) isAnyDoor() bool {
	return st.TileType == TileTypeDoor || st.TileType == TileTypeLockedDoor || st.TileType == TileTypeSecretDoor
}

func (st *StructTile) IsAnyWalkable() bool {
	return st.TileType == TileTypeRoomFloor || st.TileType == TileTypeCaveFloor
}

const (
	TileTypeUnset uint8 = iota
	TileTypeRoomFloor
	TileTypeCaveFloor
	TileTypeWall
	TileTypeBarrier
	TileTypeWindow
	// Door tiles
	TileTypeDoor
	TileTypeLockedDoor
	TileTypeSecretDoor
)
