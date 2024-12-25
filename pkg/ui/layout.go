package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var CurrentMethod = "GET"

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// URL Input View
	if v, err := g.SetView("url", 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "URL"
		v.Editable = true
		fmt.Fprint(v, "http://") // Placeholder text
	}

	// Method Input View
	if v, err := g.SetView("method", 0, 3, 10, 5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Method"
		fmt.Fprint(v, CurrentMethod)
	}

	// Output View
	if v, err := g.SetView("output", 0, 6, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Response"
		v.Wrap = true
	}

	// Set default focus
	g.SetCurrentView("url")
	return nil
}
