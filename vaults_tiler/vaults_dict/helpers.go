package vaultsdict

import (
	"fmt"
	"time"
)

const (
	// Masks
	DIR_NORTH = 1
	DIR_EAST  = 1 << 1
	DIR_SOUTH = 1 << 2
	DIR_WEST  = 1 << 3
	DIR_ANY   = DIR_NORTH | DIR_EAST | DIR_SOUTH | DIR_WEST // alias
)

// Temporary one, should add proper PRNG later
func rnd(mod int) int {
	if mod == 0 {
		mod = 10000000
	}
	seed1 := int(time.Now().UnixNano()) % 10000
	seed2 := int(time.Now().UnixMicro()) % 10000
	seed3 := int(time.Now().UnixMilli()) % 10000
	return (seed1*10007 + seed2*503 + seed3) % mod
}

func rndChancePercent(perc int) bool {
	return rnd(100) < perc
}

func rotateCoordsCW(coords [2]int) [2]int {
	return [2]int{-coords[1], coords[0]}
}

func rotateCoordsCCW(coords [2]int) [2]int {
	return [2]int{coords[1], -coords[0]}
}

func rotateDirFlagsCCW(flags int) int {
	// 4-bit cyclic shift.
	// 1. Shift left, it's a rotation itself. Gather last 4 bits.
	// 2. Shift right, gather rightmost bit. it's the missing leftmost bit shifted in 1.
	// 3. Sum the values
	// return ((flags << 1) & 0b1111) | ((flags >> 3) & 0b1) // It works, but too complicated to read
	return ((flags << 1) | ((flags & 0b1000) >> 3)) & 0b1111 // The same as the comment above
}

func rotateDirFlagsCW(flags int) int {
	// Same as with CW, but the other way around
	// return ((flags >> 1) & 0b1111) | ((flags << 3) & 0b1000)
	return (flags >> 1) | ((flags & 0b0001) << 3)
}

func vectorsToDirFlags(vectors [][2]int) int {
	flags := 0
	for i := range vectors {
		x, y := vectors[i][0], vectors[i][1]
		if x == 0 && y == -1 {
			flags |= DIR_NORTH
		} else if x == 1 && y == 0 {
			flags |= DIR_EAST
		} else if x == 0 && y == 1 {
			flags |= DIR_SOUTH
		} else if x == -1 && y == 0 {
			flags |= DIR_WEST
		} else {
			panic("Bad vector given")
		}
	}
	return flags
}

func dirFlagsToVectors(flags int) [][2]int {
	vectors := make([][2]int, 0)
	if flags&DIR_NORTH != 0 {
		vectors = append(vectors, [2]int{0, -1})
	}
	if flags&DIR_EAST != 0 {
		vectors = append(vectors, [2]int{1, 0})
	}
	if flags&DIR_SOUTH != 0 {
		vectors = append(vectors, [2]int{0, 1})
	}
	if flags&DIR_WEST != 0 {
		vectors = append(vectors, [2]int{-1, 0})
	}
	return vectors
}

func rotateXYAroundCenterOfSquareCW(x, y, squareSideLength, timesToRotate int) (int, int) {
	// Odd-size correction
	x *= 2
	y *= 2
	squareSideLength *= 2

	halfSide := squareSideLength / 2
	relX, relY := x-halfSide, y-halfSide
	for i := 0; i < timesToRotate; i++ {
		relX, relY = -relY, relX
	}
	return (relX + halfSide) / 2, (relY + halfSide) / 2
}

// Returns how many times rotatingFlags should be rotated CW until rotatingFlags == equalTo. Returns -1 if it's not possible.
func rotationsUntilDirFlagsAreEqual(rotatingFlags, equalTo int) int {
	for rot := 0; rot < 4; rot++ {
		if rotatingFlags == equalTo {
			return rot
		}
		rotatingFlags = rotateDirFlagsCW(rotatingFlags)
	}
	return -1
}

// Returns how many times rotatingFlags should be rotated CW until all inclusion 1 bits are 1 in rotatingFlags. Returns -1 if it's not possible.
func rotationsUntilFlagsIncluded(rotatingFlags, inclusion int) int {
	for rot := 0; rot < 4; rot++ {
		if rotatingFlags&inclusion == inclusion {
			return rot
		}
		rotatingFlags = rotateDirFlagsCW(rotatingFlags)
	}
	return -1
}

func DoTestStuff() {
	// vector := [][2]int{
	// 	{0, -1},
	// 	{-1, 0},
	// 	{1, 0},
	// 	{0, 1},
	// }
	// flags := vectorsToDirFlags(vector)
	// fmt.Printf("Vector %v to flags: %b\n", vector, flags)
	// rotFlags := rotateDirFlagsCCW(flags)
	// fmt.Printf("Rotated flags: %b\n", rotFlags)
	// fmt.Printf("Rotated flags to array: %v\n", dirFlagsToVectors(rotFlags))
	// x, y := 4, 0
	// side := 5
	// rotx, roty := rotateXYAroundCenterOfSquareCW(x, y, side, 0)
	// fmt.Printf("%d, %d rotated for %dx%d pic: %d,%d\n", x, y, side, side, rotx, roty)
	fmt.Printf("Rotate %d times\n", rotationsUntilFlagsIncluded(DIR_NORTH|DIR_SOUTH, DIR_SOUTH))
}
