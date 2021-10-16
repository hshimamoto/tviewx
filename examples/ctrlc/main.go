// tviewx/examples/ctrlc
// MIT License Copyright(c) 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "fmt"
    "github.com/hshimamoto/tviewx"
)

func main() {
    ctrlc := 0
    app := tviewx.NewApplication()
    button := tviewx.NewButton("Exit Button")
    button.SetSelectedFunc(func() {
	app.Stop()
    })
    app.SetCtrlCFunc(func() {
	ctrlc++
	button.SetLabel(fmt.Sprintf("Exit Button (%d)", ctrlc))
    })
    app.SetRoot(button, true)
    app.Run()
}
