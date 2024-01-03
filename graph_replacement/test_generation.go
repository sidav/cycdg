package replacement

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grammar"
	"cycdg/lib/random"
	"fmt"
	"time"
)

func TestGen(prng random.PRNG, width, height, tests, fillPerc int) (testResultString string) {
	if fillPerc > 100 {
		testResultString = "Inadequate fill percentage\n"
		return
	}
	var appliedRules int
	gen := &GraphReplacementApplier{
		MaxTeleports: 2,
	}
	var totalGenTime, worstTime, bestTime time.Duration
	worstRules := make(map[string]time.Duration, 0)

	progressBarCLI("Benchmarking", 0, tests+1, 20)
	for i := 0; i < tests; i++ {
		start := time.Now()
		gen.MinFilledPercentage = fillPerc
		gen.MaxFilledPercentage = fillPerc

		gen.Init(prng, width, height)

		gen.BenchGenerate(worstRules)

		thisGenTime := time.Since(start)
		totalGenTime += thisGenTime
		if worstTime < thisGenTime {
			worstTime = thisGenTime
		}
		if bestTime == 0 || bestTime > thisGenTime {
			bestTime = thisGenTime
		}
		appliedRules += gen.AppliedRulesCount
		progressBarCLI("Benchmarking", i+1, tests+1, 20)
	}
	testResultString = showRulesInfo()
	testResultString += "=========================\n"
	testResultString += fmt.Sprintf("TEST: Total %d graphs of size %dx%d, filled for %d percents\n", tests, width, height, fillPerc)
	testResultString += fmt.Sprintf("Total time %v, mean time per single gen %v\n", totalGenTime, totalGenTime/time.Duration(tests))
	testResultString += fmt.Sprintf("Worst gen time %v, best gen time %v\n", worstTime, bestTime)
	testResultString += fmt.Sprintf("Total rules applied %d, mean %d rules per map, mean time per rule %v\n",
		appliedRules, (appliedRules+tests/2)/tests, totalGenTime/time.Duration(appliedRules))
	testResultString += fmt.Sprintf("Worst rule coords pick times:\n")
	testResultString += formatDurationMap(worstRules)

	return
}

func (ra *GraphReplacementApplier) BenchGenerate(worstRules map[string]time.Duration) map[string]time.Duration {
	for !ra.FilledEnough() {
		var rule *ReplacementRule
		var applicableCoords [][]Coords
		try := 0
		for {
			rule = ra.SelectRandomRuleToApply()
			start := time.Now()

			applicableCoords = rule.FindAllApplicableCoordVariantsRecursively(ra.graph)

			took := time.Since(start)
			if worstRules[rule.Name] < took {
				worstRules[rule.Name] = took
			}

			if len(applicableCoords) > 0 {
				break
			}
			try++
			if try > 10000 {
				ra.debugPanic("No applicable coords even after 10000 tries!")
			}
		}
		ra.applyReplacementRule(rule, applicableCoords)
	}
	return worstRules
}

func showRulesInfo() string {
	variants := 0
	for _, r := range AllInitialRules {
		variants++
		variants += len(r.MandatoryFeatures)
	}
	str := fmt.Sprintf("Total initial rules %d (%d counting all the features)\n", len(AllInitialRules), variants)

	mandatory := 0
	features := 0
	totalVariants := 0
	for _, r := range AllReplacementRules {
		mandatory += max(1, len(r.MandatoryFeatures))
		features += len(r.OptionalFeatures)
		totalVariants += max(1, len(r.MandatoryFeatures)) * (1 + len(r.OptionalFeatures))
	}
	str += fmt.Sprintf("Total replacement rules %d (%d with variants), total %d optional features\n",
		len(AllReplacementRules), mandatory, features)
	str += fmt.Sprintf("Total replacement rules variants: %d\n", totalVariants)
	return str
}
