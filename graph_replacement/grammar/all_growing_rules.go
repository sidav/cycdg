package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

var allGrowingRules = []*ReplacementRule{
	// 0   1       0 > 1  ; where 1 is inactive
	{
		Name: "ADDNODE",
		Metadata: ruleMetadata{
			EnablesNodes:     1,
			AdditionalWeight: -2,
		},
		searchNearPrevIndex: []int{-1, 0},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentTo(c) && !g.IsNodeActive(c)
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
					g.AddNodeTagByCoords(crds[1], TagBoss)
				},
			},
			{
				Name: "Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[1], TagTreasure)
				},
			},
		},
	},

	// 0 ...  1  2       0 (teleport)> 1 > 2  ; where 1 and 2 are inactive
	{
		Name: "TELEPORT",
		Metadata: ruleMetadata{
			AddsTeleport:     true,
			AdditionalWeight: -3,
			EnablesNodes:     2,
		},
		searchNearPrevIndex: []int{-1, -1, 1},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c) && !g.DoesNodeHaveAnyTags(c) &&
					!g.DoesNodeHaveTag(c, TagTeleportBidir) &&
					!g.DoesNodeHaveTag(c, TagStart)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c)
			},
			// node 2
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) && prevСoords[1].IsAdjacentTo(c)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[1])
			g.EnableNodeByCoords(applyAt[2])
			g.EnableDirLinkByCoords(applyAt[1], applyAt[2])
			moveRandomNodeTag(g, applyAt[0], applyAt[2])
			g.AddNodeTagByCoords(applyAt[0], TagTeleportBidir)
			g.AddNodeTagByCoordsPreserveLastId(applyAt[1], TagTeleportBidir)
		},
		OptionalFeatures: []*FeatureAdder{
			{
				Name: "Boss",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[2], TagBoss)
				},
			},
			{
				Name: "Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[2], TagTreasure)
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
			EnablesNodes:      2,
		},
		searchNearPrevIndex: []int{-1, 0, 0, 1},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentTo(c) && g.IsNodeActive(c) && g.IsEdgeDirectedFromCoords(prevСoords[0], c)
			},
			// node 2
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) && prevСoords[0].IsAdjacentTo(c)
			},
			// node 3
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) && prevСoords[1].IsAdjacentTo(c) && prevСoords[2].IsAdjacentTo(c)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[2])
			g.EnableNodeByCoords(applyAt[3])
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
					g.AddNodeTagByCoords(crds[ind], TagBoss)
				},
			},
			makeWindowFeature(0, 1),
		},
	},
}
