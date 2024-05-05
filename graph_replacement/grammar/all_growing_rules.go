package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	"cycdg/graph_replacement/grid_graph/graph_element"
)

var allGrowingRules = []ReplacementRule{
	// 0   1       0 > 1  ; where 1 is inactive
	{
		Name: "ADDNODE",
		Metadata: ruleMetadata{
			EnablesNodes: 1,
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
				return areCoordsAdjacent(x, y, x1, y1) && !g.IsNodeActive(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[1])
			g.EnableDirLinkByCoords(applyAt[0], applyAt[1])
			moveRandomNodeTag(g, applyAt[0], applyAt[1])
		},
		MandatoryFeatures: []*FeatureAdder{
			makeKeyLockFeature(0, 1),
			makeMasterKeyLockFeature(0, 1),
			makeSecretPassageFeature(0, 1),
			// makeOneTimePassageFeature(0, 1), // CAUSES UNPASSABLE MAPS TO CREATE (WITH THIS RULE)
		},
		OptionalFeatures: []*FeatureAdder{
			{
				Name: "Boss",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[1], graph_element.TagBoss)
				},
			},
			{
				Name: "Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[1], graph_element.TagTreasure)
				},
			},
		},
	},

	// 0 ...  1  2       0 (teleport)> 1 > 2  ; where 1 and 2 are inactive
	{
		Name: "TELEPORT",
		Metadata: ruleMetadata{
			AddsTeleport:     true,
			AdditionalWeight: -2,
			EnablesNodes:     2,
		},
		searchNearPrevIndex: []int{-1, -1, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y) && g.DoesNodeHaveAnyTags(x, y) &&
					!g.DoesNodeHaveTag(x, y, graph_element.TagTeleportBidir) &&
					!g.DoesNodeHaveTag(x, y, graph_element.TagStart)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[1].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[1])
			g.EnableNodeByCoords(applyAt[2])
			g.EnableDirLinkByCoords(applyAt[1], applyAt[2])
			moveRandomNodeTag(g, applyAt[0], applyAt[2])
			g.AddNodeTagByCoords(applyAt[0], graph_element.TagTeleportBidir)
			g.AddNodeTagByCoordsPreserveLastId(applyAt[1], graph_element.TagTeleportBidir)
		},
		OptionalFeatures: []*FeatureAdder{
			{
				Name: "Boss",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[2], graph_element.TagBoss)
				},
			},
			{
				Name: "Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[2], graph_element.TagTreasure)
				},
			},
		},
	},

	// 0   2       0 > 2
	// V       >       V
	// 1   3       1 < 3
	{
		Name: "U-RULE",
		Metadata: ruleMetadata{
			EnablesNodes: 2,
		},
		searchNearPrevIndex: []int{-1, 0, 0, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x1, y1) && g.IsNodeActive(x, y) && g.IsEdgeDirectedBetweenCoords(x1, y1, x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[0].Unwrap()
				return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x1, y1)
			},
			// node 3
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x2, y2 := prevСoords[1].Unwrap()
				x3, y3 := prevСoords[2].Unwrap()
				return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x2, y2) && areCoordsAdjacent(x, y, x3, y3)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNode(applyAt[2][0], applyAt[2][1])
			g.EnableNode(applyAt[3][0], applyAt[3][1])
			g.EnableDirLinkByCoords(applyAt[0], applyAt[2])
			g.EnableDirLinkByCoords(applyAt[2], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[1])
			g.DisableDirLinkByCoords(applyAt[0], applyAt[1])
		},
		MandatoryFeatures: []*FeatureAdder{
			{
				Name: "Swap 01-02",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.SwapEdgeTags(crds[0], crds[1], crds[0], crds[2])
				},
			},
			{
				Name: "Swap 01-23",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.SwapEdgeTags(crds[0], crds[1], crds[2], crds[3])
				},
			},
			{
				Name: "Swap 01-31",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.SwapEdgeTags(crds[0], crds[1], crds[3], crds[1])
				},
			},
		},
		OptionalFeatures: []*FeatureAdder{
			{
				Name: "Boss",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					ind := rnd.Rand(2) + 2
					g.AddNodeTagByCoords(crds[ind], graph_element.TagBoss)
				},
			},
			makeWindowFeature(0, 1),
		},
	},

	///////////////////////////////////////////////////////
	// EXPERIMENTAL RULES BELOW
	///////////////////////////////////////////////////////

	// Add a random straight line with the length at least of 3 and with return one-dir teleport at the end.
	// 0  X ... X   ->   0 > 1 > ... > 2
	{
		Name: "RND-LINE",
		Metadata: ruleMetadata{
			AddsTeleport:        true,
			EnablesNodesUnknown: true,
		},
		searchNearPrevIndex: []int{-1, 0, -1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0 - just a node near which the cycle will be appended
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y)
			},
			// node 2 - cardinal to 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[1].IsCardinalToPair(x, y) &&
					prevСoords[1].ManhattanDistToXY(x, y) >= 3 &&
					g.CheckFuncForAllNodesInCardinalLine(
						func(xc, yc int) bool {
							return !g.IsNodeActive(xc, yc) && !g.IsNodeFinalized(xc, yc)
						},
						x, y, prevСoords[1][0], prevСoords[1][1],
					)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.DrawEnabledConnectedCardinalLine(applyAt[1], applyAt[2])
			g.EnableDirLinkByCoords(applyAt[0], applyAt[1])
			g.AddNodeTagByCoords(applyAt[0], graph_element.TagTeleportTo)
			g.AddNodeTagByCoords(applyAt[2], graph_element.TagTeleportFrom)
		},
		MandatoryFeatures: []*FeatureAdder{
			makeMasterKeyLockFeature(0, 1),
			makeKeyLockFeature(0, 1),
		},
	},
}
