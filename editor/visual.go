package editor

import (
	"github.com/nsf/termbox-go"
)

func (e *Editor) handleVisualModeKeyPress(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyEsc:
		e.mode = CommandMode
	case termbox.KeyArrowUp:
		e.moveCursor(0, -1)
	case termbox.KeyArrowDown:
		e.moveCursor(0, 1)
	case termbox.KeyArrowLeft:
		e.moveCursor(-1, 0)
	case termbox.KeyArrowRight:
		e.moveCursor(1, 0)
	}
}

func (e *Editor) renderVisualMode() {
	startX, startY := e.selectionStartX, e.selectionStartY
	endX, endY := e.cursorX, e.cursorY
	if startY > endY || (startY == endY && startX > endX) {
		startX, startY, endX, endY = endX, endY, startX, startY
	}

	for y := startY; y <= endY; y++ {
		xStart, xEnd := 0, len(e.buffer[y])
		if y == startY {
			xStart = startX
		}
		if y == endY {
			xEnd = endX
		}
		for x := xStart; x < xEnd; x++ {
			termbox.SetCell(x, y, rune(e.buffer[y][x]), termbox.ColorDefault, termbox.ColorBlue)
		}
	}
}
