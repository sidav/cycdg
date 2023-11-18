package main

import (
	graph "cycdg/graph_replacement/grid_graph"
	"cycdg/graph_replacement/grid_graph/graph_element"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const (
	nodeWidth      = 5
	halfNodeWidth  = nodeWidth / 2
	nodeHeight     = 3
	halfNodeHeight = nodeHeight / 2
	nodeSpacing    = 2
)

var (
	allDirections  = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	edgeDirections = [2][2]int{{1, 0}, {0, 1}}
)

func drawGraph(g *graph.Graph) {
	cw.ClearScreen()
	w, h := g.GetSize()
	for nx := 0; nx < w; nx++ {
		for ny := 0; ny < h; ny++ {
			drawNodeAt(g, nx, ny)
		}
	}
	cw.ResetStyle()
	cw.PutStringf(w*(nodeWidth+nodeSpacing), 0, "%d rules, %d cycles, filled: %d%%",
		g.AppliedRulesCount, g.CyclesCount, g.GetFilledNodesPercentage())
	cw.PutString("Applied rules: ", w*(nodeWidth+nodeSpacing), 1)
	for i := range g.AppliedRules {
		cw.PutStringf(w*(nodeWidth+nodeSpacing), i+2, "%d:%s", i, g.AppliedRules[i])
	}
}

func drawNodeAt(g *graph.Graph, nx, ny int) {
	x, y := 1+nx*(nodeWidth+nodeSpacing), 1+ny*(nodeHeight+nodeSpacing)
	background := tcell.ColorDarkBlue
	if !g.IsNodeActive(nx, ny) {
		background = tcell.ColorDarkGray
		if g.IsNodeFinalized(nx, ny) {
			background = tcell.ColorBlack
		}
	}
	cw.SetStyle(tcell.ColorBlack, background)
	cw.DrawFilledRect(' ', x, y, nodeWidth-1, nodeHeight-1)
	drawNodeEdges(g, nx, ny)
	cw.SetStyle(tcell.ColorRed, background)
	for i, tag := range g.NodeAt(nx, ny).GetTags() {
		str := tag.GetStringIdiom()
		if tag.Kind != graph_element.TagStart {
			str = fmt.Sprintf("%s%d", str, tag.Id)
		}
		cw.PutStringCenteredAt(str, x+halfNodeWidth, y+i)
	}
	// cw.PutStringCenteredAt(fmt.Sprintf("%d", g.CountEdgesAt(nx, ny)), x+halfNodeWidth, y+nodeHeight-2)
	// cw.PutStringCenteredAt(fmt.Sprintf("%d-%d",
	// 	g.CountDirEdgesAt(nx, ny, true, false), g.CountDirEdgesAt(nx, ny, false, true)),
	// 	x+halfNodeWidth, y+nodeHeight-1)
}

func drawNodeEdges(g *graph.Graph, nx, ny int) {
	x, y := 1+nx*(nodeWidth+nodeSpacing), 1+ny*(nodeHeight+nodeSpacing)
	background := tcell.ColorDarkBlue
	cx, cy := x+halfNodeWidth, y+halfNodeHeight
	for _, dir := range edgeDirections {
		if g.IsEdgeByVectorActive(nx, ny, dir[0], dir[1]) {
			cw.SetStyle(tcell.ColorDarkGreen, background)
			char := ' '
			if g.IsEdgeByVectorDirectional(nx, ny, dir[0], dir[1]) {
				if g.IsEdgeDirectedByVector(nx, ny, dir[0], dir[1]) {
					if dir[0] == 1 {
						char = '>'
					} else if dir[1] == 1 {
						char = 'V'
					} else {
						panic(fmt.Sprintf("Strange directon - %v", dir))
					}
				} else {
					if dir[0] == 1 {
						char = '<'
					} else if dir[1] == 1 {
						char = '^'
					} else {
						panic(fmt.Sprintf("Strange directon - %v", dir))
					}
				}
			}
			if len(g.GetEdgeByVector(nx, ny, dir[0], dir[1]).GetTags()) > 0 {
				tag := g.GetEdgeByVector(nx, ny, dir[0], dir[1]).GetTags()[0]
				switch tag.Kind {
				case graph_element.TagLockedEdge:
					char = rune(fmt.Sprintf("%d", tag.Id)[0])
					cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkRed)
				case graph_element.TagSecretEdge:
					char = '?'
					cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkGreen)
				}
			}
			cw.PutChar(char, cx+dir[0]*(halfNodeWidth+1), cy+dir[1]*(halfNodeHeight+1))
			cw.PutChar(char, cx+dir[0]*(halfNodeWidth+2), cy+dir[1]*(halfNodeHeight+2))
		}
	}
}
