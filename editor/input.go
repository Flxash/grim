package editor

import (
	"github.com/nsf/termbox-go"
)

func (e *Editor) handleKeyPress(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowUp:
		e.moveCursor(0, -1)
	case termbox.KeyArrowDown:
		e.moveCursor(0, 1)
	case termbox.KeyArrowLeft:
		e.moveCursor(-1, 0)
	case termbox.KeyArrowRight:
		e.moveCursor(1, 0)
	default:
		if ev.Ch != 0 {
			if len(e.buffer) == 0 {
				e.buffer = append(e.buffer, "")
			}
			e.buffer[e.cursorY] += string(ev.Ch)
		}
	}
}
