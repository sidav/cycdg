package main

import (
	"cycdg/grid_graph"
	"cycdg/lib/random"
	"cycdg/lib/random/pcgrandom"
	"cycdg/lib/tcell_console_wrapper"
	"time"
)

var (
	cw  tcell_console_wrapper.ConsoleWrapper
	rnd random.PRNG
)

func main() {
	cw.Init()
	defer cw.Close()
	rnd = pcgrandom.NewPCG64()
	rnd.SetSeed(int(time.Now().UnixNano()))

	graph.SetRandom(rnd)
	gr := &graph.Graph{}

	testGen(gr)
	key := ""
	for key != "ESCAPE" {
		if key == "" || gr.GetFilledNodesPercentage() > 65 {
			gr.InitWithConnectedNodes(5, 5)
		}
		gr.AlterSomething()

		drawGraph(gr)
		cw.FlushScreen()
		key = cw.ReadKey()
	}
}

func testGen(gr *graph.Graph) {
	for i := 0; i < 1000; i++ {
		gr.InitWithConnectedNodes(5, 5)
		for gr.GetFilledNodesPercentage() < 100 {
			gr.AlterSomething()
		}
	}
}
