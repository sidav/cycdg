package main

import (
	replacement "cycdg/graph_replacement"
	"cycdg/graph_replacement/grammar"
	. "cycdg/lib/tcell_console_wrapper/console_menus"
)

func CreateGraWithParamsMenu() *replacement.GraphReplacementApplier {
	gra := &replacement.GraphReplacementApplier{}
	var width, height int

	menu := ConsoleIntValuesSelection{
		Title: "SELECT GENERATION PARAMS",
		Entries: []*ValueSelectorEntry{
			NewPointerSelectorEntry(&width, "Width", "", 4, 4, 25, 1),
			NewPointerSelectorEntry(&height, "Height", "", 4, 5, 25, 1),
			NewPointerSelectorEntry(&gra.MinRulesToApply, "Minimum applied rules", "", 1, 7, 1000, 1),
			NewPointerSelectorEntry(&gra.MinFilledPercentage, "Min fill percentage", "%", 25, 65, 100, 5),
			NewPointerSelectorEntry(&gra.MaxFilledPercentage, "Max fill percentage", "%", 25, 85, 100, 5),
			NewPointerSelectorEntry(&gra.MinCycles, "Min cycles", "", 0, 3, 100, 1),
			NewPointerSelectorEntry(&gra.MaxCycles, "Max cycles", "", 0, 8, 100, 1),
			NewPointerSelectorEntry(&gra.DesiredFeatures, "Desired features", "", 0, 5, 1000, 1),
			NewPointerSelectorEntry(&gra.MaxTeleports, "Max teleports", "", 0, 2, 1000, 1),
		},
	}
	menu.Init()
	key := ""
	for key != "ENTER" {
		menu.Show(&cw)
		cw.FlushScreen()
		key = cw.ReadKey()
		menu.UpdateForKeypress(key)

		if gra.MaxFilledPercentage < gra.MinFilledPercentage {
			gra.MaxFilledPercentage = gra.MinFilledPercentage
		}
		if gra.MaxCycles < gra.MinCycles {
			gra.MaxCycles = gra.MinCycles
		}
		if gra.MinRulesToApply < 1 {
			gra.MinRulesToApply = 1
		}
		maxRules := (2 * width * height / 3)
		if gra.MinRulesToApply > maxRules {
			gra.MinRulesToApply = maxRules
		}
	}

	grammar := grammar.CreateExampleGrammarObject()
	gra.Init(grammar, rnd, width, height)
	return gra
}
