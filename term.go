package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

// For now, there is a global screen object only visible to Go code, to avoid
// leaking the screen object outside of the abstraction provided by this file.
var screen tcell.Screen

func termStart() error {
	var err error
	if screen != nil {
		return fmt.Errorf("screen already initialized")
	}
	screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := screen.Init(); err != nil {
		return err
	}
	return nil
}

func termClear() error {
	if screen == nil {
		return fmt.Errorf("screen not initialized")
	}
	screen.Clear()
	screen.Show()
	return nil
}

func termEnd() error {
	if screen == nil {
		return fmt.Errorf("screen not initialized")
	}
	screen.Fini()
	screen = nil
	return nil
}

func termDrawText(x, y int, str string) {
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
}

func termSize() (int, int, error) {
	if screen == nil {
		return 0, 0, fmt.Errorf("screen not initialized")
	}
	x, y := screen.Size()
	return x, y, nil
}

func termGetKey() (string, error) {
	return "", fmt.Errorf("not implemented")
}
