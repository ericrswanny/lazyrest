// Package http provides utilities for displaying the gocui ui for lazyrest.
package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var CurrentMethod = "GET"

// Layout provides the layout configuration for the application. It uses the provided goui.Gui for building the ui, and returns nil unless an error is encountered.
// It returns nil unless an error is encountered.
//
// Parameters:
//   - g: The gocui.Gui to build the ui on
//
// Returns:
//   - error: Any error encountered while building the ui.
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
