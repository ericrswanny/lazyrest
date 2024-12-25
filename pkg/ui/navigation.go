package ui

import (
	"github.com/jroimartin/gocui"
)

func SetupKeybindings(g *gocui.Gui) error {
	// Quit application
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// Navigate views
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Send request
	if err := g.SetKeybinding("url", gocui.KeyEnter, gocui.ModNone, sendRequest); err != nil {
		return err
	}

	// Toggle HTTP methods
	if err := g.SetKeybinding("method", gocui.KeyEnter, gocui.ModNone, toggleMethod); err != nil {
		return err
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	views := g.Views()
	for i, view := range views {
		if view == v {
			next := (i + 1) % len(views)
			g.SetCurrentView(views[next].Name())
			break
		}
	}
	return nil
}
