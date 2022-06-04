// tviewx.TextLineList
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type TextLineListItem struct {
    *TextLine
    menu *ButtonPanel
    selected func(b *Button)
}

func NewTextLineListItem() *TextLineListItem {
    ti := &TextLineListItem{}
    ti.TextLine = NewTextLine()
    ti.selected = nil
    ti.menu = NewButtonPanel().SetSelectedFunc(ti.defSelected)
    return ti
}

func (ti *TextLineListItem)AddText(text string, size int) *TextLineListItem {
    ti.TextLine.AddText(text, size)
    return ti
}

func (ti *TextLineListItem)AddButton(label string) *TextLineListItem {
    ti.menu.AddButton(label)
    return ti
}

func (ti *TextLineListItem)defSelected(b *Button) {
    if ti.selected != nil {
	ti.selected(b)
    }
}

func (ti *TextLineListItem)SetSelectedFunc(selected func(b *Button)) *TextLineListItem {
    ti.selected = selected
    return ti
}

type TextLineList struct {
    *tview.Box
    header *TextLine
    items []*TextLineListItem
    cur, last int
    drawst int
    open bool
    hasFocus bool
    blurFunc func(tcell.Key)
    changed func(*TextLineListItem)
    scrollup string
    scrolldown string
}

func NewTextLineList() *TextLineList {
    tl := &TextLineList{}
    tl.Box = tview.NewBox()
    tl.items = []*TextLineListItem{}
    tl.cur = 0
    tl.last = -1
    tl.drawst = 0
    tl.open = false
    tl.hasFocus = false
    tl.blurFunc = func(tcell.Key){}
    tl.changed = nil
    tl.scrollup = "\u25b2"
    tl.scrolldown = "\u25bc"
    return tl
}

func (tl *TextLineList)GetHeader() *TextLine {
    return tl.header
}

func (tl *TextLineList)SetHeader(header *TextLine) *TextLineList {
    tl.header = header
    return tl
}

func (tl *TextLineList)AddItem(item *TextLineListItem) *TextLineList {
    tl.items = append(tl.items, item)
    tl.last++
    return tl
}

func (tl *TextLineList)GetItem(index int) *TextLineListItem {
    if index < 0 || index > tl.last {
	return nil
    }
    return tl.items[index]
}

func (tl *TextLineList)OpenMenu() {
    tl.open = true
    tl.items[tl.cur].menu.SetCurrentIndex(0)
}

func (tl *TextLineList)CloseMenu() {
    tl.open = false
}

func (tl *TextLineList)CursorUp() {
    tl.cur--
    if tl.cur < 0 {
	tl.cur = 0
    }
}

func (tl *TextLineList)CursorDown() {
    tl.cur++
    if tl.cur > tl.last {
	tl.cur = tl.last
    }
}

func (tl *TextLineList)CursorTop() {
    tl.cur = 0
}

func (tl *TextLineList)CursorBottom() {
    tl.cur = tl.last
}

func (tl *TextLineList)InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
    return tl.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	if len(tl.items) == 0 {
	    // nothing to do if there is no item
	    return
	}
	prevopen := tl.open
	defer func() {
	    if tl.open != prevopen {
		if tl.open {
		    setFocus(tl.items[tl.cur].menu)
		} else {
		    setFocus(tl)
		}
	    }
	}()
	menu := tl.items[tl.cur].menu
	if tl.open {
	    if event.Key() == tcell.KeyEscape {
		tl.CloseMenu()
		return
	    }
	    if handler := menu.InputHandler(); handler != nil {
		handler(event, setFocus)
	    }
	    return
	}
	prev := tl.cur
	switch key := event.Key(); key {
	case tcell.KeyEnter: tl.OpenMenu()
	case tcell.KeyUp: tl.CursorUp()
	case tcell.KeyDown: tl.CursorDown()
	case tcell.KeyHome: tl.CursorTop()
	case tcell.KeyEnd: tl.CursorBottom()
	case tcell.KeyEscape, tcell.KeyTab, tcell.KeyBacktab:
	    tl.hasFocus = false
	    tl.blurFunc(key)
	}
	switch event.Rune() {
	case ' ': tl.OpenMenu()
	case 'k': tl.CursorUp()
	case 'j': tl.CursorDown()
	case 'g': tl.CursorTop()
	case 'G': tl.CursorBottom()
	}

	// selected item changed
	if prev != tl.cur {
	    if tl.changed != nil {
		tl.changed(tl.items[tl.cur])
	    }
	}
    })
}

func (tl *TextLineList)Focus(delegate func(tview.Primitive)) {
    tl.hasFocus = true
}

func (tl *TextLineList)HasFocus() bool {
    return tl.hasFocus
}

func (tl *TextLineList)SetBlurFunc(handler func(tcell.Key)) {
    tl.blurFunc = handler
}

func (tl *TextLineList)GetBlurFunc() func(tcell.Key) {
    return tl.blurFunc
}

func (tl *TextLineList)SetChangedFunc(handler func(*TextLineListItem)) {
    tl.changed = handler
}

func (tl *TextLineList)GetChangedFunc() func(*TextLineListItem) {
    return tl.changed
}

func (tl *TextLineList)Draw(scr tcell.Screen) {
    tl.Box.Draw(scr)
    if len(tl.items) == 0 {
	// no items
	return
    }
    x, y, w, h := tl.GetInnerRect()
    x += 1
    w -= 2 // for cursor and scroll bar
    h -= 1 // for button
    hdr := 0
    if tl.header != nil {
	// draw header here
	tl.header.SetRect(x, y, w, 1)
	tl.header.Draw(scr)
	y += 1
	hdr = 1
    }
    if h - hdr > len(tl.items) {
	h = len(tl.items) + hdr
    }
    top := y
    btm := y + h - hdr
    // check lines for draw
    if tl.cur < tl.drawst {
	tl.drawst = tl.cur
    } else if tl.cur >= tl.drawst + (h - hdr) - 1 {
	tl.drawst = tl.cur - (h - hdr) + 1
    }
    for i, item := range tl.items {
	if i < tl.drawst {
	    continue
	}
	// cursor> textline
	if i == tl.cur {
	    cursor := "\u25b7"
	    if tl.hasFocus {
		cursor = "\u25b6"
		if tl.open {
		    cursor = "\u25bd"
		}
	    }
	    PrintR(scr, cursor, x-1, y, 1)
	}
	item.SetRect(x, y, w, 1)
	item.Draw(scr)
	y++
	if y >= btm {
	    break
	}
    }
    // scroll bar
    if tl.drawst > 0 {
	PrintR(scr, tl.scrollup, x+w, top, 1)
    }
    if tl.drawst <= tl.last - (h - hdr) {
	PrintR(scr, tl.scrolldown, x+w, btm-1, 1)
    }
    sptr := int(((h - 1) * (tl.cur + 1)) / len(tl.items)) - 1
    Dbg.Printf(
	"TextLineList<%p>.Draw: cur=%d/%d *%d => %d",
	tl, tl.cur + 1, len(tl.items), h-1, sptr)
    if tl.cur == 0 {
	PrintR(scr, "*", x+w, top, 1)
    } else if tl.cur == (len(tl.items) - 1) {
	PrintR(scr, "*", x+w, btm-1, 1)
    } else {
	if sptr <= 0 {
	    sptr = 1
	} else if sptr == (h-2) {
	    sptr = h - 3
	}
	PrintR(scr, "*", x+w, top+sptr, 1)
    }
    // menu buttons
    menu := tl.items[tl.cur].menu
    menu.SetRect(x, btm, w, 1)
    menu.Draw(scr)
    // for debug
    Dbg.Printf("TextLineList<%p>.Draw: cur=%d", tl, tl.cur)
}
