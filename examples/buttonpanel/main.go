// tviewx/examples/buttonpanel
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "github.com/rivo/tview"
    "github.com/hshimamoto/tviewx"
)

func main() {
    label := tview.NewTextView().SetText("undefined")
    bp := tviewx.NewButtonPanel()
    bp.AddButton("Button1")
    bp.AddButton("Button2")
    bp.AddButton("Button3")
    bp.AddButton("Button4")
    bp.SetSelectedFunc(func(b *tviewx.Button) {
	label.SetText(b.GetLabel())
    })
    flex := tview.NewFlex().SetDirection(tview.FlexRow)
    flex.AddItem(label, 1, 1, false)
    flex.AddItem(bp, 0, 1, true)
    tviewx.NewApplication().SetRoot(flex, true).Run()
}
