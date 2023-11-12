package tcell_console_wrapper

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"strings"
	"time"
)

type ConsoleWrapper struct {
	screen tcell.Screen
	style  tcell.Style
}

func (c *ConsoleWrapper) Init() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	var e error
	c.screen, e = tcell.NewScreen()
	if e != nil {
		panic(e)
	}
	if e = c.screen.Init(); e != nil {
		panic(e)
	}
	// c.screen.EnableMouse()
	c.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
	c.screen.SetStyle(c.style)
	c.screen.Clear()
}

func (c *ConsoleWrapper) Close() {
	c.screen.Fini()
}

func (c *ConsoleWrapper) ClearScreen() {
	c.screen.Clear()
}

func (c *ConsoleWrapper) FlushScreen() {
	c.screen.Show()
}

func (c *ConsoleWrapper) GetConsoleSize() (int, int) {
	return c.screen.Size()
}

func (c *ConsoleWrapper) PutChar(chr rune, x, y int) {
	c.screen.SetCell(x, y, c.style, chr)
}

func (c *ConsoleWrapper) PutString(str string, x, y int) {
	for i := 0; i < len(str); i++ {
		c.screen.SetCell(x+i, y, c.style, rune(str[i]))
	}
}

func (c *ConsoleWrapper) PutStringf(x, y int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)
	for i := 0; i < len(str); i++ {
		c.screen.SetCell(x+i, y, c.style, rune(str[i]))
	}
}

func (c *ConsoleWrapper) PutStringCenteredAt(str string, x, y int) {
	length := len(str)
	for i := range str {
		c.screen.SetCell(x+i-length/2, y, c.style, rune(str[i]))
	}
}

func (c *ConsoleWrapper) SetStyle(fg, bg tcell.Color) {
	c.style = c.style.Background(bg).Foreground(fg)
}

func (c *ConsoleWrapper) InverseStyle() {
	fg, bg, _ := c.style.Decompose()
	c.SetStyle(bg, fg)
}

func (c *ConsoleWrapper) ResetStyle() {
	c.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
}

func (c *ConsoleWrapper) DrawFilledRect(char rune, fx, fy, w, h int) {
	for x := fx; x <= fx+w; x++ {
		for y := fy; y <= fy+h; y++ {
			c.PutChar(char, x, y)
		}
	}
}

func (c *ConsoleWrapper) DrawRect(fx, fy, w, h int) {
	for x := fx; x <= fx+w; x++ {
		c.PutChar(' ', x, fy)
		c.PutChar(' ', x, fy+h)
	}
	for y := fy; y <= fy+h; y++ {
		c.PutChar(' ', fx, y)
		c.PutChar(' ', fx+w, y)
	}
}

func (c *ConsoleWrapper) ReadKey() string {
	for {
		ev := c.screen.PollEvent()
		// consider only recent key presses
		if time.Since(ev.When()) < time.Duration(100)*time.Millisecond {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyCtrlC {
					return "EXIT"
				}
				return eventToKeyString(ev)
			}
		}
		// time.Sleep(50 * time.Millisecond)
	}
}

func (c *ConsoleWrapper) ReadKeyAsync(maxMsSinceKeyPress int) string { // returns an empty string if no key was pressed
	for c.screen.HasPendingEvent() {
		ev := c.screen.PollEvent()
		// consider only recent key presses
		if time.Since(ev.When()) < time.Duration(maxMsSinceKeyPress)*time.Millisecond {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyCtrlC {
					return "EXIT"
				}
				return eventToKeyString(ev)
			}
		}
	}
	return ""
}

func eventToKeyString(ev *tcell.EventKey) string {
	switch ev.Key() {
	case tcell.KeyUp:
		return "UP"
	case tcell.KeyRight:
		return "RIGHT"
	case tcell.KeyDown:
		return "DOWN"
	case tcell.KeyLeft:
		return "LEFT"
	case tcell.KeyEscape:
		return "ESCAPE"
	case tcell.KeyEnter:
		return "ENTER"
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		return "BACKSPACE"
	case tcell.KeyTab:
		return "TAB"
	case tcell.KeyDelete:
		return "DELETE"
	case tcell.KeyInsert:
		return "INSERT"
	case tcell.KeyEnd:
		return "END"
	case tcell.KeyHome:
		return "HOME"
	default:
		return string(ev.Rune())
	}
}

func (c *ConsoleWrapper) PutTextInRect(text string, x, y, w int) {
	if w == 0 {
		w, _ = c.GetConsoleSize()
	}
	cx, cy := x, y
	splittedText := strings.Split(text, " ")
	for _, word := range splittedText {
		if cx-x+len(word) > w || word == "\\n" || word == "\n" {
			cx = x
			cy += 1
		}
		if word != "\\n" && word != "\n" {
			c.PutString(word, cx, cy)
			cx += len(word) + 1
		}
	}
}

func (c *ConsoleWrapper) RenderProgressBar(name string, x, y, width, value, maxValue int, textColor, fullColor, emptyColor tcell.Color) {
	c.SetStyle(textColor, tcell.ColorBlack)
	text := ""
	if len(name) > 0 {
		text := fmt.Sprintf("%s %d/%d", name, value, maxValue)
		c.PutString(text, x, y)
	}
	titleLen := len(text)
	barWidth := width - titleLen - 1
	if barWidth > 2 {
		barX := x + len(text) + 1
		filledChars := barWidth * value / maxValue
		for i := 0; i < filledChars; i++ {
			c.SetStyle(fullColor, tcell.ColorBlack)
			c.PutChar('=', barX+i, y)
		}
		for i := filledChars; i < barWidth; i++ {
			c.SetStyle(emptyColor, tcell.ColorBlack)
			c.PutChar('-', barX+i, y)
		}
	}
	c.ResetStyle()
}

func (c *ConsoleWrapper) DrawPartiallyFilledEllipse(x, y, w, h, value, maxValue int, emptyRune, filledRune rune) {
	filledH := value * h / maxValue
	a := w / 2
	b := h / 2
	// scale := a / b
	rectX, rectY := x+a, y+b
	for i := 0; i < w; i++ {
		dx := 2 * (x + i - rectX)
		if w%2 == 0 {
			dx++
		}
		for j := 0; j < h; j++ {
			dy := 2 * (a * (y + j - rectY) / b)
			if h%2 == 0 {
				dy++
			}
			// check if the coords are in the ellipse
			if getApproxDistFromTo(0, 0, dx, dy) <= w {
				if j >= h-filledH {
					c.PutChar(filledRune, i+x, j+y)
				} else {
					c.PutChar(emptyRune, i+x, j+y)
				}
			}
		}
	}
}

func getApproxDistFromTo(x1, y1, x2, y2 int) int {
	diffX := x1 - x2
	if diffX < 0 {
		diffX = -diffX
	}
	diffY := y1 - y2
	if diffY < 0 {
		diffY = -diffY
	}
	if diffX > diffY {
		return diffX + (diffY >> 1)
	} else {
		return diffY + (diffX >> 1)
	}
}
