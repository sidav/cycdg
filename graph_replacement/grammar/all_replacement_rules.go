package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	"cycdg/graph_replacement/grid_graph/graph_element"
)

var AllReplacementRules = []*ReplacementRule{

	// 0  X  ; just finalize disabled node
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
				return !g.IsNodeActive(x, y) && !g.IsNodeFinalized(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.FinalizeNode(applyAt[0])
		},
	},

	// Disable two neighbouring nodes
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
				return !g.IsNodeActive(x, y) && !g.IsNodeFinalized(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && !g.IsNodeFinalized(x, y) && prevСoords[0].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.FinalizeNode(applyAt[0])
			g.FinalizeNode(applyAt[1])
		},
	},

	// Disable three neighbouring nodes
	{
		Name: "DISAB-3",
		Metadata: ruleMetadata{
			AdditionalWeight:       -4,
			FinalizesDisabledNodes: 3,
		},
		searchNearPrevIndex: []int{-1, 0, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && !g.IsNodeFinalized(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && !g.IsNodeFinalized(x, y) && prevСoords[0].IsAdjacentToXY(x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && !g.IsNodeFinalized(x, y) && prevСoords[1].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.FinalizeNode(applyAt[0])
			g.FinalizeNode(applyAt[1])
			g.FinalizeNode(applyAt[2])
		},
	},

	// 0; just add someting to empty active node
	{
		Name: "THING",
		Metadata: ruleMetadata{
			AdditionalWeight: 2,
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
			AdditionalWeight: -5,
			EnablesNodes:     2,
		},
		searchNearPrevIndex: []int{-1, -1, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y) && g.DoesNodeHaveAnyTags(x, y) &&
					!g.DoesNodeHaveTag(x, y, graph_element.TagTeleportBidirectional) &&
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
			g.AddNodeTagByCoords(applyAt[0], graph_element.TagTeleportBidirectional)
			g.AddNodeTagByCoordsPreserveLastId(applyAt[1], graph_element.TagTeleportBidirectional)
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

	// 0   2       0 < 2
	// V       >   V   ^
	// 1   3       1 > 3
	// {
	// 	Name:                "RET-LOOP",
	// 	Metadata: ruleMetadata{
	//		AddsCycle: true,
	//	},
	// 	searchNearPrevIndex: []int{-1, 0, 0, 1},
	// 	applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
	// 		// node 0
	// 		func(g *Graph, x, y int, prevСoords ...Coords) bool {
	// 			return g.IsNodeActive(x, y)
	// 		},
	// 		// node 1
	// 		func(g *Graph, x, y int, prevСoords ...Coords) bool {
	// 			x1, y1 := prevСoords[0].Unwrap()
	// 			return areCoordsAdjacent(x, y, x1, y1) && g.IsNodeActive(x, y) && g.IsEdgeDirectedBetweenCoords(x1, y1, x, y)
	// 		},
	// 		// node 2
	// 		func(g *Graph, x, y int, prevСoords ...Coords) bool {
	// 			x1, y1 := prevСoords[0].Unwrap()
	// 			return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x1, y1)
	// 		},
	// 		// node 3
	// 		func(g *Graph, x, y int, prevСoords ...Coords) bool {
	// 			x2, y2 := prevСoords[1].Unwrap()
	// 			x3, y3 := prevСoords[2].Unwrap()
	// 			return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x2, y2) && areCoordsAdjacent(x, y, x3, y3)
	// 		},
	// 	},
	// 	ApplyToGraph: func(g *Graph, applyAt ...Coords) {
	// 		g.EnableNode(applyAt[2][0], applyAt[2][1])
	// 		g.EnableNode(applyAt[3][0], applyAt[3][1])
	// 		g.EnableDirectionalLinkBetweenCoords(applyAt[2], applyAt[0])
	// 		g.EnableDirectionalLinkBetweenCoords(applyAt[3], applyAt[2])
	// 		g.EnableDirectionalLinkBetweenCoords(applyAt[1], applyAt[3])
	// 	},
	// 	MandatoryFeatures: []*FeatureAdder{
	// 		makeTwoMasterKeyLocksFeature(1, 3, 2, 0),
	// 		makeOneWayPassagesFeature(1, 3, 2, 0),
	// 	},
	// 	OptionalFeatures: []*FeatureAdder{
	// 		makeRandomHazardFeature(3),
	// 	},
	// },

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
				return g.IsNodeActive(x, y) && g.CountEdgesAt(x, y) == 2
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
				x1, y1 := prevСoords[1].Unwrap()
				x2, y2 := prevСoords[2].Unwrap()
				return !g.IsNodeActive(x, y) &&
					areCoordsAdjacent(x, y, x1, y1) && areCoordsAdjacent(x, y, x2, y2)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNode(applyAt[3][0], applyAt[3][1])
			g.EnableDirLinkByCoords(applyAt[1], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[2])
			g.DisableDirLinkByCoords(applyAt[1], applyAt[0])
			g.DisableDirLinkByCoords(applyAt[0], applyAt[2])
			g.SwapNodeTags(applyAt[3], applyAt[0])
			g.CopyEdgeTagsPreservingIds(applyAt[1], applyAt[0], applyAt[1], applyAt[3])
			g.CopyEdgeTagsPreservingIds(applyAt[0], applyAt[2], applyAt[3], applyAt[2])
			g.ResetNodeAndConnections(applyAt[0])
			// g.FinalizeNode(applyAt[0]) disabled so that it won't affect fill percentage.
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
}
