package graph

var initialRules = []ReplacementRule{
	// U U    R-R
	//     >  | |
	// U U    R-R
	// random start and goal
	{
		Name:      "PL 4-CYCLE",
		AddsCycle: true,
		IsApplicableAt: func(g *Graph, x, y, _, _ int) bool {
			if !g.areCoordsInBounds(x+1, y+1) {
				return false
			}
			return true
		},
		ApplyOnGraphAt: func(g *Graph, x, y, _, _ int) {
			g.drawConnectedDirectionalRect(x, y, 2, 2, rnd.OneChanceFrom(2))

			sx, sy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j)
			})
			g.addNodeTag(sx, sy, "START")
			gx, gy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j) && !g.DoesNodeHaveAnyTags(i, j)
			})
			g.addNodeTag(gx, gy, "GOAL")
		},
	},
	// R-R-R
	// |   |
	// R   R
	// |   |
	// R-R-R
	// random start, neighbouring goal
	{
		Name:      "LONG WAY",
		AddsCycle: true,
		IsApplicableAt: func(g *Graph, x, y, _, _ int) bool {
			if !g.areCoordsInBounds(x+2, y+2) {
				return false
			}
			return true
		},
		ApplyOnGraphAt: func(g *Graph, x, y, _, _ int) {
			g.drawConnectedDirectionalRect(x, y, 3, 3, rnd.OneChanceFrom(2))

			sx, sy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j)
			})
			g.addNodeTag(sx, sy, "START")
			gx, gy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j) && !g.DoesNodeHaveAnyTags(i, j) &&
					areCoordsAdjacent(i, j, sx, sy) && g.IsEdgeDirectedBetweenCoords(i, j, sx, sy)
			})
			g.addNodeTag(gx, gy, "GOAL")
			g.addEdgeTagFromCoordsByVector(sx, sy, gx-sx, gy-sy, "LOCK")
			g.finalizeNode(x+1, y+1)
		},
	},
}
