package editor

import (
	"github.com/nsf/termbox-go"
)

func (e *Editor) handleCommandModeKeyPress(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowUp:
		e.moveCursor(0, -1)
	case termbox.KeyArrowDown:
		e.moveCursor(0, 1)
	case termbox.KeyArrowLeft:
		e.moveCursor(-1, 0)
	case termbox.KeyArrowRight:
		e.moveCursor(1, 0)
	case termbox.KeyEnter:
		e.buffer = append(e.buffer[:e.cursorY+1], e.buffer[e.cursorY:]...)
		e.cursorY++
		e.cursorX = 0
	default:
		if ev.Ch == 'i' {
			e.mode = InsertMode
		} else if ev.Ch == 'v' {
			e.mode = VisualMode
			e.selectionStartX = e.cursorX
			e.selectionStartY = e.cursorY
		}
	}
}
