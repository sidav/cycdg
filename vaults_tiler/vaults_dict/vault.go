package vaultsdict

// Vault is a MxN predefined "map piece" which is used to fill the map
type Vault struct {
	applicableDoorsMask int
	charmap             []string
	implementsTags      []string

	// inner:
	rotation int // How many times the vault is rotated CW (used in order not to rotate the charmap)
}

func (v *Vault) Rotate(times int) {
	for i := 0; i < times; i++ {
		v.applicableDoorsMask = rotateDirFlagsCW(v.applicableDoorsMask)
		v.rotation++
	}
	v.rotation = v.rotation % 4
}

func (v *Vault) GetCharAt(x, y int) rune {
	// Consider the rotation
	x, y = rotateXYAroundCenterOfSquareCW(x, y, len(v.charmap)-1, v.rotation)
	s := v.charmap[y]
	return rune(s[x])
}
