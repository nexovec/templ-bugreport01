package main

import (
	"context"
	"os"

	"github.com/a-h/templ"
)

func main() {
	var t templ.Component = AnotherThing()
	t.Render(context.Background(), os.Stdout)
}
