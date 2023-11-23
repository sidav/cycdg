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
	width, height    int
	rnd              random.PRNG
	testResultString string
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
		if key == "" || gen.GetGraph().GetFilledNodesPercentage() >= 65 {
			gen.Init(rnd, width, height)
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

// returns true if the program should exit
func execArgs() bool {
	exit := false
	tests, fill := 1000, 100

	argsCount := len(os.Args)
	if argsCount != 1 && argsCount != 3 && argsCount != 5 {
		fmt.Println("USAGE: ")
		fmt.Println("  Generation: go run *.go [graph width] [graph height]")
		fmt.Println("  Benchmark:  go run *.go [graph width] [graph height] [total generated graphs] [desired fill percentage]")
		return true
	}

	if len(os.Args) <= 1 {
		width, height = 5, 5
	} else if len(os.Args) > 2 {
		for i := range os.Args {
			if i == 0 {
				continue
			}
			if _, err := strconv.Atoi(os.Args[i]); err != nil {
				fmt.Println("Please use numbers as args.")
				return true
			}
		}
		width, _ = strconv.Atoi(os.Args[1])
		height, _ = strconv.Atoi(os.Args[2])
	}
	if len(os.Args) > 3 {
		tests, _ = strconv.Atoi(os.Args[3])
		fill, _ = strconv.Atoi(os.Args[4])
		exit = true
	}

	testResultString = replacement.TestGen(rnd, width, height, tests, fill)
	return exit
}
