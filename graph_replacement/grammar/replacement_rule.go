package grammar

import (
	"cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
)

// it's a replacement rule indifferent to mirroring and rotations
type ReplacementRule struct {
	Name string

	Metadata ruleMetadata

	WorksWithFinalizedNodes bool // true if coords search should not skip finalized nodes; USE CAUTIOUSLY

	// each value is coords index, near which the applicability func will be checked
	searchNearPrevIndex []int // -1 means "any coords"

	applicabilityFuncs []func(g *Graph, x, y int, prevСoords ...Coords) bool
	ApplyToGraph       func(g *Graph, applyAt ...Coords)
	MandatoryFeatures  []*FeatureAdder // One (and only) of them SHOULD apply! (May have nil though)
	OptionalFeatures   []*FeatureAdder // One (or more?) of them could be applied. Should NOT conflict with any of the mandatory and optional features.
}

func (ir *ReplacementRule) FindAllApplicableCoordVariantsRecursively(g *Graph) (result [][]Coords) {
	return ir.tryFindAllCoordVariantsRecursively(g)
}

func (ir *ReplacementRule) tryFindAllCoordVariantsRecursively(g *Graph, argsForFunc ...Coords) [][]Coords {
	currFuncIndex := len(argsForFunc)
	w, h := g.GetSize()
	var result [][]Coords

	xFrom, xTo := 0, w-1
	yFrom, yTo := 0, h-1
	if len(ir.searchNearPrevIndex) != len(ir.applicabilityFuncs) {
		debugPanic("Rule %s has wrong searchNearPrevIndex count", ir.Name)
	}
	if ir.searchNearPrevIndex[currFuncIndex] != -1 {
		searchNearX, searchNearY := argsForFunc[ir.searchNearPrevIndex[currFuncIndex]].Unwrap()
		xFrom, yFrom = maxint(searchNearX-1, 0), maxint(searchNearY-1, 0)
		xTo, yTo = minint(searchNearX+1, w-1), minint(searchNearY+1, h-1)
	}

	// try all coordinates
	for x := xFrom; x <= xTo; x++ {
		for y := yFrom; y <= yTo; y++ {

			if !ir.WorksWithFinalizedNodes {
				if g.IsNodeFinalized(x, y) {
					continue
				}
			}

			if geometry.AreXYCoordsInCoordsArray(x, y, argsForFunc) {
				continue
			}
			if ir.applicabilityFuncs[currFuncIndex](g, x, y, argsForFunc...) {
				// This function is not the last in rule
				if currFuncIndex < len(ir.applicabilityFuncs)-1 {
					// Warning: possible bug in append(argsForFunc, NewCoords(x, y)) usage.
					res := ir.tryFindAllCoordVariantsRecursively(g, append(argsForFunc, NewCoords(x, y))...)
					if len(res) > 0 { // next coords are good, so we can add them to the list
						result = append(result, res...)
					}
				} else { // it's last in rule, should add the previous and current coords to the list
					allArguments := make([]Coords, len(argsForFunc)+1)
					copy(allArguments, argsForFunc)
					allArguments[len(argsForFunc)] = NewCoords(x, y)
					result = append(result, allArguments)
					// Prevously: result = append(result, append(argsForFunc, NewCoords(x, y)))
					// It's fixed; was severely buggy with large total coords num (was observed with 8 coords, was ok with 6 or less)
					// Thus append() seems to be destructive to argsForFunc in that case; it's better to use slice copy for appending to.
				}
			}
		}
	}
	return result
}
