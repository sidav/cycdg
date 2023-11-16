package replacement

import . "cycdg/graph_replacement/grid_graph/geometry"
import . "cycdg/graph_replacement/grid_graph"

// TODO: delete this whole file?

func (ir *indifferentRule) getApplicableCoordsForFunc(g *Graph,
	afunc func(*Graph, int, int, ...Coords) bool, argsForFunc ...Coords) []Coords {
	var crds []Coords
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// TODO: maybe some rules should want to ignore that?..
			if g.IsNodeFinalized(x, y) {
				continue
			}
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

// func (ar *alternateRule) tryAllCoordsForFuncs(g *Graph) (result [][]Coords) {
// 	var applicableArray1 [][]Coords
// 	all1 := ir.getApplicableCoordsForFunc(g, ir.applicabilityFuncs[0])
// 	for i := range all1 {
// 		applicableArray1 = append(applicableArray1, []Coords{all1[i]})
// 	}
// 	if len(ir.applicabilityFuncs) == 1 {
// 		return applicableArray1
// 	}

// 	var applicableArray2 [][]Coords
// 	for i := range all1 {
// 		all2curr := ir.getApplicableCoordsForFunc(g, ir.applicabilityFuncs[1], all1[i])
// 		for j := range all2curr {
// 			applicableArray2 = append(applicableArray2, append(applicableArray1[i], all2curr[j]))
// 		}
// 	}
// 	if len(ir.applicabilityFuncs) == 2 {
// 		return applicableArray2
// 	}

// 	var applicableArray3 [][]Coords
// 	for i := range applicableArray2 {
// 		all3curr := ir.getApplicableCoordsForFunc(g, ir.applicabilityFuncs[2], applicableArray2[i]...)
// 		for j := range all3curr {
// 			applicableArray3 = append(applicableArray3, append(applicableArray2[i], all3curr[j]))
// 		}
// 	}
// 	if len(ir.applicabilityFuncs) == 3 {
// 		return applicableArray3
// 	}

// 	var applicableArray4 [][]Coords
// 	for i := range applicableArray3 {
// 		all4curr := ir.getApplicableCoordsForFunc(g, ir.applicabilityFuncs[3], applicableArray3[i]...)
// 		for j := range all4curr {
// 			applicableArray4 = append(applicableArray4, append(applicableArray3[i], all4curr[j]))
// 		}
// 	}
// 	if len(ir.applicabilityFuncs) == 4 {
// 		return applicableArray4
// 	}
// 	panic("too many dimensions?")
// }

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
