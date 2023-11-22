package replacement

import (
	"cycdg/lib/random"
	"fmt"
	"strings"
	"time"
)

func TestGen(prng random.PRNG, size, tests, fillPerc int) (testResultString string) {
	if fillPerc > 100 {
		testResultString = "Inadequate fill percentage\n"
		return
	}
	var appliedRules int
	gen := &GraphReplacementApplier{}

	start := time.Now()
	for i := 0; i < tests; i++ {
		gen.Init(prng, size, size)
		for gen.GetGraph().GetFilledNodesPercentage() < fillPerc {
			gen.ApplyRandomReplacementRuleToTheGraph()
		}
		appliedRules += gen.GetGraph().AppliedRulesCount
		progressBarCLI("Benchmarking", i+1, tests+1, 20)
	}
	totalGenTime := time.Since(start)
	testResultString = fmt.Sprintf("TEST: Total %d graphs of size %dx%d, filled for %d percents\n", tests, size, size, fillPerc)
	testResultString += fmt.Sprintf("Total time %v, mean time per single gen %v\n", totalGenTime, totalGenTime/time.Duration(tests))
	testResultString += fmt.Sprintf("Total rules applied %d, mean time per rule %v\n",
		appliedRules, totalGenTime/time.Duration(appliedRules))

	return
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
