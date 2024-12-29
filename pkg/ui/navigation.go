package ui

import (
	"github.com/jroimartin/gocui"
)

// SetupKeybindings sets the keybindings for the application.
// It returns nil unless an error is encountered.
//
// Parameters:
//   - g: The gocui.Gui to set the keybindings on
//
// Returns:
//   - error: Any error encountered while setting the keybindings.
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

// quit exits the lazyrest application.
//
// Parameters:
//   - g: The gocui.Gui to call quit on.
//
// Returns:
//   - the response of gocui.ErrQuit
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// nextView navigates to the next view in the application
//
// Parameters:
//
//	g: The gocui.Gui
//	v: the gocui.View to be navigated to
//
// Returns:
//   - error: Any error encountered while navigating to the next view
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
