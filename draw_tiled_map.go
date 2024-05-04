package main

import (
	graph "cycdg/graph_replacement/grid_graph"
	. "cycdg/tiler"

	"github.com/gdamore/tcell/v2"
)

func drawTiledMap(g *graph.Graph) {
	w, _ := g.GetSize()
	offsetX, offsetY := w*(nodeWidth+nodeSpacing), 1

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
			case TileTypeRoomFloor:
				color = tcell.ColorBlue
				symbol = '.'
			case TileTypeCaveFloor:
				color = tcell.ColorDarkGray
				symbol = ','
			case TileTypeDoor:
				color = tcell.ColorGreen
				symbol = '+'
			case TileTypeWall:
				color = tcell.ColorDarkRed
				symbol = '#'
			}
			cw.SetStyle(tcell.ColorBlack, color)

			cw.PutChar(symbol, x+offsetX, y+offsetY)
		}
	}
}
