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
	// It's needed for optimization, in order not to check each and every coord out there for each tile
	searchNearPrevIndex []int // -1 means "any coords"

	applicabilityFuncs []func(g *Graph, c Coords, prevСoords ...Coords) bool
	ApplyToGraph       func(g *Graph, applyAt ...Coords)
	MandatoryFeatures  []*FeatureAdder // One (and only) of them SHOULD apply! (May have nil though)
	OptionalFeatures   []*FeatureAdder // One (or more?) of them could be applied. Should NOT conflict with any of the mandatory and optional features.
}

func (ir *ReplacementRule) isApplicableAtStep(currentStep int) bool {
	if (ir.Metadata.StepApplicability == nil) {
		return true
	}
	return ir.Metadata.StepApplicability(currentStep)
}

func (ir *ReplacementRule) FindAllApplicableCoordVariantsRecursively(g *Graph) (result [][]Coords) {
	return ir.tryFindAllCoordVariantsRecursively(g)
}

func (ir *ReplacementRule) tryFindAllCoordVariantsRecursively(g *Graph, picked ...Coords) [][]Coords {
	stepIndex := len(picked)
	w, h := g.GetSize()
	if len(ir.searchNearPrevIndex) != len(ir.applicabilityFuncs) {
		debugPanic("Rule %s has wrong searchNearPrevIndex count", ir.Name)
	}

	xFrom, xTo, yFrom, yTo := ir.stepSearchBounds(stepIndex, picked, w, h)

	var result [][]Coords
	for x := xFrom; x <= xTo; x++ {
		for y := yFrom; y <= yTo; y++ {
			if !ir.WorksWithFinalizedNodes && g.IsNodeFinalized(x, y) {
				continue
			}
			if geometry.AreXYCoordsInCoordsArray(x, y, picked) {
				continue
			}
			if !ir.applicabilityFuncs[stepIndex](g, NewCoords(x, y), picked...) {
				continue
			}

			next := appendCoordCopy(picked, NewCoords(x, y))
			if stepIndex < len(ir.applicabilityFuncs)-1 {
				result = append(result, ir.tryFindAllCoordVariantsRecursively(g, next...)...)
			} else {
				result = append(result, next)
			}
		}
	}
	return result
}

// stepSearchBounds returns the grid area to search for the current step.
// If searchNearPrevIndex[step] != -1, narrows to 3×3 around the referenced picked coord.
func (ir *ReplacementRule) stepSearchBounds(step int, picked []Coords, w, h int) (xFrom, xTo, yFrom, yTo int) {
	xFrom, xTo = 0, w-1
	yFrom, yTo = 0, h-1
	if nearIdx := ir.searchNearPrevIndex[step]; nearIdx != -1 {
		cx, cy := picked[nearIdx].Unwrap()
		xFrom, yFrom = maxint(cx-1, 0), maxint(cy-1, 0)
		xTo, yTo = minint(cx+1, w-1), minint(cy+1, h-1)
	}
	return
}

// appendCoordCopy returns a fresh slice — avoids append aliasing when multiple
// loop iterations share the same backing array behind `coords`.
func appendCoordCopy(coords []Coords, c Coords) []Coords {
	result := make([]Coords, len(coords)+1)
	copy(result, coords)
	result[len(coords)] = c
	return result
}
