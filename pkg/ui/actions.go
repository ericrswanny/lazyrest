package ui

import (
	"fmt"

	"github.com/ericrswanny/lazyrest/pkg/http"
	"github.com/jroimartin/gocui"
)

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

func displayError(g *gocui.Gui, err error) {
	outputView, _ := g.View("output")
	outputView.Clear()
	fmt.Fprintf(outputView, "Error: %s", err)
}
