package editor

import (
	"github.com/nsf/termbox-go"
)

func (e *Editor) renderStatusBar() {
	width, height := termbox.Size()
	status := ""

	switch e.mode {
	case CommandMode:
		status = "COMMAND MODE"
	case InsertMode:
		status = "INSERT MODE"
	case VisualMode:
		status = "VISUAL MODE"
	}

	for x := 0; x < width; x++ {
		termbox.SetCell(x, height-1, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}

	for i, ch := range status {
		termbox.SetCell(i, height-1, ch, termbox.ColorBlack, termbox.ColorWhite)
	}
}
