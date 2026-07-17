package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

var allCycleRules = []*ReplacementRule{
	// 0   X   1       0 > 2 > 1 ; where 0 and 1 are active; may be bent
	{
		Name: "CONNROOM",
		Metadata: ruleMetadata{
			AddsCycle:         true, // it's not guaranteed, but should be more possible than not
			EnablesNodes:      1,
		},
		searchNearPrevIndex: []int{-1, -1, 0},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c) && !prevСoords[0].IsAdjacentTo(c) // && prevСoords[0].IsCardinalToPair(x, y)
			},
			// node 2
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) && prevСoords[0].IsAdjacentTo(c) && prevСoords[1].IsAdjacentTo(c)
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
					g.AddEdgeTagByCoords(crds[0], crds[2], TagSecretEdge)
					g.AddEdgeTagByCoords(crds[2], crds[1], TagSecretEdge)
					if rnd.Rand(2) == 0 {
						AddRandomHazardAt(g, crds[2])
					}
				},
			},
		},
		OptionalFeatures: []*FeatureAdder{
			makeRandomHazardFeature(2),
			{
				Name: "Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[2], TagTreasure)
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
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentTo(c) && !g.IsNodeActive(c)
			},
			// node 2
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return prevСoords[0].IsAdjacentTo(c) && !g.IsNodeActive(c)
			},
			// node 3
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) &&
					prevСoords[1].IsAdjacentTo(c) && prevСoords[2].IsAdjacentTo(c)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[1])
			g.EnableNodeByCoords(applyAt[2])
			g.EnableNodeByCoords(applyAt[3])
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
					g.AddNodeTagByCoords(crds[1], TagBoss)
					g.AddNodeTagByCoords(crds[2], TagTreasure)
					g.AddEdgeTagByCoords(crds[2], crds[0], TagSecretEdge)
				},
			},
			{
				Name: "ForcedBoss",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[3], TagBoss)
					g.AddEdgeTagByCoords(crds[0], crds[1], TagOneTimeEdge)
					g.AddEdgeTagByCoords(crds[2], crds[0], TagOneTimeEdge)
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
			AddsCycle:        true,
			EnablesNodes:     2,
		},
		searchNearPrevIndex: []int{-1, 0, 1, 0, 1},
		applicabilityFuncs: []func(g *Graph, c Coords, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c)
			},
			// node 1
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c) && prevСoords[0].IsAdjacentTo(c) &&
					g.IsEdgeDirectedFromCoords(prevСoords[0], c) && g.DoesEdgeHaveZeroTags(prevСoords[0], c)
			},
			// node 2
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return g.IsNodeActive(c) && prevСoords[1].IsAdjacentTo(c) && g.DoesNodeHaveAnyTags(c) &&
					g.IsEdgeDirectedFromCoords(prevСoords[1], c) && g.DoesEdgeHaveZeroTags(prevСoords[1], c)
			},
			// node 3
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) && prevСoords[1].IsAdjacentTo(c)
			},
			// node 4
			func(g *Graph, c Coords, prevСoords ...Coords) bool {
				return !g.IsNodeActive(c) && prevСoords[0].IsAdjacentTo(c) && prevСoords[3].IsAdjacentTo(c)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[3])
			g.EnableNodeByCoords(applyAt[4])
			g.EnableDirLinkByCoords(applyAt[1], applyAt[3])
			g.EnableDirLinkByCoords(applyAt[3], applyAt[4])
			g.EnableDirLinkByCoords(applyAt[4], applyAt[0])
			g.AddNodeTagByCoords(applyAt[4], TagKey)
			g.AddEdgeTagByCoords(applyAt[1], applyAt[2], TagLockedEdge)
		},
		MandatoryFeatures: []*FeatureAdder{
			nil,
			{
				Name: "Singleway",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddEdgeTagByCoords(crds[1], crds[3], TagOneWayEdge)
					g.AddEdgeTagByCoordsPreserveLastId(crds[4], crds[0], TagLockedEdge)
				},
			},
			{
				Name: "Window",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddEdgeTagByCoordsPreserveLastId(crds[4], crds[0], TagWindowEdge)
				},
			},
		},
		OptionalFeatures: []*FeatureAdder{
			{
				Name: "BossGuardsKey",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[4], TagBoss)
				},
			},
			{
				Name: "Ambush",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[4], TagTrap)
				},
			},
		},
	},
}
