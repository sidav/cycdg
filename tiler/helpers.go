package tiler

import "time"

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
