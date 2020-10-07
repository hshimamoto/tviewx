// tviewx/examples/helloworld
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "github.com/rivo/tview"
    "github.com/hshimamoto/tviewx"
)

func main() {
    box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
    tviewx.NewApplication().SetRoot(box, true).Run()
}
