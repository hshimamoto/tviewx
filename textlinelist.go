// tviewx.TextLineList
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell"
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
    items []*TextLineListItem
    cur, last int
    drawst int
    open bool
}

func NewTextLineList() *TextLineList {
    tl := &TextLineList{}
    tl.Box = tview.NewBox()
    tl.items = []*TextLineListItem{}
    tl.cur = 0
    tl.last = -1
    tl.drawst = 0
    tl.open = false
    return tl
}

func (tl *TextLineList)AddItem(item *TextLineListItem) *TextLineList {
    tl.items = append(tl.items, item)
    tl.last++
    return tl
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
	switch event.Key() {
	case tcell.KeyEnter: tl.OpenMenu()
	case tcell.KeyUp: tl.CursorUp()
	case tcell.KeyDown: tl.CursorDown()
	case tcell.KeyHome: tl.CursorTop()
	case tcell.KeyEnd: tl.CursorBottom()
	}
	switch event.Rune() {
	case ' ': tl.OpenMenu()
	case 'k': tl.CursorUp()
	case 'j': tl.CursorDown()
	case 'g': tl.CursorTop()
	case 'G': tl.CursorBottom()
	}
    })
}

func (tl *TextLineList)GetFocusable() tview.Focusable {
    return tl
}

func (tl *TextLineList)HasFocus() bool {
    return true
}

func (tl *TextLineList)Draw(scr tcell.Screen) {
    tl.Box.Draw(scr)
    x, y, w, h := tl.GetInnerRect()
    x += 1
    w -= 2 // for cursor and scroll bar
    h -= 1 // for button
    // check lines for draw
    if tl.cur < tl.drawst {
	tl.drawst = tl.cur
    } else if tl.cur >= tl.drawst + h - 1 {
	tl.drawst = tl.cur - h + 1
    }
    for i, item := range tl.items {
	if i < tl.drawst {
	    continue
	}
	// cursor> textline
	if i == tl.cur {
	    cursor := ">"
	    if tl.open {
		cursor = "V"
	    }
	    tview.Print(scr, cursor, x-1, y, 1, tview.AlignRight, tview.Styles.PrimaryTextColor)
	}
	item.SetRect(x, y, w, 1)
	item.Draw(scr)
	y++
	if y >= h {
	    break
	}
    }
    // scroll bar
    if tl.drawst > 0 {
	tview.Print(scr, "^", w+1, 0, 1, tview.AlignRight, tview.Styles.PrimaryTextColor)
    }
    if tl.drawst <= tl.last - h {
	tview.Print(scr, "V", w+1, h-1, 1, tview.AlignRight, tview.Styles.PrimaryTextColor)
    }
    // menu buttons
    x--
    w++
    menu := tl.items[tl.cur].menu
    menu.SetRect(x, h, w, 1)
    menu.Draw(scr)
    // for debug
    // tview.Print(scr, fmt.Sprintf("%2d %2d %2d", tl.cur, tl.drawst, h), w-12, h, 11, tview.AlignRight, tview.Styles.PrimaryTextColor)
}
