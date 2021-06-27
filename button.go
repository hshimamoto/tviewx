// tviewx.ButtonPanel
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type Button struct {
    *tview.Button
}

func NewButton(label string) *Button {
    return &Button{ Button: tview.NewButton(label) }
}

func (b *Button)InputHandler() func(*tcell.EventKey, func(tview.Primitive)) {
    return b.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	e := event
	switch key := event.Key(); key {
	case tcell.KeyLeft:
	    e = tcell.NewEventKey(tcell.KeyBacktab, event.Rune(), event.Modifiers())
	case tcell.KeyRight:
	    e = tcell.NewEventKey(tcell.KeyTab, event.Rune(), event.Modifiers())
	}
	b.Button.InputHandler()(e, setFocus)
    })
}
