package editor

import (
	"github.com/nsf/termbox-go"
)

type Mode int

const (
	CommandMode Mode = iota
	InsertMode
	VisualMode
)

type Editor struct {
	cursorX, cursorY int
	buffer           []string
	mode             Mode
	selectionStartX  int
	selectionStartY  int
}

func NewEditor() *Editor {
	return &Editor{
		buffer: []string{""},
		mode:   CommandMode,
	}
}

func (e *Editor) moveCursor(dx, dy int) {
	e.cursorX += dx
	e.cursorY += dy
	// Add boundary checks if necessary
}

func (e *Editor) Draw() {
	e.render()
}

func Run() {
	editor := NewEditor()
	editor.Draw()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				return
			default:
				editor.handleKeyPress(ev)
			}
		}
		editor.Draw()
	}
}
