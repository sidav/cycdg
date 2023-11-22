package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	"cycdg/graph_replacement/grid_graph/graph_element"
)

var AllReplacementRules = []*ReplacementRule{
	// 0   1       0 > 1  ; where 1 is inactive
	{
		Name:                "ADDNODE",
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
			g.EnableNode(applyAt[1][0], applyAt[1][1])
			g.EnableDirectionalLinkBetweenCoords(applyAt[0], applyAt[1])
		},
		Features: []*FeatureAdder{
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
			{
				Name: "Secret Treasure",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddEdgeTagByCoords(crds[0], crds[1], graph_element.TagSecretEdge)
					g.AddNodeTagByCoords(crds[1], graph_element.TagTreasure)
				},
			},
		},
	},

	// 0   2   1       0 > 2 > 1 ; where 0 and 1 are active; may be bent
	{
		Name:                "CONNECT",
		searchNearPrevIndex: []int{-1, -1, 0},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// node 0
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// node 1
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y) && !prevСoords[0].IsAdjacentToXY(x, y) && prevСoords[0].IsCardinalToPair(x, y)
			},
			// node 2
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return !g.IsNodeActive(x, y) && prevСoords[0].IsAdjacentToXY(x, y) && prevСoords[1].IsAdjacentToXY(x, y)
			},
		},
		ApplyToGraph: func(g *Graph, applyAt ...Coords) {
			g.EnableNodeByCoords(applyAt[2])
			g.EnableDirectionalLinkBetweenCoords(applyAt[0], applyAt[2])
			g.EnableDirectionalLinkBetweenCoords(applyAt[2], applyAt[1])
		},
		Features: []*FeatureAdder{
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
	},

	// 0   2       0 > 2
	// V       >       V
	// 1   3       1 < 3
	{
		Name:                "U-RULE",
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
			g.EnableDirectionalLinkBetweenCoords(applyAt[0], applyAt[2])
			g.EnableDirectionalLinkBetweenCoords(applyAt[2], applyAt[3])
			g.EnableDirectionalLinkBetweenCoords(applyAt[3], applyAt[1])
			g.SwapEdgeTags(applyAt[0], applyAt[1], applyAt[3], applyAt[1])
			g.SetLinkBetweenCoords(applyAt[0][0], applyAt[0][1], applyAt[1][0], applyAt[1][1], false)
		},
		Features: []*FeatureAdder{
			{
				Name: "Boss",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[3], graph_element.TagBoss)
				},
			},
		},
	},

	// 0   2       0 > 2
	// V       >   V   V
	// 1   3       1 < 3
	{
		Name:                "D-RULE",
		AddsCycle:           true,
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
			g.EnableDirectionalLinkBetweenCoords(applyAt[0], applyAt[2])
			g.EnableDirectionalLinkBetweenCoords(applyAt[2], applyAt[3])
			g.EnableDirectionalLinkBetweenCoords(applyAt[3], applyAt[1])
			g.CopyEdgeTagsPreservingIds(applyAt[0], applyAt[1], applyAt[0], applyAt[2])
		},
		Features: []*FeatureAdder{
			{
				Name: "SecretOrMaybeDanger",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddEdgeTagByCoords(crds[0], crds[1], graph_element.TagSecretEdge)
					if rnd.Rand(3) < 2 {
						AddRandomHazardAt(g, crds[3])
					}
				},
			},
		},
	},

	// 1   3       1 > 3
	// V       >       V
	// 0 > 2       U   2
	{
		Name:                "FLIP",
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
			g.EnableDirectionalLinkBetweenCoords(applyAt[1], applyAt[3])
			g.EnableDirectionalLinkBetweenCoords(applyAt[3], applyAt[2])
			g.DisableDirectionalLinkBetweenCoords(applyAt[1], applyAt[0])
			g.DisableDirectionalLinkBetweenCoords(applyAt[0], applyAt[2])
			g.SwapNodeTags(applyAt[3], applyAt[0])
			g.ResetNodeAndConnections(applyAt[0])
			g.FinalizeNode(applyAt[0])
		},
	},
	// 0   1       0 > 1
	//         >   ^   V    0 is active, others not
	// 2   3       2 < 3
	{
		Name:                "CORNERLOOP",
		AddsCycle:           true,
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
			g.EnableDirectionalLinkBetweenCoords(applyAt[0], applyAt[1])
			g.EnableDirectionalLinkBetweenCoords(applyAt[1], applyAt[3])
			g.EnableDirectionalLinkBetweenCoords(applyAt[3], applyAt[2])
			g.EnableDirectionalLinkBetweenCoords(applyAt[2], applyAt[0])
		},
		Features: []*FeatureAdder{
			{
				Name: "SecretOrHazard",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					g.AddNodeTagByCoords(crds[1], graph_element.TagBoss)
					g.AddNodeTagByCoords(crds[2], graph_element.TagTreasure)
					g.AddEdgeTagByCoords(crds[2], crds[0], graph_element.TagSecretEdge)
				},
			},
		},
	},
}
