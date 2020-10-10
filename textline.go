// tviewx.ButtonPanel
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell"
    "github.com/rivo/tview"
)

type textlineItem struct {
    *tview.TextView
    size int
}

type TextLine struct {
    *tview.Box
    items []*textlineItem
    textcolor tcell.Color
    separator string
}

func NewTextLine() *TextLine {
    t := &TextLine{}
    t.Box = tview.NewBox()
    t.items = []*textlineItem{}
    t.separator = string(tview.Borders.Vertical)
    t.textcolor = tview.Styles.PrimaryTextColor
    return t
}

func (t *TextLine)AddText(text string, size int) *TextLine {
    ti := &textlineItem{}
    ti.TextView = tview.NewTextView().SetText(text)
    ti.size = size
    t.items = append(t.items, ti)
    return t
}

func (t *TextLine)GetItem(index int) *tview.TextView {
    if index < 0 || index >= len(t.items) {
	return nil
    }
    return t.items[index].TextView
}

func (t *TextLine)SetTextColor(color tcell.Color) *TextLine {
    t.textcolor = color
    return t
}

func (t *TextLine)SetSeparator(sep string) *TextLine {
    t.separator = sep
    return t
}

func (t *TextLine)Draw(scr tcell.Screen) {
    t.Box.Draw(scr)
    x, y, w, _ := t.GetInnerRect()
    rightLimit := x + w
    cx := x
    tview.Print(scr, t.separator, cx, y, 1, tview.AlignCenter, t.textcolor)
    cx += 1
    for _, ti := range t.items {
	bw := ti.size
	if cx + bw >= rightLimit {
	    bw = rightLimit - cx
	}
	ti.SetRect(cx, y, bw, 1)
	ti.SetTextColor(t.textcolor)
	ti.Draw(scr)
	cx += bw
	if cx >= rightLimit {
	    break
	}
	tview.Print(scr, t.separator, cx, y, 1, tview.AlignCenter, t.textcolor)
	cx += 1
	if cx >= rightLimit {
	    break
	}
    }
}
