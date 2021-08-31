// tviewx/examples/textlinelist
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "math/rand"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "github.com/hshimamoto/tviewx"
)

func main() {
    app := tviewx.NewApplication()

    list1 := tviewx.NewTextLineList()
    list2 := tviewx.NewTextLineList()

    create := func(list *tviewx.TextLineList) *tviewx.TextLineListItem {
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
	    it.ReplaceItem(0, b.GetLabel())
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
    list1.SetHeader(hdr)
    list2.SetHeader(hdr)
    for i := 0; i < 20; i++ {
	it := create(list1)
	list1.AddItem(it)
    }
    for i := 0; i < 20; i++ {
	it := create(list2)
	list2.AddItem(it)
    }
    //
    flex := tview.NewFlex().SetDirection(tview.FlexRow)
    flex.AddItem(list1, 10, 1, true)
    flex.AddItem(list2, 0, 1, true)

    // blur
    list1.SetBlurFunc(func(tcell.Key){
	app.SetFocus(list2)
    })
    list2.SetBlurFunc(func(tcell.Key){
	app.SetFocus(list1)
    })

    app.SetRoot(flex, true)
    app.Run()
}