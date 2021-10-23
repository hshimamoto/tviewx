// tviewx.ButtonPanel
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type textlineItem struct {
    *TextView
    size int
}

type TextLine struct {
    *tview.Box
    items []*textlineItem
    textcolor tcell.Color
    separator string
    dynamic func(int, string) string
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
    ti.TextView = NewTextView().SetText(text)
    ti.size = size
    t.items = append(t.items, ti)
    return t
}

func (t *TextLine)GetItem(index int) *TextView {
    if index < 0 || index >= len(t.items) {
	return nil
    }
    return t.items[index].TextView
}

func (t *TextLine)ReplaceText(index int, text string) *TextLine {
    if index < 0 || index >= len(t.items) {
	return t
    }
    t.items[index].TextView.SetText(text)
    return t
}

func (t *TextLine)ReplaceTexts(index int, texts []string) *TextLine {
    for i, text := range texts {
	t.ReplaceText(index + i, text)
    }
    return t
}

func (t *TextLine)SetTextColor(color tcell.Color) *TextLine {
    t.textcolor = color
    return t
}

func (t *TextLine)SetSeparator(sep string) *TextLine {
    t.separator = sep
    return t
}

func (t *TextLine)GetDynamic() func(int, string) string {
    return t.dynamic
}

func (t *TextLine)SetDynamic(dyn func(int, string) string) *TextLine {
    t.dynamic = dyn
    return t
}

func (t *TextLine)Draw(scr tcell.Screen) {
    t.Box.Draw(scr)
    x, y, w, _ := t.GetInnerRect()
    rightLimit := x + w
    cx := x
    tview.Print(scr, t.separator, cx, y, 1, tview.AlignCenter, t.textcolor)
    cx += 1
    for i, ti := range t.items {
	bw := ti.size
	if cx + bw >= rightLimit - 1 {
	    bw = rightLimit - 1 - cx
	}
	ti.SetRect(cx, y, bw, 1)
	ti.SetTextColor(t.textcolor)
	if t.dynamic != nil {
	    orig := ti.GetText(false)
	    ti.SetText(t.dynamic(i, orig))
	    ti.Draw(scr)
	    ti.SetText(orig)
	} else {
	    ti.Draw(scr)
	}
	cx += bw
	tview.Print(scr, t.separator, cx, y, 1, tview.AlignCenter, t.textcolor)
	if cx >= rightLimit {
	    break
	}
	cx += 1
	if cx >= rightLimit {
	    break
	}
    }
}
