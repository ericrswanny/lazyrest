package ui

import (
	"fmt"

	"github.com/ericrswanny/lazyrest/pkg/http"
	"github.com/jroimartin/gocui"
)

// sendRequest sends the REST request using the http package
// Pulls the url, method, and body from the gocui fields and sends the methods.
//
// Parameters:
//   - g: The gocui.Gui
//   - v: The gocui view
//
// Returns:
//   - error: Any error encountered while sending the request
func sendRequest(g *gocui.Gui, v *gocui.View) error {
	urlView, _ := g.View("url")
	url := urlView.Buffer()
	url = url[:len(url)-1] // Remove trailing newline

	headers, body, err := http.SendRequest(url, CurrentMethod)
	if err != nil {
		displayError(g, err)
		return nil
	}

	outputView, _ := g.View("output")
	outputView.Clear()
	fmt.Fprintf(outputView, "%s\n\n%s", headers, body)

	return nil
}

// toggleMethod toggles the HTTP method between different options.
//
// Parameters:
//   - g: The gocui.Gui to access the ui
//   - v: The gocui.View
//
// Returns:
//   - error: Any error encountered while toggling the method
func toggleMethod(g *gocui.Gui, v *gocui.View) error {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	for i, method := range methods {
		if method == CurrentMethod {
			CurrentMethod = methods[(i+1)%len(methods)]
			break
		}
	}
	v.Clear()
	fmt.Fprint(v, CurrentMethod)
	return nil
}

// displayError displays errors ...
//
// Parameters:
//   - g: The gocui.Gui.
//   - err: The error to be displayed.
//
// Returns:
//   - nothing, but prints an error using fmt.Fprintf
func displayError(g *gocui.Gui, err error) {
	outputView, _ := g.View("output")
	outputView.Clear()
	fmt.Fprintf(outputView, "Error: %s", err)
}
