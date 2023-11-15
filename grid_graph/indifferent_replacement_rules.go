package graph

import . "cycdg/grid_graph/geometry"

var allReplacementRules = []*indifferentRule{
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
		applyToCoords: func(g *Graph, allCoords ...Coords) {
			g.enableNode(allCoords[1][0], allCoords[1][1])
			g.enableDirectionalLinkBetweenCoords(allCoords[0], allCoords[1])
		},
	},

	// 0   2       0 > 2
	// V       >       V
	// 1   3       1 < 3
	{
		Name:                "U-RULE",
		searchNearPrevIndex: []int{-1, 0, 0, 1},
		applicabilityFuncs: []func(g *Graph, x, y int, prevСoords ...Coords) bool{
			// first node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// second node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x1, y1) && g.IsNodeActive(x, y) && g.IsEdgeDirectedBetweenCoords(x1, y1, x, y)
			},
			// third node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[0].Unwrap()
				return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x1, y1)
			},
			// fourth node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x2, y2 := prevСoords[1].Unwrap()
				x3, y3 := prevСoords[2].Unwrap()
				return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x2, y2) && areCoordsAdjacent(x, y, x3, y3)
			},
		},
		applyToCoords: func(g *Graph, allCoords ...Coords) {
			g.enableNode(allCoords[2][0], allCoords[2][1])
			g.enableNode(allCoords[3][0], allCoords[3][1])
			g.setLinkBetweenCoords(allCoords[0][0], allCoords[0][1], allCoords[1][0], allCoords[1][1], false)
			g.enableDirectionalLinkBetweenCoords(allCoords[0], allCoords[2])
			g.enableDirectionalLinkBetweenCoords(allCoords[2], allCoords[3])
			g.enableDirectionalLinkBetweenCoords(allCoords[3], allCoords[1])
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
			// first node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				return g.IsNodeActive(x, y)
			},
			// second node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[0].Unwrap()
				return areCoordsAdjacent(x, y, x1, y1) && g.IsNodeActive(x, y) && g.IsEdgeDirectedBetweenCoords(x1, y1, x, y)
			},
			// third node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x1, y1 := prevСoords[0].Unwrap()
				return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x1, y1)
			},
			// fourth node
			func(g *Graph, x, y int, prevСoords ...Coords) bool {
				x2, y2 := prevСoords[1].Unwrap()
				x3, y3 := prevСoords[2].Unwrap()
				return !g.IsNodeActive(x, y) && areCoordsAdjacent(x, y, x2, y2) && areCoordsAdjacent(x, y, x3, y3)
			},
		},
		applyToCoords: func(g *Graph, allCoords ...Coords) {
			g.enableNode(allCoords[2][0], allCoords[2][1])
			g.enableNode(allCoords[3][0], allCoords[3][1])
			g.enableDirectionalLinkBetweenCoords(allCoords[0], allCoords[2])
			g.enableDirectionalLinkBetweenCoords(allCoords[2], allCoords[3])
			g.enableDirectionalLinkBetweenCoords(allCoords[3], allCoords[1])
		},
	},

	// 1   3       1 > 3
	// V       >       V
	// 0 > 2       U   2
	{
		Name:                "FLIP",
		AddsCycle:           true,
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
		applyToCoords: func(g *Graph, allCoords ...Coords) {
			g.enableNode(allCoords[3][0], allCoords[3][1])
			g.enableDirectionalLinkBetweenCoords(allCoords[1], allCoords[3])
			g.enableDirectionalLinkBetweenCoords(allCoords[3], allCoords[2])
			g.disableDirectionalLinkBetweenCoords(allCoords[1], allCoords[0])
			g.disableDirectionalLinkBetweenCoords(allCoords[0], allCoords[2])
			g.SwapTagsAtCoords(allCoords[3][0], allCoords[3][1], allCoords[0][0], allCoords[0][1])
			g.resetNodeAndConnections(allCoords[0][0], allCoords[0][1])
			g.FinalizeNode(allCoords[0])
		},
	},
	// 0   1       0 - 1
	//         >   |   |    0 is active, others not
	// 2   3       2 - 3
	//
	// !! N has no other connections !!
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
		applyToCoords: func(g *Graph, allCoords ...Coords) {
			g.enableNode(allCoords[1].Unwrap())
			g.enableNode(allCoords[2].Unwrap())
			g.enableNode(allCoords[3].Unwrap())
			g.enableDirectionalLinkBetweenCoords(allCoords[0], allCoords[1])
			g.enableDirectionalLinkBetweenCoords(allCoords[1], allCoords[3])
			g.enableDirectionalLinkBetweenCoords(allCoords[3], allCoords[2])
			g.enableDirectionalLinkBetweenCoords(allCoords[2], allCoords[0])
		},
	},
}
