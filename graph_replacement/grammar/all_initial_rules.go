package grammar

import (
	"cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

var AllInitialRules = []InitialRule{
	// U U    R-R
	//     >  | |
	// U U    R-R
	// random start and goal
	{
		Name:      "PL 4-CYCLE",
		AddsCycle: true,
		IsApplicableAt: func(g *Graph, x, y, _, _ int) bool {
			if !g.AreCoordsInBounds(x+1, y+1) {
				return false
			}
			return true
		},
		ApplyOnGraphAt: func(g *Graph, x, y, _, _ int) {
			g.DrawConnectedDirectionalRect(x, y, 2, 2, rnd.OneChanceFrom(2))

			sx, sy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j)
			})
			g.AddNodeTag(sx, sy, TagStart)
			gx, gy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j) && !g.DoesNodeHaveAnyTags(i, j)
			})
			g.AddNodeTag(gx, gy, TagGoal)
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
			if !g.AreCoordsInBounds(x+2, y+2) {
				return false
			}
			return true
		},
		ApplyOnGraphAt: func(g *Graph, x, y, _, _ int) {
			g.DrawConnectedDirectionalRect(x, y, 3, 3, rnd.OneChanceFrom(2))

			sx, sy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j)
			})
			g.AddNodeTag(sx, sy, TagStart)
			gx, gy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return g.IsNodeActive(i, j) && !g.DoesNodeHaveAnyTags(i, j) &&
					areCoordsAdjacent(i, j, sx, sy) && g.IsEdgeDirectedBetweenCoords(i, j, sx, sy)
			})
			g.AddNodeTag(gx, gy, TagGoal)
			g.AddEdgeTagByVector(sx, sy, gx-sx, gy-sy, TagLockedEdge)
			g.FinalizeNode(geometry.NewCoords(x+1, y+1))
		},
	},
	// random start, two paths to goal
	{
		Name:      "TWO WAYS",
		AddsCycle: true,
		IsApplicableAt: func(g *Graph, x, y, _, _ int) bool {
			if !g.AreCoordsInBounds(x+2, y+2) {
				return false
			}
			return true
		},
		ApplyOnGraphAt: func(g *Graph, x, y, _, _ int) {
			sx, sy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return areCoordsOnRectangle(i, j, x, y, 3, 3)
			})
			gx, gy := g.GetRandomCoordsByFunc(func(i, j int) bool {
				return areCoordsOnRectangle(i, j, x, y, 3, 3) && i != sx && j != sy
			})
			g.DrawBiсonnectedDirectionalRect(x, y, 3, 3, sx, sy, gx, gy)
			g.AddNodeTag(sx, sy, TagStart)
			g.AddNodeTag(gx, gy, TagGoal)
			g.FinalizeNode(geometry.NewCoords(x+1, y+1))
		},
	},
}
