package consolemenus

import (
	. "cycdg/lib/tcell_console_wrapper"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type ConsoleIntValuesSelection struct {
	Title   string
	Entries []*ValueSelectorEntry

	cursorPos int
}

func (menu *ConsoleIntValuesSelection) Init() {

}

func (menu *ConsoleIntValuesSelection) Show(cw *ConsoleWrapper) {
	w, _ := cw.GetConsoleSize()
	cw.ClearScreen()
	cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
	cw.PutStringCenteredAt(menu.Title, w/2, 0)
	for i, e := range menu.Entries {
		cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
		if menu.cursorPos == i {
			cw.SetStyle(tcell.ColorBlack, tcell.ColorGray)
		}
		cw.PutString(fmt.Sprintf("%-25s%-8s", e.Title, e.getValueStringWithArrows()), 0, i+1)
	}
}

func (menu *ConsoleIntValuesSelection) UpdateForKeypress(key string) {
	switch key {
	case "UP":
		menu.cursorPos--
		if menu.cursorPos < 0 {
			menu.cursorPos = len(menu.Entries) - 1
		}
	case "DOWN":
		menu.cursorPos++
		if menu.cursorPos >= len(menu.Entries) {
			menu.cursorPos = 0
		}
	case "LEFT":
		menu.Entries[menu.cursorPos].decrease()
	case "RIGHT":
		menu.Entries[menu.cursorPos].increase()
	}
}

func (menu *ConsoleIntValuesSelection) GetValueByIndex(index int) int {
	return menu.Entries[index].currentValue
}
