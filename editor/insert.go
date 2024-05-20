package editor

import (
	"github.com/nsf/termbox-go"
)

func (e *Editor) handleInsertModeKeyPress(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyEsc:
		e.mode = CommandMode
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		if e.cursorX > 0 {
			line := e.buffer[e.cursorY]
			e.buffer[e.cursorY] = line[:e.cursorX-1] + line[e.cursorX:]
			e.cursorX--
		}
	case termbox.KeyEnter:
		e.buffer = append(e.buffer[:e.cursorY+1], e.buffer[e.cursorY:]...)
		e.buffer[e.cursorY+1] = e.buffer[e.cursorY][e.cursorX:]
		e.buffer[e.cursorY] = e.buffer[e.cursorY][:e.cursorX]
		e.cursorY++
		e.cursorX = 0
	default:
		if ev.Ch != 0 {
			line := e.buffer[e.cursorY]
			e.buffer[e.cursorY] = line[:e.cursorX] + string(ev.Ch) + line[e.cursorX:]
			e.cursorX++
		}
	}
}
