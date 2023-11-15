package main

import (
	"cycdg/grid_graph"
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
	graph.SetRandom(rnd)

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
		testGen(size, tests, fill)
		return
	} else {
		testGen(5, 1000, 100)
	}

	cw.Init()
	defer cw.Close()

	gr := &graph.Graph{}

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

func testGen(size, tests, fillPerc int) {
	if fillPerc > 100 {
		testResultString = "Inadequate fill percentage\n"
		return
	}
	var appliedRules int
	gr := &graph.Graph{}

	start := time.Now()
	for i := 0; i < tests; i++ {
		gr.InitWithConnectedNodes(size, size)
		for gr.GetFilledNodesPercentage() < fillPerc {
			gr.AlterSomething()
		}
		appliedRules += gr.AppliedRulesCount
	}
	totalGenTime := time.Since(start)
	testResultString = fmt.Sprintf("TEST: Total %d graphs of size %dx%d, filled for %d percents\n", tests, size, size, fillPerc)
	testResultString += fmt.Sprintf("Total time %v, mean time per single gen %v\n", totalGenTime, totalGenTime/time.Duration(tests))
	testResultString += fmt.Sprintf("Total rules applied %d, mean time per rule %v\n",
		appliedRules, totalGenTime/time.Duration(appliedRules))
}
