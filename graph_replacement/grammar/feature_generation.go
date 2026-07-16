package grammar

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
)

// Funcs for features' auto-generation (just for DRYifying the rules descriptions)

func makeKeyLockFeature(lockBetweenIndex1, lockBetweenIndex2 int) *FeatureAdder {
	return &FeatureAdder{
		Name: "LockedRandom",
		PrepareFeature: func(g *Graph, crds ...Coords) {
			addTagAtRandomActiveNode(g, TagKey)
		},
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[lockBetweenIndex1], crds[lockBetweenIndex2], TagLockedEdge)
		},
	}
}

func makeKeyLockAtKnownCoordsFeature(lockBetweenIndex1, lockBetweenIndex2, keyX, keyY int) *FeatureAdder {
	return &FeatureAdder{
		Name: "LockedKnown",
		PrepareFeature: func(g *Graph, crds ...Coords) {
			g.AddNodeTag(keyX, keyY, TagKey)
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
			addTagAtRandomActiveNode(g, TagKey)
		},
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[lock1ind1], crds[lock1ind2], TagLockedEdge)
			g.AddEdgeTagByCoordsPreserveLastId(crds[lock2ind1], crds[lock2ind2], TagLockedEdge)
		},
	}
}

// Adds a master key ONLY if it's not on the map
func makeMasterKeyLockFeature(lockBetweenIndex1, lockBetweenIndex2 int) *FeatureAdder {
	return &FeatureAdder{
		Name: "Master-locked",
		PrepareFeature: func(g *Graph, crds ...Coords) {
			if !doesGraphContainNodeTag(g, TagMasterkey) {
				addTagAtRandomActiveNode(g, TagMasterkey)
			}
		},
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[lockBetweenIndex1], crds[lockBetweenIndex2], TagMasterLockedEdge)
		},
	}
}

func makeTwoMasterKeyLocksFeature(lock1ind1, lock1ind2, lock2ind1, lock2ind2 int) *FeatureAdder {
	return &FeatureAdder{
		Name:             "Master-locked",
		AdditionalWeight: -5,
		PrepareFeature: func(g *Graph, crds ...Coords) {
			if !doesGraphContainNodeTag(g, TagMasterkey) {
				addTagAtRandomActiveNode(g, TagMasterkey)
			}
		},
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[lock1ind1], crds[lock1ind2], TagMasterLockedEdge)
			g.AddEdgeTagByCoordsPreserveLastId(crds[lock2ind1], crds[lock2ind2], TagMasterLockedEdge)
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
		Name:             "2x One-way",
		AdditionalWeight: -8,
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[edge1FromIndex], crds[edge1ToIndex], TagOneWayEdge)
			g.AddEdgeTagByCoords(crds[edge2FromIndex], crds[edge2ToIndex], TagOneWayEdge)
		},
	}
}

func makeOneTimePassageFeature(edgeIndex1, edgeIndex2 int) *FeatureAdder {
	return &FeatureAdder{
		Name:             "One-time",
		AdditionalWeight: -7,
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.AddEdgeTagByCoords(crds[edgeIndex1], crds[edgeIndex2], TagOneTimeEdge)
		},
	}
}

func makeWindowFeature(edgeIndex1, edgeIndex2 int) *FeatureAdder {
	return &FeatureAdder{
		Name: "Window",
		ApplyFeature: func(g *Graph, crds ...Coords) {
			g.EnableDirLinkByCoords(crds[edgeIndex1], crds[edgeIndex2])
			g.AddEdgeTagByCoords(crds[edgeIndex1], crds[edgeIndex2], TagWindowEdge)
		},
	}
}

func makeRandomHazardFeature(edgeIndex int) *FeatureAdder {
	return &FeatureAdder{
		Name: "Random hazard",
		ApplyFeature: func(g *Graph, crds ...Coords) {
			AddRandomHazardAt(g, crds[edgeIndex])
		},
	}
}
