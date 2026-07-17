package grammar

import (
	. "cycdg/graph_replacement/geometry"
	graph "cycdg/graph_replacement/grid_graph"
	. "cycdg/graph_replacement/grid_graph/graph_element"
	"cycdg/lib/random"
	"fmt"
)

var rnd random.PRNG

func SetRandom(r random.PRNG) {
	rnd = r
}

var (
	cardinalDirections = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
)

func debugPanic(msg string, args ...interface{}) {
	fmt.Println()
	panic(sprintf(msg, args...))
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

func getRandomGraphCoordsByFunc(g *graph.Graph, good func(c Coords) bool) Coords {
	var candidates []Coords
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if good(NewCoords(x, y)) {
				candidates = append(candidates, NewCoords(x, y))
			}
		}
	}
	if len(candidates) == 0 {
		debugPanic("No candidates!")
		// return geometry.NewCoords(-1, -1)
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind]
}

func getRandomGraphCoordsByTag(g *graph.Graph, tag TagKind) Coords {
	var candidates []Coords
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := NewCoords(x, y)
			if g.DoesNodeHaveTag(c, tag) {
				candidates = append(candidates, c)
			}
		}
	}
	if len(candidates) == 0 {
		debugPanic("No candidates!")
		// return geometry.NewCoords(-1, -1)
	}
	ind := rnd.Rand(len(candidates))
	return candidates[ind]
}

func getRandomGraphCoordsByScore(g *graph.Graph, score func(x, y int) int) Coords {
	var candidates []Coords
	var scores []int
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			score := score(x, y)
			if score > 0 {
				candidates = append(candidates, NewCoords(x, y))
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

func getFirstGraphCoordsWithTag(g *graph.Graph, tag TagKind) Coords {
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if g.NodeAtXY(x, y).HasTag(tag) {
				return NewCoords(x, y)
			}
		}
	}
	panic("No coords with requested tag exist!")
}

func doesGraphContainNodeTag(g *graph.Graph, tag TagKind) bool {
	w, h := g.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if g.NodeAtXY(x, y).HasTag(tag) {
				return true
			}
		}
	}
	return false
}

func isTagMovable(tag *Tag) bool {
	t := tag.Kind
	// TODO: allow moving teleports somehow.
	// currently they're disabled because it causes unreachable keys sometimes (teleport is moved behind a door, which is then locked).
	restrictedTags := [...]TagKind{
		TagKey, TagHalfkey, TagMasterkey,
		TagStart,
		TagTeleportBidir, TagTeleportFrom, TagTeleportTo,
	}
	for i := range restrictedTags {
		if t == restrictedTags[i] {
			return false
		}
	}
	return true
}

func areAllNodeTagsMovable(g *graph.Graph, crds Coords) bool {
	tags := g.NodeAtXY(crds.Unwrap()).GetTags()
	for _, t := range tags {
		if !isTagMovable(t) {
			return false
		}
	}
	return true
}

func AddRandomHazardAt(g *graph.Graph, crds Coords) {
	possibleTags := []TagKind{TagBoss, TagTrap, TagHazard}
	g.AddNodeTagByCoords(crds, possibleTags[rnd.Rand(len(possibleTags))])
}

func moveRandomNodeTag(g *graph.Graph, from, to Coords) {
	fromNode := g.NodeAtXY(from.Unwrap())
	fromTags := fromNode.GetTags()
	if len(fromTags) == 0 {
		return
	}
	index := rnd.Rand(len(fromTags))
	if !isTagMovable(fromTags[index]) {
		return
	}
	toNode := g.NodeAtXY(to.Unwrap())
	toNode.AddTag(fromTags[index].Kind, fromTags[index].Id)
	fromNode.RemoveTagByIndex(index)
}

func PushNodeContentsInRandomDirection(g *graph.Graph, crds Coords) {
	pushTo := getRandomGraphCoordsByFunc(g, func(c Coords) bool {
		return !g.IsNodeActiveXY(c.Unwrap()) && crds.IsAdjacentTo(c)
	})
	if pushTo.EqualsPair(-1, -1) || !areAllNodeTagsMovable(g, crds) {
		return
	}
	g.EnableNodeByCoords(pushTo)
	g.EnableDirLinkByCoords(crds, pushTo)
	g.SwapNodeTags(crds, pushTo)
}

func PushNodeContentsInRandomDirectionWithEdgeTag(g *graph.Graph, crds Coords, tag TagKind) {
	pushTo := getRandomGraphCoordsByFunc(g, func(c Coords) bool {
		return !g.IsNodeActiveXY(c.Unwrap()) && crds.IsAdjacentTo(c)
	})
	if pushTo.EqualsPair(-1, -1) || !areAllNodeTagsMovable(g, crds) {
		return
	}
	g.EnableNodeByCoords(pushTo)
	g.EnableDirLinkByCoords(crds, pushTo)
	g.AddEdgeTagByCoords(crds, pushTo, tag)
	g.SwapNodeTags(crds, pushTo)
}
