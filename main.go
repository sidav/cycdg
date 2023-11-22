package main

import (
	replacement "cycdg/graph_replacement"
	"cycdg/lib/random"
	"cycdg/lib/random/pcgrandom"
	"cycdg/lib/tcell_console_wrapper"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	cw               tcell_console_wrapper.ConsoleWrapper
	rnd              random.PRNG
	testResultString string
)

func main() {
	defer func() { fmt.Printf(testResultString) }()

	rnd = pcgrandom.NewPCG64()
	rnd.SetSeed(int(time.Now().UnixNano()))

	if len(os.Args) > 1 {
		if len(os.Args) < 4 {
			fmt.Println("TEST USAGE: go run *.go [graph size] [total generations] [desired fill percentage]")
			return
		}
		size, err1 := strconv.Atoi(os.Args[1])
		tests, err2 := strconv.Atoi(os.Args[2])
		fill, err3 := strconv.Atoi(os.Args[3])
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Please use numbers as args.")
			return
		}
		testResultString = replacement.TestGen(rnd, size, tests, fill)
		return
	} else {
		testResultString = replacement.TestGen(rnd, 5, 1000, 100)
	}

	cw.Init()
	defer cw.Close()

	gen := replacement.GraphReplacementApplier{}

	key := ""
	for key != "ESCAPE" {
		if key == "" || gen.GetGraph().GetFilledNodesPercentage() >= 65 {
			gen.Init(rnd, 5, 5)
		} else {
			gen.ApplyRandomReplacementRuleToTheGraph()
			for key == "ENTER" && gen.GetGraph().GetFilledNodesPercentage() < 65 {
				gen.ApplyRandomReplacementRuleToTheGraph()
			}
		}

		drawGraph(gen.GetGraph())
		cw.FlushScreen()
		key = cw.ReadKey()
	}
}
