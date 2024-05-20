package editor

import (
	"github.com/nsf/termbox-go"
)

func (e *Editor) handleKeyPress(ev termbox.Event) {
	switch e.mode {
	case CommandMode:
		e.handleCommandModeKeyPress(ev)
	case InsertMode:
		e.handleInsertModeKeyPress(ev)
	case VisualMode:
		e.handleVisualModeKeyPress(ev)
	}
}
