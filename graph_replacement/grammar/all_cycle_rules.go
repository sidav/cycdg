package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	"cycdg/graph_replacement/grid_graph/graph_element"
)

var allCycleRules = []*ReplacementRule{
	// 0   2   1       0 > 2 > 1 ; where 0 and 1 are active; may be bent
	{
		Name: "CONNROOM",
		Metadata: ruleMetadata{
			AddsCycle:    true, // it's not guaranteed, but should be more possible than not
			EnablesNodes: 1,
		},
		searchNearPrevIndex: []int{-1, -1, 0},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y) && !prevСoords[0].IsAdjacentToXY(x, y) // && prevСoords[0].IsCardinalToPair(x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y) && prevСoords[1].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[2])
			g.EnableDirLinkByCoords(applyAt[0], applyAt[2])
			g.EnableDirLinkByCoords(applyAt[2], applyAt[1])
		},
		MandatoryFeatures: []*FeatureAdder{
			makeOneTimePassageFeature(0, 2),
			makeMasterKeyLockFeature(0, 2),
			makeOneWayPassagesFeature(0, 2, 2, 1),
			makeTwoMasterKeyLocksFeature(0, 2, 2, 1),
			{
				Name: "SecretPassage",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddEdgeTagByCoords(crds[0], crds[2], graph_element.TagSecretEdge)
					g.AddEdgeTagByCoords(crds[2], crds[1], graph_element.TagSecretEdge)
					if rnd.Rand(2) == 0 {
						AddRandomHazardAt(g, crds[2])
					}
				},
			},
		},
		OptionalFeatures: []*FeatureAdder{
			makeRandomHazardFeature(2),
			makeMasterKeyLockFeature(2, 1),
			{
				Name: "Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[2], graph_element.TagTreasure)
				},
			},
		},
	},

	// 0   2       0 > 2
	// V       >   V   V
	// 1   3       1 < 3
	{
		Name: "D-RULE",
		Metadata: ruleMetadata{
			AddsCycle:    true,
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
		},
		MandatoryFeatures: []*FeatureAdder{
			{
				Name: "Copy 01-02",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.CopyEdgeTagsPreservingIds(crds[0], crds[1], crds[0], crds[2])
				},
			},
			makeSecretPassageFeature(0, 2),
			makeMasterKeyLockFeature(0, 2),
			makeTwoMasterKeyLocksFeature(0, 2, 3, 1),
			makeOneWayPassagesFeature(0, 2, 3, 1),
		},
		OptionalFeatures: []*FeatureAdder{
			makeRandomHazardFeature(3),
		},
	},

	// 0   1       0 > 1
	//         >   ^   V    0 is active, others not
	// 2   3       2 < 3
	{
		Name: "CORNERLOOP",
		Metadata: ruleMetadata{
			AddsCycle:    true,
			EnablesNodes: 3,
		},
		searchNearPrevIndex: []int{-1, 0, 0, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x0, y0 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x0, y0) && !g.IsNodeActive(x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x0, y0 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x0, y0) && !g.IsNodeActive(x, y)
			},
			// node 3
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[1].Unwrap()
				x2, y2 := prevСoords[2].Unwrap()
				return !g.IsNodeActive(x, y) &&
					areCoordsAdjacent(x, y, x1, y1) && areCoordsAdjacent(x, y, x2, y2)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNode(applyAt[1].Unwrap())
			g.EnableNode(applyAt[2].Unwrap())
			g.EnableNode(applyAt[3].Unwrap())
			g.EnableDirLinkByCoords(applyAt[0], applyAt[1])
			g.EnableDirLinkByCoords(applyAt[1], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[2])
			g.EnableDirLinkByCoords(applyAt[2], applyAt[0])
		},
		OptionalFeatures: []*FeatureAdder{
			makeOneWayPassagesFeature(0, 1, 2, 0),
			makeTwoMasterKeyLocksFeature(0, 1, 2, 0),
			makeMasterKeyLockFeature(0, 1),
			{
				Name: "SecretOrHazard",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[1], graph_element.TagBoss)
					g.AddNodeTagByCoords(crds[2], graph_element.TagTreasure)
					g.AddEdgeTagByCoords(crds[2], crds[0], graph_element.TagSecretEdge)
				},
			},
			{
				Name: "ForcedBoss",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[3], graph_element.TagBoss)
					g.AddEdgeTagByCoords(crds[0], crds[1], graph_element.TagOneTimeEdge)
					g.AddEdgeTagByCoords(crds[2], crds[0], graph_element.TagOneTimeEdge)
				},
			},
		},
	},

	// 4   3           4 < 3
	//             >  ~V   ^          0, 1, 2 are active, others not, ~> means locked path
	// 0 > 1 > 2       0 > 1 ~> 2
	{
		Name: "GOODLOCK+2",
		Metadata: ruleMetadata{
			AddsCycle:    true,
			EnablesNodes: 2,
		},
		searchNearPrevIndex: []int{-1, 0, 1, 0, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x0, y0 := prevСoords[0].Unwrap()
				return g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x0, y0) &&
					g.IsEdgeDirectedBetweenCoords(x0, y0, x, y) && g.DoesEdgeHaveZeroTags(prevСoords[0], NewCoords(x, y))
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[1].Unwrap()
				return g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x1, y1) && g.DoesNodeHaveAnyTags(x, y) &&
					g.IsEdgeDirectedBetweenCoords(x1, y1, x, y) && g.DoesEdgeHaveZeroTags(prevСoords[1], NewCoords(x, y))
			},
			// node 3
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[1].Unwrap()
				return !g.IsNodeActive(x, y) &&
					areCoordsAdjacent(x, y, x1, y1)
			},
			// node 4
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x0, y0 := prevСoords[0].Unwrap()
				return !g.IsNodeActive(x, y) &&
					areCoordsAdjacent(x, y, x0, y0) && prevСoords[3].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNode(applyAt[3].Unwrap())
			g.EnableNode(applyAt[4].Unwrap())
			g.EnableDirLinkByCoords(applyAt[1], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[4])
			g.EnableDirLinkByCoords(applyAt[4], applyAt[0])
			g.AddNodeTagByCoords(applyAt[4], graph_element.TagKey)
			g.AddEdgeTagByCoords(applyAt[1], applyAt[2], graph_element.TagLockedEdge)
		},
		MandatoryFeatures: []*FeatureAdder{
			nil,
			{
				Name: "Singleway",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddEdgeTagByCoords(crds[1], crds[3], graph_element.TagOneWayEdge)
					g.AddEdgeTagByCoordsPreserveLastId(crds[4], crds[0], graph_element.TagLockedEdge)
				},
			},
			{
				Name: "Window",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddEdgeTagByCoordsPreserveLastId(crds[4], crds[0], graph_element.TagWindowEdge)
				},
			},
		},
		OptionalFeatures: []*FeatureAdder{
			{
				Name: "BossGuardsKey",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[4], graph_element.TagBoss)
				},
			},
			{
				Name: "Ambush",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[4], graph_element.TagTrap)
				},
			},
		},
	},

	// STRAIGHT EXAMPLE
	// 0 > 1 > 2       0 > 1 > 2
	//             >   V       ^   0-2 are active, others not, may be bent
	// 3   4   5       3 > 4 > 5
	//
	// BENT EXAMPLE:
	// 0   3   4       0 > 3 > 4
	// V           >   V       V
	// 1 > 2   5       1 > 2 < 5
	{
		Name: "ALTWAY",
		Metadata: ruleMetadata{
			AddsCycle:    true,
			EnablesNodes: 3,
		},
		searchNearPrevIndex: []int{-1, 0, 1, 0, 3, 2},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentToXY(x, y) && g.IsNodeActive(x, y) &&
					g.IsEdgeDirectedBetweenCoords(prevСoords[0][0], prevСoords[0][1], x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[1].IsAdjacentToXY(x, y) && g.IsNodeActive(x, y) &&
					g.IsEdgeDirectedBetweenCoords(prevСoords[1][0], prevСoords[1][1], x, y)
			},
			// node 3
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentToXY(x, y) && !g.IsNodeActive(x, y)
			},
			// node 4
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[3].IsAdjacentToXY(x, y) && !g.IsNodeActive(x, y)
			},
			// node 5
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[2].IsAdjacentToXY(x, y) && prevСoords[4].IsAdjacentToXY(x, y) && !g.IsNodeActive(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNode(applyAt[3].Unwrap())
			g.EnableNode(applyAt[4].Unwrap())
			g.EnableNode(applyAt[5].Unwrap())
			g.EnableDirLinkByCoords(applyAt[0], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[4])
			g.EnableDirLinkByCoords(applyAt[4], applyAt[5])
			g.EnableDirLinkByCoords(applyAt[5], applyAt[2])
			if !g.DoesNodeHaveAnyTags(applyAt[1].Unwrap()) {
				AddRandomHazardAt(g, applyAt[1])
			}
		},
		MandatoryFeatures: []*FeatureAdder{
			nil,
			makeSecretPassageFeature(0, 3),
			makeMasterKeyLockFeature(0, 3),
			makeOneKeyTwoLocksFeature(0, 3, 5, 2),
			makeOneWayPassagesFeature(0, 3, 5, 2),
			makeTwoMasterKeyLocksFeature(0, 3, 5, 2),
			// {
			// 	Name: "OneTime",
			// 	ApplyFeature: func(g *Graph, crds ...Coords) {
			// 		g.AddEdgeTagByCoords(crds[0], crds[3], graph_element.TagOnetimeEdge)
			// 		g.AddEdgeTagByCoords(crds[5], crds[2], graph_element.TagOnetimeEdge)
			// 	},
			// },
		},
		OptionalFeatures: []*FeatureAdder{
			makeRandomHazardFeature(4),
		},
	},

	//  X   X   X     2 > 3 > 4
	//            >>  ^       V
	//  0 > 1   X     0 > 1 < 5
	{
		Name: "2ADJ-CYCL+4",
		Metadata: ruleMetadata{
			AddsCycle:    true,
			EnablesNodes: 4,
		},
		searchNearPrevIndex: []int{-1, 0, 0, 2, 3, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentToXY(x, y) && g.IsNodeActive(x, y) &&
					g.IsEdgeDirectedFromCoordsToPair(prevСoords[0], x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y)
			},
			// node 3
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[2].IsAdjacentToXY(x, y)
			},
			// node 4
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[3].IsAdjacentToXY(x, y)
			},
			// node 5
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[1].IsAdjacentToXY(x, y) && prevСoords[4].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[2])
			g.EnableNodeByCoords(applyAt[3])
			g.EnableNodeByCoords(applyAt[4])
			g.EnableNodeByCoords(applyAt[5])
			g.EnableDirLinkByCoords(applyAt[0], applyAt[2])
			g.EnableDirLinkByCoords(applyAt[2], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[4])
			g.EnableDirLinkByCoords(applyAt[4], applyAt[5])
			g.EnableDirLinkByCoords(applyAt[5], applyAt[1])
		},
		MandatoryFeatures: []*FeatureAdder{
			{
				Name: "Copy 01-02",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.CopyEdgeTagsPreservingIds(crds[0], crds[1], crds[0], crds[2])
				},
			},
			makeSecretPassageFeature(0, 2),
			makeMasterKeyLockFeature(0, 2),
			makeTwoMasterKeyLocksFeature(0, 2, 5, 1),
			makeOneWayPassagesFeature(0, 2, 5, 1),
		},
		OptionalFeatures: []*FeatureAdder{
			makeRandomHazardFeature(2),
		},
	},

	//  X   X   X   X     3 > 4 > 5 > 6
	//                >>  ^           V
	//  0 > 1 > 2   X     0 > 1 > 2 < 7
	{
		Name: "3ADJ-CYCL+5",
		Metadata: ruleMetadata{
			AddsCycle:        true,
			EnablesNodes:     5,
			AdditionalWeight: -2,
		},
		searchNearPrevIndex: []int{-1, 0, 1, 0, 3, 4, 5, 2},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentToXY(x, y) && g.IsNodeActive(x, y) &&
					g.IsEdgeDirectedFromCoordsToPair(prevСoords[0], x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return prevСoords[1].IsAdjacentToXY(x, y) && g.IsNodeActive(x, y) &&
					g.IsEdgeDirectedFromCoordsToPair(prevСoords[1], x, y)
			},
			// node 3
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y)
			},
			// node 4
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[3].IsAdjacentToXY(x, y)
			},
			// node 5
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[4].IsAdjacentToXY(x, y)
			},
			// node 6
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[5].IsAdjacentToXY(x, y)
			},
			// node 7
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[2].IsAdjacentToXY(x, y) && prevСoords[6].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[3])
			g.EnableNodeByCoords(applyAt[4])
			g.EnableNodeByCoords(applyAt[5])
			g.EnableNodeByCoords(applyAt[6])
			g.EnableNodeByCoords(applyAt[7])
			g.EnableDirLinkByCoords(applyAt[0], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[4])
			g.EnableDirLinkByCoords(applyAt[4], applyAt[5])
			g.EnableDirLinkByCoords(applyAt[5], applyAt[6])
			g.EnableDirLinkByCoords(applyAt[6], applyAt[7])
			g.EnableDirLinkByCoords(applyAt[7], applyAt[2])
		},
	},

	///////////////////////////////////////////////////////
	// EXPERIMENTAL RULES BELOW
	///////////////////////////////////////////////////////

	// Add a random adjacent to 0 cycle, size is at least 3x3
	// 0             0 > 1 > ... > 2
	//     ->            V         V
	//                  ...       ...
	//                   V         V
	//                   3 > ... > 4
	{
		Name: "RND-ADJ-CYCL",
		Metadata: ruleMetadata{
			AddsCycle:           true,
			EnablesNodesUnknown: true,
		},
		searchNearPrevIndex: []int{-1, 0, -1, -1, -1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0 - just a node near which the cycle will be appended
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1 - corners
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y)
			},
			// node 2 - cardinal to 1 corner
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
			// node 3 - another cardinal to 1 corner, should NOT be cardinal to 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[1].IsCardinalToPair(x, y) &&
					prevСoords[1].ManhattanDistToXY(x, y) >= 3 && !prevСoords[2].IsCardinalToPair(x, y) &&
					g.CheckFuncForAllNodesInCardinalLine(
						func(xc, yc int) bool {
							return !g.IsNodeActive(xc, yc) && !g.IsNodeFinalized(xc, yc)
						},
						x, y, prevСoords[1][0], prevСoords[1][1],
					)
			},
			// node 4 - cardinal to both 2 and 3, diaginal (NOT cardinal) to 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[2].IsCardinalToPair(x, y) &&
					prevСoords[3].IsCardinalToPair(x, y) && !prevСoords[1].IsCardinalToPair(x, y) &&
					g.CheckFuncForAllNodesInCardinalLine(
						func(xc, yc int) bool {
							return !g.IsNodeActive(xc, yc) && !g.IsNodeFinalized(xc, yc)
						},
						x, y, prevСoords[2][0], prevСoords[2][1],
					) &&
					g.CheckFuncForAllNodesInCardinalLine(
						func(xc, yc int) bool {
							return !g.IsNodeActive(xc, yc) && !g.IsNodeFinalized(xc, yc)
						},
						x, y, prevСoords[3][0], prevСoords[3][1],
					)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			x, y, w, h := applyAt[1].GetRectangleForAnotherCornerCoords(applyAt[4])
			g.DrawBiсonnectedDirectionalRect(x, y, w, h, applyAt[1], applyAt[4])
			g.EnableDirLinkByCoords(applyAt[0], applyAt[1])
		},
	},
}
