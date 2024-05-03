package tiler

import graph "cycdg/graph_replacement/grid_graph"

// Transforms graph replacement result to a tiled map
type Tiler struct {
	graph    *graph.Graph
	nodeSize int
	tiledMap [][]Tile
}

func (t *Tiler) Init(g *graph.Graph, nodeSize int) {
	t.graph = g
	t.nodeSize = nodeSize
}
