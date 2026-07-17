package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
)

var allNonGrowingRules = []*ReplacementRule{

	// 0  X  ; just finalize disabled node.
	{
		Name: "DISAB-1",
		Metadata: ruleMetadata{
			StepApplicability:      LessThan(5),
			AdditionalWeight:       -1,
			FinalizesDisabledNodes: 1,
		},
		searchNearPrevIndex: []int{-1},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				x, y := c.Unwrap()
				return !g.IsNodeActive(c) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.FinalizeNode(applyAt[0])
		},
	},

	// Finalize two adjacent disabled nodes.
	{
		Name: "DISAB-2",
		Metadata: ruleMetadata{
			StepApplicability:      LessThan(3),
			AdditionalWeight:       -2,
			FinalizesDisabledNodes: 2,
		},
		searchNearPrevIndex: []int{-1, 0},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				x, y := c.Unwrap()
				return !g.IsNodeActive(c) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				x, y := c.Unwrap()
				return !g.IsNodeActive(c) && prevСoords[0].IsAdjacentTo(c) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.FinalizeNode(applyAt[0])
			g.FinalizeNode(applyAt[1])
		},
	},

	// Finalize three adjacent disabled nodes. Prevents a node from being locked by a finalized L-shape in a corner.
	{
		Name: "DISAB-3",
		Metadata: ruleMetadata{
			StepApplicability:      LessThan(3),
			AdditionalWeight:       -7,
			FinalizesDisabledNodes: 3,
		},
		searchNearPrevIndex: []int{-1, 0, 1},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				x, y := c.Unwrap()
				return !g.IsNodeActive(c) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				x, y := c.Unwrap()
				return !g.IsNodeActive(c) && prevСoords[0].IsAdjacentTo(c) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
			// node 2
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				x, y := c.Unwrap()
				w, h := g.GetSize()
				// Prevent a not yet used node from being locked by a finalized L-shape in a corner.
				// Removal of this check causes a creation of unfillable nodes.
				if prevСoords[0].IsAdjacentToRectangleCorner(0, 0, w, h) && !prevСoords[1].IsOnRectangle(0, 0, w, h) {
					if c.IsAdjacentToRectangleCorner(0, 0, w, h) {
						return false
					}
				}
				return !g.IsNodeActive(c) && prevСoords[1].IsAdjacentTo(c) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.FinalizeNode(applyAt[0])
			g.FinalizeNode(applyAt[1])
			g.FinalizeNode(applyAt[2])
		},
	},

	// 0   1       0 > 1  ; where both are active
	{
		Name: "CONNECT",
		Metadata: ruleMetadata{
			AddsCycle:        true, // it's not guaranteed, but should be more possible than not
			AdditionalWeight: -3,
		},
		searchNearPrevIndex: []int{-1, 0},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentTo(c) && g.IsNodeActive(c) && !g.AreCoordsLinked(c, prevСoords[0])
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableDirLinkByCoords(applyAt[0], applyAt[1])
		},
		MandatoryFeatures: []*FeatureAdder{
			makeKeyLockFeature(0, 1),
			makeMasterKeyLockFeature(0, 1),
			makeSecretPassageFeature(0, 1),
			makeWindowFeature(0, 1),
			makeOneTimePassageFeature(0, 1),
			makeOneWayPassagesFeature(0, 1, 0, 1), // repeat on purpose
		},
	},

	// 1   3       1 > 3
	// V       >       V
	// 0 > 2       U   2
	{
		Name: "L-FLIP",
		Metadata: ruleMetadata{
			EnablesNodes:     0, // Enables 1 node, disables 1 node -> 0 in general
			AdditionalWeight: -2,
		},
		searchNearPrevIndex: []int{-1, 0, 0, 1},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				x, y := c.Unwrap()
				return g.IsNodeActive(c) && g.CountEdgesAtXY(x, y) == 2 && !g.NodeAt(c).IsFlagged()
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentTo(c) && g.IsNodeActive(c) && g.IsEdgeDirectedFromCoords(c, prevСoords[0])
			},
			// node 2
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentTo(c) && g.IsNodeActive(c) && g.IsEdgeDirectedFromCoords(prevСoords[0], c)
			},
			// node 3
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) && prevСoords[1].IsAdjacentTo(c) && prevСoords[2].IsAdjacentTo(c)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNode(applyAt[3].Unwrap())
			g.EnableDirLinkByCoords(applyAt[1], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[2])
			g.DisableDirLinkByCoords(applyAt[1], applyAt[0])
			g.DisableDirLinkByCoords(applyAt[0], applyAt[2])
			g.SwapNodeTags(applyAt[3], applyAt[0])
			g.NodeAt(applyAt[3]).MarkFlagged() // so there will be no need to finalize the node 0
			g.CopyEdgeTagsPreservingIds(applyAt[1], applyAt[0], applyAt[1], applyAt[3])
			g.CopyEdgeTagsPreservingIds(applyAt[0], applyAt[2], applyAt[3], applyAt[2])
			g.ResetNodeAndConnections(applyAt[0])
		},
	},
}
