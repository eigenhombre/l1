package main

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

// For now, there is a global screen object only visible to Go code, to avoid
// leaking the screen object outside of the abstraction provided by this file.
var screen tcell.Screen

func termStart() error {
	var err error
	if screen != nil {
		return baseError("screen already initialized")
	}
	screen, err = tcell.NewScreen()
	if err != nil {
		return extendError("termStart NewScreen", err)
	}
	if err := screen.Init(); err != nil {
		return extendError("termStart Init", err)
	}
	return nil
}

func termClear() error {
	if screen == nil {
		return baseError("screen not initialized")
	}
	screen.Clear()
	screen.Show()
	return nil
}

func termEnd() error {
	if screen == nil {
		return baseError("screen not initialized")
	}
	screen.Fini()
	screen = nil
	return nil
}

func termDrawText(x, y int, str string) error {
	if screen == nil {
		return baseError("screen not initialized")
	}
	for _, c := range str {
		var combc []rune
		w := runewidth.RuneWidth(c)
		// Handle variable-width runes:
		if w == 0 {
			combc = []rune{c}
			c = ' '
			w = 1
		}
		screen.SetContent(x, y, c, combc, tcell.StyleDefault)
		x += w
	}
	screen.Show()
	return nil
}

func termSize() (int, int, error) {
	if screen == nil {
		return 0, 0, baseError("screen not initialized")
	}
	x, y := screen.Size()
	return x, y, nil
}

func termGetKey() (string, error) {
	if screen == nil {
		return "", baseError("screen not initialized")
	}
	for {
		ev := screen.PollEvent()
		if ev == nil {
			return "", nil
		}
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				return string(ev.Rune()), nil
			case tcell.KeyCtrlC:
				return "INTR", nil
			case tcell.KeyCtrlD:
				return "EOF", nil
			case tcell.KeyCtrlL:
				return "CLEAR", nil
			case tcell.KeyBackspace:
				return "BSP", nil
			case tcell.KeyBackspace2:
				return "BSP", nil
			case tcell.KeyDelete:
				return "DEL", nil
			case tcell.KeyDown:
				return "DOWNARROW", nil
			case tcell.KeyEnd:
				return "END", nil
			case tcell.KeyLeft:
				return "LEFTARROW", nil
			case tcell.KeyRight:
				return "RIGHTARROW", nil
			case tcell.KeyUp:
				return "UPARROW", nil
			case tcell.KeyEnter:
				return "ENTER", nil
			case tcell.KeyEscape:
				return "ESC", nil
			default:
				return "", nil
			}
		}
	}
}
