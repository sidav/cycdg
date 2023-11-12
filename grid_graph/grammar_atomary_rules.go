package graph

var atomaryRules = []ReplacementRule{
	// X U   >   X-N
	{
		Name:                   "ADDNODE",
		vectorable:             true,
		cardinalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, vx, vy int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+vx, y+vy) {
				return false
			}
			return g.IsNodeActive(x, y) && !g.IsNodeActive(x+vx, y+vy)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, vx, vy int) {
			g.enableAndInterlinkNodeFromCoords(x, y, vx, vy, true)
		},
	},

	// X   U       X - R
	// |       >       |
	// R   U       R - R
	{
		Name:                   "U-HORIZ",
		vectorable:             true,
		cardinalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, vx, _ int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+vx, y+1) {
				return false
			}
			return g.IsNodeActive(x, y) && g.IsNodeActive(x, y+1) && g.AreCoordsInterlinked(x, y, x, y+1) &&
				!g.IsNodeActive(x+vx, y) && !g.IsNodeActive(x+vx, y+1)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, vx, _ int) {
			g.setLinkBetweenCoords(x, y, x, y+1, false)
			g.enableAndInterlinkNodeFromCoords(x, y, vx, 0, false)
			g.enableAndInterlinkNodeFromCoords(x, y+1, vx, 0, false)
			g.setLinkBetweenCoords(x+vx, y, x+vx, y+1, true)
		},
	},

	// X - R       X   R
	//         >   |   |
	// U   U       R - R
	{
		Name:                   "U-VERT",
		vectorable:             true,
		cardinalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, _, vy int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+1, y+vy) {
				return false
			}
			return g.IsNodeActive(x, y) && g.IsNodeActive(x+1, y) && g.AreCoordsInterlinked(x, y, x+1, y) &&
				!g.IsNodeActive(x, y+vy) && !g.IsNodeActive(x+1, y+vy)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, _, vy int) {
			g.setLinkByVector(x, y, 1, 0, false)
			g.enableAndInterlinkNodeFromCoords(x, y, 0, vy, false)
			g.enableAndInterlinkNodeFromCoords(x+1, y, 0, vy, false)
			g.setLinkByVector(x, y+vy, 1, 0, true)
		},
	},

	// X   U       X - R
	// |       >   |   |
	// R   U       R - R
	{
		Name:                   "CYC-HOR",
		AddsCycle:              true,
		vectorable:             true,
		cardinalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, vx, _ int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+vx, y+1) {
				return false
			}
			return g.IsNodeActive(x, y) && g.IsNodeActive(x, y+1) && g.AreCoordsInterlinked(x, y, x, y+1) &&
				!g.IsNodeActive(x+vx, y) && !g.IsNodeActive(x+vx, y+1)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, vx, _ int) {
			if vx < 0 {
				x--
			}
			g.drawConnectedNodeRect(x, y, 2, 2)
		},
	},

	// X - R       X - R
	//         >   |   |
	// U   U       R - R
	{
		Name:                   "CYC-VER",
		AddsCycle:              true,
		vectorable:             true,
		cardinalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, _, vy int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+1, y+vy) {
				return false
			}
			return g.IsNodeActive(x, y) && g.IsNodeActive(x+1, y) && g.AreCoordsInterlinked(x, y, x+1, y) &&
				!g.IsNodeActive(x, y+vy) && !g.IsNodeActive(x+1, y+vy)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, _, vy int) {
			if vy < 0 {
				y--
			}
			g.drawConnectedNodeRect(x, y, 2, 2)
		},
	},

	// R   U       R - R         U   R     R - R
	// |       >       |     or      |  >  |
	// N - R       U   R         R - N     R   U
	//
	// !! N has no other connections !!
	{
		Name:                   "L-FLIP",
		vectorable:             true,
		cardinalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, vx, _ int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+vx, y+1) {
				return false
			}
			return !g.IsNodeActive(x+vx, y) && g.IsNodeActive(x, y) && g.IsNodeActive(x+vx, y+1) && g.IsNodeActive(x, y+1) &&
				g.AreCoordsInterlinked(x, y, x, y+1) && g.AreCoordsInterlinked(x, y+1, x+vx, y+1) &&
				!g.IsEdgeByVectorActive(x, y+1, -vx, 0) && !g.IsEdgeByVectorActive(x, y+1, 0, 1)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, vx, _ int) {
			g.enableAndInterlinkNodeFromCoords(x, y, vx, 0, false)
			g.SwapTagsAtCoords(x+vx, y, x, y+1)
			g.setLinkBetweenCoords(x+vx, y, x+vx, y+1, true)
			g.resetNodeAndConnections(x, y+1)
			g.finalizeNode(x, y+1)
			g.addNodeTag(x, y, "L-F")
		},
	},

	// X - N       X   U         N - X     U   X
	//     |   >   |        or   |      >      |
	// U   R       R - R         R   U     R - R
	//
	// !! N has no other connections !!
	{
		Name:                   "L-R-FLIP",
		vectorable:             true,
		cardinalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, vx, _ int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+vx, y+1) {
				return false
			}
			return !g.IsNodeActive(x, y+1) && g.IsNodeActive(x, y) && g.IsNodeActive(x+vx, y) && g.IsNodeActive(x+vx, y+1) &&
				g.AreCoordsInterlinked(x, y, x+vx, y) && g.AreCoordsInterlinked(x+vx, y, x+vx, y+1) &&
				!g.IsEdgeByVectorActive(x+vx, y, 0, -1) && !g.IsEdgeByVectorActive(x+vx, y, vx, 0)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, vx, _ int) {
			g.enableAndInterlinkNodeFromCoords(x, y, 0, 1, false)
			g.SwapTagsAtCoords(x, y+1, x+vx, y)
			g.setLinkBetweenCoords(x+vx, y+1, x, y+1, true)
			g.resetNodeAndConnections(x+vx, y)
			g.finalizeNode(x+vx, y)
			g.addNodeTag(x, y, "LRF")
		},
	},
	// X   U       X - R
	//         >   |   |    or mirrored in any direction
	// U   U       R - R
	//
	// !! N has no other connections !!
	{
		Name:                   "CYC-CRNER",
		AddsCycle:              true,
		vectorable:             true,
		diagonalVectorsAllowed: true,
		IsApplicableAt: func(g *Graph, x, y, vx, vy int) bool {
			if !g.AreNodesBetweenCoordsEditable(x, y, x+vx, y+vy) {
				return false
			}
			return g.IsNodeActive(x, y) && !g.IsNodeActive(x+vx, y) && !g.IsNodeActive(x, y+vy) && !g.IsNodeActive(x+vx, y+vy)
		},
		ApplyOnGraphAt: func(g *Graph, x, y, vx, vy int) {
			if vx == -1 {
				x--
			}
			if vy == -1 {
				y--
			}
			g.drawConnectedNodeRect(x, y, 2, 2)
			g.addNodeTag(x, y, "C-L")
		},
	},
}
