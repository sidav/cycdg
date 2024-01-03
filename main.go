package main

import (
	replacement "cycdg/graph_replacement"
	"cycdg/lib/random"
	"cycdg/lib/random/pcgrandom"
	"cycdg/lib/tcell_console_wrapper"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
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

	if execArgs() {
		return
	}

	cw.Init()
	defer cw.Close()

	gen := CreateGraWithParamsMenu()

	key := ""
	for key != "ESCAPE" {
		if key == "e" {
			logName := exportRulesLog(gen)
			cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
			cw.PutString("Exported the log to file "+logName, 0, 0)
			cw.FlushScreen()
			key = cw.ReadKey()
			continue
		}
		if key == "" || gen.FilledEnough() {
			gen.Reset()
		} else {
			gen.ApplyRandomReplacementRuleToTheGraph()
			for key == "ENTER" && !gen.FilledEnough() {
				gen.ApplyRandomReplacementRuleToTheGraph()
			}
		}

		drawGraph(gen)
		cw.FlushScreen()
		key = cw.ReadKey()
	}
}

// returns true if the program should exit
func execArgs() bool {
	var width, height, fill int
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

func exportRulesLog(gen *replacement.GraphReplacementApplier) string {
	rlog := ""
	for i := range gen.AppliedRules {
		if i == 0 {
			rlog += "Initial rule:\n"
		} else {
			rlog += fmt.Sprintf("Rule %d:\n", i+1)
		}
		rlog += fmt.Sprintf("  %s\n", gen.AppliedRules[i].StringifyRule())
		rlog += fmt.Sprintf("  %s\n", gen.AppliedRules[i].StringifyCoords())
	}
	fileName := time.Now().Format("2006-01-02_15-04-05") + ".log"
	os.WriteFile(fileName, []byte(rlog), 0644)
	return fileName
}
