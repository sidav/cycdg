package main

import (
	replacement "cycdg/graph_replacement"
	"cycdg/lib/random"
	"cycdg/lib/random/pcgrandom"
	"cycdg/lib/tcell_console_wrapper"
	"flag"
	"fmt"
	"time"
)

var (
	cw                  tcell_console_wrapper.ConsoleWrapper
	width, height, fill int
	rnd                 random.PRNG
	testResultString    string
)

func main() {
	defer func() { fmt.Printf(testResultString) }()

	rnd = pcgrandom.NewPCG64()
	rnd.SetSeed(int(time.Now().UnixNano()))

	if execArgs() {
		return
	}

	cw.Init()
	defer cw.Close()

	gen := replacement.GraphReplacementApplier{}

	key := ""
	for key != "ESCAPE" {
		if key == "" || gen.GetGraph().GetFilledNodesPercentage() >= fill {
			gen.Init(rnd, width, height)
		} else {
			gen.ApplyRandomReplacementRuleToTheGraph()
			for key == "ENTER" && gen.GetGraph().GetFilledNodesPercentage() < fill {
				gen.ApplyRandomReplacementRuleToTheGraph()
			}
		}

		drawGraph(&gen)
		cw.FlushScreen()
		key = cw.ReadKey()
	}
}

// returns true if the program should exit
func execArgs() bool {
	benchOnly := false
	var testMapsCount int
	flag.IntVar(&width, "w", 5, "Graph width")
	flag.IntVar(&height, "h", 5, "Graph height")
	flag.IntVar(&fill, "fill", 70, "Fill percentage")
	flag.BoolVar(&benchOnly, "b", false, "Run benchmark only")
	flag.IntVar(&testMapsCount, "total", 1000, "Generated maps count")
	flag.Parse()

	if testMapsCount > 0 {
		testResultString = replacement.TestGen(rnd, width, height, testMapsCount, fill)
	}
	return benchOnly
}
