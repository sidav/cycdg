package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

// Funcs for features' auto-generation (just for DRYifying the rules descriptions)

func makeKeyLockFeature(lockBetweenIndex1, lockBetweenIndex2 int) *FeatureAdder {
	return &FeatureAdder{
		Name: "Locked",
		PrepareFeature: func(g *Graph, crds ...Coords) {
			addKeyAtRandom(g)
		},
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[lockBetweenIndex1], crds[lockBetweenIndex2], TagLockedEdge)
		},
	}
}

func makeOneKeyTwoLocksFeature(lock1ind1, lock1ind2, lock2ind1, lock2ind2 int) *FeatureAdder {
	return &FeatureAdder{
		Name: "2x Locked",
		PrepareFeature: func(g *Graph, crds ...Coords) {
			addKeyAtRandom(g)
		},
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[lock1ind1], crds[lock1ind2], TagLockedEdge)
			g.AddEdgeTagByCoordsPreserveLastId(crds[lock2ind1], crds[lock2ind2], TagLockedEdge)
		},
	}
}

func makeSecretPassageFeature(edgeIndex1, edgeIndex2 int) *FeatureAdder {
	return &FeatureAdder{
		Name: "Secret",
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[edgeIndex1], crds[edgeIndex2], TagSecretEdge)
		},
	}
}

func makeOneWayPassagesFeature(edge1FromIndex, edge1ToIndex, edge2FromIndex, edge2ToIndex int) *FeatureAdder {
	return &FeatureAdder{
		Name: "2x One-way",
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[edge1FromIndex], crds[edge1ToIndex], TagOneWayEdge)
			g.AddEdgeTagByCoords(crds[edge2FromIndex], crds[edge2ToIndex], TagOneWayEdge)
		},
	}
}

func makeOneTimePassageFeature(edgeIndex1, edgeIndex2 int) *FeatureAdder {
	return &FeatureAdder{
		Name: "One-time",
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[edgeIndex1], crds[edgeIndex2], TagOneTimeEdge)
		},
	}
}
