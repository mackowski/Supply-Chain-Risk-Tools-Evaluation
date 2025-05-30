package main

import (
	"fmt"
	"github.com/ornatedoctrin/layout"
)

func main() {
	// Create a new default layout.
	// The layout package offers different types of layouts,
	// but a default one is often a good starting point.
	l := layout.New()

	// Add some elements to the layout.
	// You can add various types of content, such as text, views, or even other layouts.
	l.AddText("Hello from our Go Layout App!")
	l.AddText("This is a simple demonstration of the layout package.")

	// You can also add views (which represent a more structured piece of content)
	// For this example, let's just add another text element to keep it simple.
	l.AddText("---") // A separator
	l.AddText("Organizing content with Go!")

	// Render the layout to see its output.
	// The Render method typically returns a string or a byte slice
	// that represents the formatted content based on the layout's configuration.
	renderedOutput := l.Render()

	// Print the rendered output.
	fmt.Println(renderedOutput)
}