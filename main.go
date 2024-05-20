package main

import (
	"log"

	"github.com/Flxash/grim/editor"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	editor.Run()
}
