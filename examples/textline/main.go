// tviewx/examples/textline
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "fmt"
    "time"

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
    text.SetDynamic(func(idx int, orig string) string {
	if idx == 0 {
	    return fmt.Sprintf("%d", time.Now().Unix())
	}
	return orig
    })
    app := tviewx.NewApplication()
    go func() {
	for {
	    time.Sleep(time.Second)
	    app.QueueUpdateDraw(func(){})
	}
    }()
    app.SetRoot(text, true).Run()
}
