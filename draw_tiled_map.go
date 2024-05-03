package main

import (
	graph "cycdg/graph_replacement/grid_graph"
	. "cycdg/tiler"

	"github.com/gdamore/tcell/v2"
)

func drawTiledMap(g *graph.Graph, sx, sy int) {
	tiler := Tiler{}
	tiler.Init(g, 5)
	itm := tiler.GetTileMap()

	for x := range itm {
		for y := range itm[x] {
			color := tcell.ColorBlack
			symbol := ' '
			switch itm[x][y].TileType {
			case TileTypeUnset:
				color = tcell.ColorBlack
				symbol = ' '
			case TileTypeBarrier:
				color = tcell.ColorRed
			case TileTypeFloor:
				color = tcell.ColorBlue
				symbol = '.'
			case TileTypeDoor:
				color = tcell.ColorGreen
				symbol = '+'
			}
			cw.SetStyle(tcell.ColorBlack, color)
			cw.PutChar(symbol, x+sx, y+sy)
		}
	}
}
