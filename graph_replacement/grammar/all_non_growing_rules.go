package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	"cycdg/graph_replacement/grid_graph/graph_element"
)

var allNonGrowingRules = []ReplacementRule{

	// 0  X  ; just finalize disabled node.
	{
		Name: "DISAB-1",
		Metadata: ruleMetadata{
			AdditionalWeight:       0,
			FinalizesDisabledNodes: 1,
		},
		searchNearPrevIndex: []int{-1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && g.HasNoFinalizedNodesNearXY(x, y, true)
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
			AdditionalWeight:       -2,
			FinalizesDisabledNodes: 2,
		},
		searchNearPrevIndex: []int{-1, 0},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y) && g.HasNoFinalizedNodesNearXY(x, y, true)
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
			AdditionalWeight:       -7,
			FinalizesDisabledNodes: 3,
		},
		searchNearPrevIndex: []int{-1, 0, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				// Additional check:
				w, h := g.GetSize()
				x0, y0 := prevСoords[0].Unwrap()
				x1, y1 := prevСoords[1].Unwrap()
				// Prevent a not yet used node from being locked by a finalized L-shape in a corner.
				// Removal of this check causes a creation of unfillable nodes.
				if areCoordsAdjacentToRectangleCorner(x0, y0, 0, 0, w, h) && !areCoordsOnRectangle(x1, y1, 0, 0, w, h) {
					if areCoordsAdjacentToRectangleCorner(x, y, 0, 0, w, h) {
						return false
					}
				}
				return !g.IsNodeActive(x, y) && prevСoords[1].IsAdjacentToXY(x, y) && g.HasNoFinalizedNodesNearXY(x, y, true)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.FinalizeNode(applyAt[0])
			g.FinalizeNode(applyAt[1])
			g.FinalizeNode(applyAt[2])
		},
	},

	// WORKAROUND RULE
	// just unfinalize disabled node adjacent to an active one
	// Used as a workaround, so that DISABLE-rules won't "wall up" the nodes' growth.
	// {
	// 	Name:                    "~~UNDISABLE",
	// 	WorksWithFinalizedNodes: true,
	// 	Metadata: ruleMetadata{
	// 		AdditionalWeight:         -8,
	// 		UnfinalizesDisabledNodes: 1,
	// 	},
	// 	searchNearPrevIndex: []int{-1, 0, -1},
	// 	applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
	// 		// node 0
	// 		func(g *Graph, x, y int, prevСoords ...Coords) bool {
	// 			return !g.IsNodeActive(x, y) && g.IsNodeFinalized(x, y)
	// 		},
	// 		// node 1
	// 		func(g *Graph, x, y int, prevСoords ...Coords) bool {
	// 			return g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y)
	// 		},
	// 		// node 2 (not an actual node, just an applicability check)
	// 		func(g *Graph, x, y int, prevСoords ...Coords) bool {
	// 			return x == 0 && y == 0 && g.CountEmptyEditableNodesNearEnabledOnes() < 3
	// 		},
	// 	},
	// 	ApplyToGraph: func(g *Graph, applyAt ...Coords) {
	// 		g.UnsafeUnfinalizeNode(applyAt[0])
	// 	},
	// },

	// 0; just add someting to empty active node
	{
		Name: "THING",
		Metadata: ruleMetadata{
			AdditionalWeight: -4,
		},
		searchNearPrevIndex: []int{-1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y) && !g.DoesNodeHaveAnyTags(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {},
		MandatoryFeatures: []*FeatureAdder{
			{
				Name: "Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[0], graph_element.TagTreasure)
				},
			},
			makeRandomHazardFeature(0),
		},
	},

	// 0   1       0 > 1  ; where both are active
	{
		Name: "CONNECT",
		Metadata: ruleMetadata{
			AddsCycle: true, // it's not guaranteed, but should be more possible than not
		},
		searchNearPrevIndex: []int{-1, 0},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x1, y1) && g.IsNodeActive(x, y) && !g.AreCoordsInterlinked(x, y, x1, y1)
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
			EnablesNodes: 0, // Enables 1 node, disables 1 node -> 0 in general
		},
		searchNearPrevIndex: []int{-1, 0, 0, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y) && g.CountEdgesAt(x, y) == 2 && !g.NodeAt(x, y).IsFlagged()
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x0, y0 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x0, y0) && g.IsNodeActive(x, y) && g.IsEdgeDirectedBetweenCoords(x, y, x0, y0)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x0, y0 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x0, y0) && g.IsNodeActive(x, y) && g.IsEdgeDirectedBetweenCoords(x0, y0, x, y)
			},
			// node 3
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[1].IsAdjacentToXY(x, y) && prevСoords[2].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNode(applyAt[3][0], applyAt[3][1])
			g.EnableDirLinkByCoords(applyAt[1], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[2])
			g.DisableDirLinkByCoords(applyAt[1], applyAt[0])
			g.DisableDirLinkByCoords(applyAt[0], applyAt[2])
			g.SwapNodeTags(applyAt[3], applyAt[0])
			g.NodeAt(applyAt[3].Unwrap()).MarkFlagged() // so there will be no need to finalize the node 0
			g.CopyEdgeTagsPreservingIds(applyAt[1], applyAt[0], applyAt[1], applyAt[3])
			g.CopyEdgeTagsPreservingIds(applyAt[0], applyAt[2], applyAt[3], applyAt[2])
			g.ResetNodeAndConnections(applyAt[0])
		},
	},
}
