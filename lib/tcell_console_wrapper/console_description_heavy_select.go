package tcell_console_wrapper

import "github.com/gdamore/tcell/v2"

type descriptionHeavySelectMenuItem struct {
	Title            string
	DescriptionLines []string
}

type DescriptionHeavySelectMenu struct {
	Title           string
	UndertitleLines []string
	Entries         []*descriptionHeavySelectMenuItem
}

func (menu *DescriptionHeavySelectMenu) AddMenuItem(title string, descriptionLines []string) {
	menu.Entries = append(menu.Entries, &descriptionHeavySelectMenuItem{Title: title, DescriptionLines: descriptionLines})
}

func (menu *DescriptionHeavySelectMenu) Call(cw *ConsoleWrapper) (selectedIndex int) {
	cursor := 0
	x, y := 0, 0
	w, h := cw.GetConsoleSize()
	w--
	h--
	for {
		// outline and title
		cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
		cw.DrawFilledRect(' ', x, y, w, h)
		cw.InverseStyle()
		cw.DrawRect(x, y, w, h)
		cw.InverseStyle()
		cw.PutStringCenteredAt(menu.Title, x+w/2, y)
		// undertitles
		currentY := y + 1
		cw.ResetStyle()
		for _, s := range menu.UndertitleLines {
			cw.PutString(s, x+1, currentY)
			currentY++
		}
		// the entries
		for i, e := range menu.Entries {
			cw.SetStyle(tcell.ColorBlue, tcell.ColorBlack)
			if i == cursor {
				cw.InverseStyle()
			}
			cw.PutString(e.Title, x+1, currentY)
			currentY++
			cw.ResetStyle()
			for _, s := range e.DescriptionLines {
				cw.PutString(s, x+2, currentY)
				currentY++
			}
		}

		cw.FlushScreen()

		key := cw.ReadKey()
		switch key {
		case "ESCAPE":
			return -1
		case "ENTER":
			return cursor
		case "DOWN":
			cursor++
		case "UP":
			cursor--
		}
		if cursor < 0 {
			cursor = len(menu.Entries) - 1
		}
		if cursor >= len(menu.Entries) {
			cursor = 0
		}
	}
}
