package grammar

// import (
// 	. "cycdg/graph_replacement/geometry"
// 	. "cycdg/graph_replacement/grid_graph"
// 	"fmt"
// )

// func print(msg string, args ...interface{}) {
// 	fmt.Printf(sprintf(msg+"\n", args...))
// }

// func pp(name string, a []Coords) {
// 	fmt.Printf(sprintf("%s\n", name))
// 	for _, v := range a {
// 		fmt.Printf("%v\n", v)
// 	}
// }

// func ppp(name string, a [][]Coords) {
// 	fmt.Printf(sprintf("%s\n", name))
// 	for _, v := range a {
// 		fmt.Printf("%v\n", v)
// 	}
// }

// func (ir *ReplacementRule) getApplicableCoordsForFunc(g *Graph, index int, argsForFunc ...Coords) []Coords {
// 	var crds []Coords
// 	w, h := g.GetSize()

// 	xFrom, xTo := 0, w-1
// 	yFrom, yTo := 0, h-1
// 	if len(ir.searchNearPrevIndex) < len(ir.applicabilityFuncs) {
// 		debugPanic("Rule %s has wrong searchNearPrevIndex count", ir.Name)
// 	}
// 	if ir.searchNearPrevIndex[index] != -1 {
// 		searchNearX, searchNearY := argsForFunc[ir.searchNearPrevIndex[index]].Unwrap()
// 		xFrom, yFrom = maxint(searchNearX-1, 0), maxint(searchNearY-1, 0)
// 		xTo, yTo = minint(searchNearX+1, w-1), minint(searchNearY+1, h-1)
// 	}

// 	for x := xFrom; x <= xTo; x++ {
// 		for y := yFrom; y <= yTo; y++ {
// 			// TODO: maybe some rules should want to ignore that?..
// 			if g.IsNodeFinalized(x, y) {
// 				continue
// 			}
// 			// uniqueness check:
// 			if AreXYCoordsInCoordsArray(x, y, argsForFunc) {
// 				continue
// 			}
// 			if ir.applicabilityFuncs[index](g, x, y, argsForFunc...) {
// 				crds = append(crds, NewCoords(x, y))
// 			}
// 		}
// 	}
// 	return crds
// }

// func (ir *ReplacementRule) legacySearchApplicableCoords(g *Graph) (result [][]Coords) {
// 	var applicableArray [][]Coords
// 	var prevApplicableArray [][]Coords

// 	applicable0 := ir.getApplicableCoordsForFunc(g, 0)
// 	for i := range applicable0 {
// 		applicableArray = append(applicableArray, []Coords{applicable0[i]})
// 	}
// 	if len(ir.applicabilityFuncs) == 1 {
// 		return applicableArray
// 	}

// 	prevApplicableArray = applicableArray
// 	for funcnum := 1; funcnum < len(ir.applicabilityFuncs); funcnum++ {
// 		applicableArray = make([][]Coords, 0)
// 		for i := range prevApplicableArray {
// 			allcurr := ir.getApplicableCoordsForFunc(g, funcnum, prevApplicableArray[i]...)
// 			for j := range allcurr {
// 				newArray := make([]Coords, len(prevApplicableArray[i]))
// 				copy(newArray, prevApplicableArray[i])
// 				newArray = append(newArray, allcurr[j])
// 				applicableArray = append(applicableArray, newArray)
// 			}
// 			// if len(allcurr) > 0 && hasRepeats(applicableArray) {
// 			// 	ppp("BEFORE", prevApplicableArray)
// 			// 	ppp("Error", applicableArray)
// 			// 	pp("Caused by", allcurr)
// 			// 	debugPanic("Oh duck!")
// 			// }
// 		}
// 		// prevApplicableArray = make([][]Coords, len(applicableArray))
// 		// copy(prevApplicableArray, applicableArray)
// 		prevApplicableArray = applicableArray
// 	}

// 	return applicableArray
// }

// func hasRepeats(a [][]Coords) bool {
// 	for i := 0; i < len(a)-1; i++ {
// 		for j := i + 1; j < len(a); j++ {
// 			if arrsEqual(a[i], a[j]) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func arrsEqual(a, b []Coords) bool {
// 	for i := range a {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
