// tviewx/examples/textlinelist
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "math/rand"

    "github.com/rivo/tview"
    "github.com/hshimamoto/tviewx"
)

func main() {
    list := tviewx.NewTextLineList()
    create := func() *tviewx.TextLineListItem {
	it := tviewx.NewTextLineListItem()
	rtext := func(n int) string {
	    text := make([]byte, n)
	    for n > 0 {
		n--
		ascii := rand.Intn(27) + 65 // 'A'
		text[n] = byte(ascii)
	    }
	    return string(text)
	}
	it.AddText(rtext(20), 20)
	it.AddText(rtext(10), 10)
	it.AddText(rtext(10), 10)
	it.AddText(rtext(30), 20)
	it.AddText(rtext(20), 20)
	it.AddButton(rtext(15))
	it.AddButton(rtext(15))
	it.AddButton(rtext(15))
	it.SetSelectedFunc(func(b *tviewx.Button) {
	    it.GetItem(0).SetText(b.GetLabel())
	    list.CloseMenu()
	})
	return it
    }
    for i := 0; i < 20; i++ {
	it := create()
	list.AddItem(it)
    }
    //
    flex := tview.NewFlex().SetDirection(tview.FlexRow)
    flex.AddItem(list, 10, 1, true)
    flex.AddItem(tview.NewBox(), 0, 1, false)
    flex.AddItem(tview.NewBox(), 0, 1, false)
    flex.AddItem(tview.NewBox(), 0, 1, false)
    tviewx.NewApplication().SetRoot(flex, true).Run()
}
