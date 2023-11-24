package replacement

import (
	"cycdg/graph_replacement/grammar"
	"cycdg/lib/random"
	"fmt"
	"strings"
	"time"
)

func TestGen(prng random.PRNG, width, height, tests, fillPerc int) (testResultString string) {
	if fillPerc > 100 {
		testResultString = "Inadequate fill percentage\n"
		return
	}
	var appliedRules int
	gen := &GraphReplacementApplier{}

	progressBarCLI("Benchmarking", 0, tests+1, 20)
	start := time.Now()
	for i := 0; i < tests; i++ {
		gen.Init(prng, width, height)
		for gen.GetGraph().GetFilledNodesPercentage() < fillPerc {
			gen.ApplyRandomReplacementRuleToTheGraph()
		}
		appliedRules += gen.GetGraph().AppliedRulesCount
		progressBarCLI("Benchmarking", i+1, tests+1, 20)
	}
	totalGenTime := time.Since(start)
	testResultString = showRulesInfo()
	testResultString += fmt.Sprintf("TEST: Total %d graphs of size %dx%d, filled for %d percents\n", tests, width, height, fillPerc)
	testResultString += fmt.Sprintf("Total time %v, mean time per single gen %v\n", totalGenTime, totalGenTime/time.Duration(tests))
	testResultString += fmt.Sprintf("Total rules applied %d, mean time per rule %v\n",
		appliedRules, totalGenTime/time.Duration(appliedRules))

	return
}

func showRulesInfo() string {
	variants := 0
	for _, r := range grammar.AllInitialRules {
		variants++
		variants += len(r.Features)
	}
	str := fmt.Sprintf("Total initial rules %d (%d counting all the features)\n", len(grammar.AllInitialRules), variants)
	variants = 0
	for _, r := range grammar.AllReplacementRules {
		variants++
		variants += len(r.Features)
	}
	str += fmt.Sprintf("Total replacement rules %d (%d counting all the features)\n", len(grammar.AllReplacementRules), variants)
	return str
}

func progressBarCLI(title string, value, endvalue, bar_length int) { // because I can
	endvalue -= 1
	percent := float64(value) / float64(endvalue)
	arrow := ">"
	for i := 0; i < int(percent*float64(bar_length)); i++ {
		arrow = "-" + arrow
	}
	spaces := strings.Repeat(" ", bar_length-len(arrow)+1)
	percent_with_dec := fmt.Sprintf("%.1f", percent*100.0)
	fmt.Printf("\r%s [%s%s] %s%% (%d out of %d)", title, arrow, spaces, percent_with_dec, value, endvalue)
	if value == endvalue {
		fmt.Printf("\n")
	}
}
