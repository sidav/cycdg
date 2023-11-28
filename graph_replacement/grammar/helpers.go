package grammar

import (
	"cycdg/graph_replacement/geometry"
	graph "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
	"fmt"
)

var (
	cardinalDirections = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
)

func debugPanic(msg string, args ...interface{}) {
	panic(sprintf(msg, args...))
}

// note: it's not IN rectangle!
func areCoordsOnRectangle(x, y, rx, ry, w, h int) bool {
	if x < rx || x >= rx+w || y < ry || y >= ry+h {
		return false
	}
	return x == rx || x == rx+w-1 || y == ry || y == ry+h-1
}

func areCoordsAdjacent(x1, y1, x2, y2 int) bool {
	dx := intabs(x2 - x1)
	dy := intabs(y2 - y1)
	return dx+dy == 1
}

func sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func intabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minint(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getRandomGraphCoordsByFunc(g *graph.Graph, good func(x, y int) bool) geometry.Coords {
	var candidates []geometry.Coords
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if good(x, y) {
				candidates = append(candidates, [2]int{x, y})
			}
		}
	}
	if len(candidates) == 0 {
		panic("No candidates!")
		// return geometry.NewCoords(-1, -1)
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind]
}

func getRandomGraphCoordsByScore(g *graph.Graph, score func(x, y int) int) geometry.Coords {
	var candidates []geometry.Coords
	var scores []int
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			score := score(x, y)
			if score > 0 {
				candidates = append(candidates, [2]int{x, y})
				scores = append(scores, score)
			}
		}
	}
	if len(candidates) == 0 {
		panic("No scored candidates!")
	}
	ind := rnd.SelectRandomIndexFromWeighted(len(candidates), func(i int) int { return scores[i] })
	return candidates[ind]
}

func doesGraphContainNodeTag(g *graph.Graph, tag TagKind) bool {
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if g.NodeAt(x, y).HasTag(tag) {
				return true
			}
		}
	}
	return false
}

func AddRandomHazardAt(g *graph.Graph, crds geometry.Coords) {
	possibleTags := []TagKind{TagBoss, TagTrap, TagHazard}
	g.AddNodeTagByCoords(crds, possibleTags[rnd.Rand(len(possibleTags))])
}

func moveRandomNodeTag(g *graph.Graph, from, to geometry.Coords) {
	fromNode := g.NodeAt(from.Unwrap())
	fromTags := fromNode.GetTags()
	if len(fromTags) == 0 {
		return
	}
	index := rnd.Rand(len(fromTags))
	toNode := g.NodeAt(to.Unwrap())
	toNode.AddTag(fromTags[index].Kind, fromTags[index].Id)
	fromNode.RemoveTagByIndex(index)
}

func PushNodeContentsInRandomDirection(g *graph.Graph, crds geometry.Coords) {
	pushTo := getRandomGraphCoordsByFunc(g, func(x, y int) bool {
		return !g.IsNodeActive(x, y) && crds.IsAdjacentToXY(x, y)
	})
	if pushTo.EqualsPair(-1, -1) {
		return
	}
	g.EnableNodeByCoords(pushTo)
	g.EnableDirectionalLinkBetweenCoords(crds, pushTo)
	g.SwapNodeTags(crds, pushTo)
}

func PushNodeContentsInRandomDirectionWithEdgeTag(g *graph.Graph, crds geometry.Coords, tag TagKind) {
	pushTo := getRandomGraphCoordsByFunc(g, func(x, y int) bool {
		return !g.IsNodeActive(x, y) && crds.IsAdjacentToXY(x, y)
	})
	if pushTo.EqualsPair(-1, -1) {
		return
	}
	g.EnableNodeByCoords(pushTo)
	g.EnableDirectionalLinkBetweenCoords(crds, pushTo)
	g.AddEdgeTagByCoords(crds, pushTo, tag)
	g.SwapNodeTags(crds, pushTo)
}
