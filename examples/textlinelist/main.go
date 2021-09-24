// tviewx/examples/textlinelist
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
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
	    it.ReplaceText(0, b.GetLabel())
	    list.CloseMenu()
	})
	return it
    }
    hdr := tviewx.NewTextLine()
    hdr.AddText("Header0", 20)
    hdr.AddText("Header1", 10)
    hdr.AddText("Header2", 10)
    hdr.AddText("Header3", 20)
    hdr.AddText("Header4", 20)
    list.SetHeader(hdr)
    for i := 0; i < 20; i++ {
	it := create()
	list.AddItem(it)
    }
    list.GetItem(0).ReplaceTexts(0, []string{"First", "FIRST", "first"})
    list.GetItem(1).ReplaceTexts(0, []string{"Second", "SECOND", "second"})
    //
    flex := tviewx.NewFlexRow()
    flex.AddItem(list, 10, 1, true)
    flex.AddItem(tview.NewBox(), 0, 1, false)
    flex.AddItem(tview.NewBox(), 0, 1, false)
    flex.AddItem(tview.NewBox(), 0, 1, false)
    tviewx.NewApplication().SetRoot(flex, true).Run()
}
