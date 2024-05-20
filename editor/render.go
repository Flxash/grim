package editor

import (
	"github.com/nsf/termbox-go"
)

func (e *Editor) render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y, line := range e.buffer {
		for x, ch := range line {
			termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	switch e.mode {
	case VisualMode:
		e.renderVisualMode()
	}

	e.renderStatusBar()

	termbox.SetCell(e.cursorX, e.cursorY, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}
