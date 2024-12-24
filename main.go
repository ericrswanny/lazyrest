package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jroimartin/gocui"
)

var currentMethod = "GET"

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
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
		fmt.Fprint(v, currentMethod)
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

func keybindings(g *gocui.Gui) error {
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

func toggleMethod(g *gocui.Gui, v *gocui.View) error {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	for i, method := range methods {
		if method == currentMethod {
			currentMethod = methods[(i+1)%len(methods)]
			break
		}
	}
	v.Clear()
	fmt.Fprint(v, currentMethod)
	return nil
}

func sendRequest(g *gocui.Gui, v *gocui.View) error {
	urlView, _ := g.View("url")
	url := urlView.Buffer()
	url = url[:len(url)-1] // Remove trailing newline

	var req *http.Request
	var err error

	if currentMethod == "POST" || currentMethod == "PUT" {
		// Add request body (hardcoded for now)
		reqBody := bytes.NewBuffer([]byte(`{"key":"value"}`))
		req, err = http.NewRequest(currentMethod, url, reqBody)
	} else {
		req, err = http.NewRequest(currentMethod, url, nil)
	}

	if err != nil {
		displayError(g, err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		displayError(g, err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		displayError(g, err)
		return nil
	}

	outputView, _ := g.View("output")
	outputView.Clear()
	fmt.Fprintf(outputView, "Status: %s\nHeaders:\n", resp.Status)
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Fprintf(outputView, "%s: %s\n", key, value)
		}
	}
	fmt.Fprintf(outputView, "\nBody:\n%s", string(body))

	return nil
}

func displayError(g *gocui.Gui, err error) {
	outputView, _ := g.View("output")
	outputView.Clear()
	fmt.Fprintf(outputView, "Error: %s", err)
}
