// tviewx.TextView
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type TextView struct {
    *tview.TextView
    blurFunc func(tcell.Key)
}

func NewTextView() *TextView {
    return &TextView{ TextView: tview.NewTextView() }
}

func (t *TextView)Clear() *TextView {
    t.TextView.Clear()
    return t
}

func (t *TextView)Highlight(ids ...string) *TextView {
    t.TextView.Highlight(ids...)
    return t
}

func (t *TextView)ScrollTo(row, col int) *TextView {
    t.TextView.ScrollTo(row, col)
    return t
}

func (t *TextView)ScrollToBeginning() *TextView {
    t.TextView.ScrollToBeginning()
    return t
}

func (t *TextView)ScrollToEnd() *TextView {
    t.TextView.ScrollToEnd()
    return t
}

func (t *TextView)ScrollToHighlight() *TextView {
    t.TextView.ScrollToHighlight()
    return t
}

func (t *TextView)SetChangedFunc(handler func()) *TextView {
    t.TextView.SetChangedFunc(handler)
    return t
}

func (t *TextView)SetDoneFunc(handler func(tcell.Key)) *TextView {
    t.TextView.SetDoneFunc(handler)
    return t
}

func (t *TextView)SetDynamicColors(d bool) *TextView {
    t.TextView.SetDynamicColors(d)
    return t
}

func (t *TextView)SetHighlightedFunc(handler func(a, b, c []string)) *TextView {
    t.TextView.SetHighlightedFunc(handler)
    return t
}

func (t *TextView)SetMaxLines(n int) *TextView {
    t.TextView.SetMaxLines(n)
    return t
}

func (t *TextView)SetRegions(r bool) *TextView {
    t.TextView.SetRegions(r)
    return t
}

func (t *TextView)SetScrollable(s bool) *TextView {
    t.TextView.SetScrollable(s)
    return t
}

func (t *TextView)SetText(s string) *TextView {
    t.TextView.SetText(s)
    return t
}

func (t *TextView)SetTextAlign(a int) *TextView {
    t.TextView.SetTextAlign(a)
    return t
}

func (t *TextView)SetTextColor(c tcell.Color) *TextView {
    t.TextView.SetTextColor(c)
    return t
}

func (t *TextView)SetWordWrap(w bool) *TextView {
    t.TextView.SetWordWrap(w)
    return t
}

func (t *TextView)SetWrap(w bool) *TextView {
    t.TextView.SetWrap(w)
    return t
}

func (t *TextView)SetToggleHighlights(toggle bool) *TextView {
    t.TextView.SetToggleHighlights(toggle)
    return t
}

func (t *TextView)GetBlurFunc() func(tcell.Key) {
    return t.blurFunc
}

func (t *TextView)SetBlurFunc(handle func(tcell.Key)) {
    t.blurFunc = handle
}
