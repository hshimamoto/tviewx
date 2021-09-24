// tviewx/examples/textline
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "github.com/rivo/tview"
    "github.com/hshimamoto/tviewx"
)

func main() {
    text := tviewx.NewTextLine()
    text.AddText("aaaaa", 10)
    text.AddText("bbbbb", 15)
    text.AddText("ccccc", 3)
    text.AddText("ddddd", 10)
    text.AddText("eeeee", 10)
    text.AddText("fffff", 10)
    text.SetSeparator("[green::b]" + string(tview.Borders.Vertical))
    text.GetItem(2).SetDynamicColors(true).SetText("[yellow::b]0123456789")
    text.ReplaceTexts(3, []string{"DDDDD", "EEEEE", "FFFFF"})
    tviewx.NewApplication().SetRoot(text, true).Run()
}
