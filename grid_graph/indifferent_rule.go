package graph

import . "cycdg/grid_graph/geometry"

// it's a replacement rule indifferent to mirroring and rotations
type indifferentRule struct {
	Name string

	// metadata:
	AddsCycle bool

	applicabilityFuncs []func(g *Graph, x, y int, prevСoords ...Coords) bool
	applyToCoords      func(g *Graph, allCoords ...Coords)
}

func (ir *indifferentRule) getApplicableCoordsForFunc(g *Graph,
	afunc func(*Graph, int, int, ...Coords) bool, argsForFunc ...Coords) []Coords {
	var crds []Coords
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// uniqueness check:
			if areXYCoordsInCoordsArray(x, y, argsForFunc) {
				continue
			}

			if afunc(g, x, y, argsForFunc...) {
				crds = append(crds, NewCoords(x, y))
			}
		}
	}
	return crds
}

func areXYCoordsInCoordsArray(x, y int, coords []Coords) bool {
	for i := range coords {
		if coords[i].EqualsPair(x, y) {
			return true
		}
	}
	return false
}

func (ir *indifferentRule) FindAllApplicableCoordVariants(g *Graph) (result [][]Coords) {
	var applicableArray [][]Coords
	var prevApplicableArray [][]Coords

	applicable1 := ir.getApplicableCoordsForFunc(g, ir.applicabilityFuncs[0])
	for i := range applicable1 {
		applicableArray = append(applicableArray, []Coords{applicable1[i]})
	}
	if len(ir.applicabilityFuncs) == 1 {
		return applicableArray
	}

	prevApplicableArray = applicableArray
	for funcnum := 1; funcnum < len(ir.applicabilityFuncs); funcnum++ {
		applicableArray = nil
		for i := range prevApplicableArray {
			allcurr := ir.getApplicableCoordsForFunc(g, ir.applicabilityFuncs[funcnum], prevApplicableArray[i]...)
			for j := range allcurr {
				applicableArray = append(applicableArray, append(prevApplicableArray[i], allcurr[j]))
			}
		}
		prevApplicableArray = applicableArray
	}
	return applicableArray
}
