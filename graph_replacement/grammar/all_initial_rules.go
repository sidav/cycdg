package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

var AllInitialRules = []InitialRule{
	// random start and non-adjacent goal, biconnected, random size
	{
		Name:      "nAj-CYCLE",
		AddsCycle: true,
		IsApplicableAt: func(g *Graph, x, y int) bool {
			if !g.AreCoordsInBounds(x+3, y+3) || x == 0 || y == 0 {
				return false
			}
			return true
		},
		ApplyOnGraphAt: func(g *Graph, x, y int) {
			w, h := g.GetSize()
			rw, rh := rnd.RandInRange(3, w-x-1), rnd.RandInRange(3, h-y-1)
			start := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
				return areCoordsOnRectangle(i, j, x, y, rw, rh)
			})
			goal := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
				return areCoordsOnRectangle(i, j, x, y, rw, rh) && !start.EqualsPair(i, j) && start.ManhattanDistToXY(i, j) >= min(rw, rh)
			})
			g.DrawBiсonnectedDirectionalRect(x, y, rw, rh, start, goal)
			g.AddNodeTagByCoords(start, TagStart)
			g.AddNodeTagByCoords(goal, TagGoal)
		},
		Features: []*FeatureAdder{
			{
				Name: "Alt paths w hazards",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					goalCrd := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.DoesNodeHaveTag(i, j, TagGoal)
					})
					crds1 := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.IsNodeActive(i, j) && goalCrd.IsAdjacentToXY(i, j)
					})
					crds2 := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.IsNodeActive(i, j) && goalCrd.IsAdjacentToXY(i, j) && !crds1.EqualsPair(i, j)
					})
					AddRandomHazardAt(g, crds1)
					AddRandomHazardAt(g, crds2)
					if rnd.Rand(3) == 0 {
						PushNodeContentsInRandomDirection(g, goalCrd)
					}
				},
			},
			{
				Name: "Two keys",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					startCrd := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.DoesNodeHaveTag(i, j, TagStart)
					})
					goalCrd := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.DoesNodeHaveTag(i, j, TagGoal)
					})
					crds1 := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.IsNodeActive(i, j) && (goalCrd.IsAdjacentToXY(i, j) || startCrd.IsAdjacentToXY(i, j))
					})
					crds2 := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.IsNodeActive(i, j) && !crds1.EqualsPair(i, j) && !crds1.IsAdjacentToXY(i, j) &&
							(goalCrd.IsAdjacentToXY(i, j) || startCrd.IsAdjacentToXY(i, j))
					})
					g.AddNodeTagByCoords(crds1, TagKeyForEdge)
					g.AddNodeTagByCoordsPreserveLastId(crds2, TagKeyForEdge)
					PushNodeContentsInRandomDirectionWithEdgeTag(g, goalCrd, TagBilockedEdge)
				},
			},
		},
	},
	// random start and adjacent goal, biconnected, random size
	{
		Name:      "Aj-CYCLE",
		AddsCycle: true,
		IsApplicableAt: func(g *Graph, x, y int) bool {
			if x == 0 || y == 0 || !g.AreCoordsInBounds(x+2, y+2) {
				return false
			}
			return true
		},
		ApplyOnGraphAt: func(g *Graph, x, y int) {
			w, h := g.GetSize()
			rw, rh := rnd.RandInRange(3, w-x), rnd.RandInRange(3, h-y)
			start := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
				return areCoordsOnRectangle(i, j, x, y, rw, rh)
			})
			goal := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
				return areCoordsOnRectangle(i, j, x, y, rw, rh) && start.IsAdjacentToXY(i, j)
			})
			g.DrawBiсonnectedDirectionalRect(x, y, rw, rh, start, goal)
			g.AddNodeTagByCoords(start, TagStart)
			g.AddNodeTagByCoords(goal, TagGoal)
		},
		Features: []*FeatureAdder{
			{
				Name: "Foresee",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					startCrd := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.DoesNodeHaveTag(i, j, TagStart)
					})
					goalCrd := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.DoesNodeHaveTag(i, j, TagGoal)
					})
					g.AddEdgeTagByCoords(startCrd, goalCrd, TagWindowEdge)
				},
			},
			{
				Name: "OpenableShortcut",
				ApplyFeature: func(g *Graph, crds ...Coords) {
					startCrd := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.DoesNodeHaveTag(i, j, TagStart)
					})
					goalCrd := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.DoesNodeHaveTag(i, j, TagGoal)
					})
					crds1 := getRandomGraphCoordsByFunc(g, func(i, j int) bool {
						return g.IsNodeActive(i, j) && !startCrd.EqualsPair(i, j) && !startCrd.IsAdjacentToXY(i, j)
					})
					g.AddNodeTagByCoords(crds1, TagKeyForEdge)
					g.AddEdgeTagByCoords(startCrd, goalCrd, TagLockedEdge)
				},
			},
		},
	},
}
