package editor

import (
	"github.com/nsf/termbox-go"
)

type Editor struct {
	cursorX, cursorY int
	buffer           []string
}

func NewEditor() *Editor {
	return &Editor{
		buffer: []string{""},
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
