// Package main is the entry point for for the lazyrest application.
// lazyrest is a terminal-based rest API client using gocui
package main

import (
	"log"

	"github.com/ericrswanny/lazyrest/pkg/ui"
	"github.com/jroimartin/gocui"
)

// main is the entry point of the lazyrest application. It initializes the
// gocui GUI, sets up the layout and keybindings, and starts the main event loop.
func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(ui.Layout)

	if err := ui.SetupKeybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
