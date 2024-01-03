package main

import . "cycdg/lib/tcell_console_wrapper/console_menus"
import replacement "cycdg/graph_replacement"

func CreateGraWithParamsMenu() *replacement.GraphReplacementApplier {
	menu := ConsoleIntValuesSelection{
		Title: "SELECT GENERATION PARAMS",
		Entries: []*ValueSelectorEntry{
			NewValueSelectorEntry("Width", "", 4, 4, 25, 1),
			NewValueSelectorEntry("Height", "", 4, 4, 25, 1),
			NewValueSelectorEntry("Fill percentage", "%", 25, 100, 100, 5),
			NewValueSelectorEntry("Min cycles", "", 0, 1, 100, 1),
			NewValueSelectorEntry("Max cycles", "", 0, 4, 100, 1),
			NewValueSelectorEntry("Desired features", "", 0, 5, 1000, 1),
			NewValueSelectorEntry("Max teleports", "", 0, 2, 1000, 1),
		},
	}
	menu.Init()
	key := ""
	for key != "ENTER" {
		menu.Show(&cw)
		cw.FlushScreen()
		key = cw.ReadKey()
		menu.UpdateForKeypress(key)
	}

	gra := &replacement.GraphReplacementApplier{}

	width := menu.GetValueByIndex(0)
	height := menu.GetValueByIndex(1)
	gra.DesiredFillPercentage = menu.GetValueByIndex(2)
	gra.Init(rnd, width, height)
	gra.MinCycles = menu.GetValueByIndex(3)
	gra.MaxCycles = menu.GetValueByIndex(4)
	gra.DesiredFeatures = menu.GetValueByIndex(5)
	gra.MaxTeleports = menu.GetValueByIndex(6)

	return gra
}
