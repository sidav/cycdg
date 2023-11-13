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
	afunc func(g *Graph, x, y int, prevСoords ...Coords) bool, argsForFunc ...Coords) []Coords {
	var crds []Coords
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if afunc(g, x, y, argsForFunc...) {
				crds = append(crds, NewCoords(x, y))
			}
		}
	}
	return crds
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

var URR = indifferentRule{
	applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
		// first node
		func(g *Graph, x, y int, prevСoords ...Coords) bool {
			return g.IsNodeActive(x, y)
		},
		// second node
		func(g *Graph, x, y int, prevСoords ...Coords) bool {
			x1, y1 := prevСoords[0].Unwrap()
			return areCoordsAdjacent(x, y, x1, y1) && g.IsNodeActive(x, y) && g.IsEdgeDirectedBetweenCoords(x1, y1, x, y)
		},
		// third node
		func(g *Graph, x, y int, prevСoords ...Coords) bool {
			x1, y1 := prevСoords[0].Unwrap()
			return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x1, y1)
		},
		// fourth node
		func(g *Graph, x, y int, prevСoords ...Coords) bool {
			x2, y2 := prevСoords[1].Unwrap()
			x3, y3 := prevСoords[2].Unwrap()
			return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x2, y2) && areCoordsAdjacent(x, y, x3, y3)
		},
	},
	applyToCoords: func(g *Graph, allCoords ...Coords) {
		g.enableNode(allCoords[2][0], allCoords[2][1])
		g.enableNode(allCoords[3][0], allCoords[3][1])
		g.setLinkBetweenCoords(allCoords[0][0], allCoords[0][1], allCoords[1][0], allCoords[1][1], false)
		g.enableDirectionalLinkBetweenCoords(allCoords[0][0], allCoords[0][1], allCoords[2][0], allCoords[2][1])
		g.enableDirectionalLinkBetweenCoords(allCoords[2][0], allCoords[2][1], allCoords[3][0], allCoords[3][1])
		g.enableDirectionalLinkBetweenCoords(allCoords[3][0], allCoords[3][1], allCoords[1][0], allCoords[1][1])
	},
}
