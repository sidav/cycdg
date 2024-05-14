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
			fgcolor := tcell.ColorBlack
			bgcolor := tcell.ColorDarkMagenta
			symbol := '?'
			switch itm[x][y].TileType {
			case TileTypeUnset:
				bgcolor = tcell.ColorBlack
				symbol = ' '
			case TileTypeBarrier:
				bgcolor = tcell.ColorRed
			case TileTypeRoomFloor:
				bgcolor = tcell.ColorBlue
				symbol = '.'
			case TileTypeCaveFloor:
				bgcolor = tcell.ColorDarkGray
				symbol = ','
			case TileTypeDoor:
				bgcolor = tcell.ColorGreen
				symbol = '+'
			case TileTypeWall:
				bgcolor = tcell.ColorDarkRed
				symbol = '#'
			case TileTypeSecretDoor:
				fgcolor = tcell.ColorDarkGray
				bgcolor = tcell.ColorDarkRed
			case TileTypeLockedDoor:
				fgcolor = tcell.ColorRed
				bgcolor = tcell.ColorDarkRed
				symbol = 'X'
			}
			cw.SetStyle(fgcolor, bgcolor)

			cw.PutChar(symbol, x+offsetX, y+offsetY)
		}
	}
}
