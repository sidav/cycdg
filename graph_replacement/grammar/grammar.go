package grammar

import "cycdg/lib/random"

var rnd random.PRNG

func SetRandom(r random.PRNG) {
	rnd = r
}
