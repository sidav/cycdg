package graph

import . "cycdg/grid_graph/geometry"

var allReplacementRules = []*indifferentRule{
	// 0   2       0 > 2
	// V       >       V
	// 1   3       1 < 3
	{
		Name: "U-rule",
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
			g.enableDirectionalLinkBetweenCoords(allCoords[0][0], allCoords[0][1], allCoords[2][0], allCoords[2][1])
			g.enableDirectionalLinkBetweenCoords(allCoords[2][0], allCoords[2][1], allCoords[3][0], allCoords[3][1])
			g.enableDirectionalLinkBetweenCoords(allCoords[3][0], allCoords[3][1], allCoords[1][0], allCoords[1][1])
		},
	},
}
