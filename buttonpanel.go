// tviewx.ButtonPanel
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell"
    "github.com/rivo/tview"
)

type ButtonPanel struct {
    *tview.Box
    buttons []*tview.Button
    hasFocus bool
    focusedButton int
    selected func(*tview.Button)
}

func NewButtonPanel() *ButtonPanel {
    bp := &ButtonPanel{}
    bp.Box = tview.NewBox()
    bp.hasFocus = false
    bp.focusedButton = -1
    bp.selected = nil
    return bp
}

func (bp *ButtonPanel)AddButton(label string) *ButtonPanel {
    b := tview.NewButton(label)
    selected := func() {
	if bp.selected != nil {
	    bp.selected(b)
	}
    }
    b.SetSelectedFunc(selected)
    bp.buttons = append(bp.buttons, b)
    return bp
}

func (bp *ButtonPanel)SetSelectedFunc(selected func(b *tview.Button)) {
    bp.selected = selected
}

func (bp *ButtonPanel)Draw(scr tcell.Screen) {
    bp.Box.Draw(scr)
    x, y, w, _ := bp.GetInnerRect()
    rightLimit := x + w

    cx := x
    for _, b := range bp.buttons {
	bw := tview.TaggedStringWidth(b.GetLabel()) + 4 // [__Label__]
	if cx + bw >= rightLimit {
	    bw = rightLimit - cx
	}
	b.SetRect(cx, y, bw, 1)
	b.Draw(scr)
	cx += bw + 1
	if cx >= rightLimit {
	    break
	}
    }
}

func (bp *ButtonPanel)Focus(delegate func(p tview.Primitive)) {
    if len(bp.buttons) == 0 {
	bp.hasFocus = true
	return
    }
    bp.hasFocus = false
    if bp.focusedButton == -1 {
	// pick the first button
	bp.focusedButton = 0
    }
    handler := func(key tcell.Key) {
	switch key {
	case tcell.KeyTab, tcell.KeyEnter:
	    bp.focusedButton++
	    if bp.focusedButton >= len(bp.buttons) {
		bp.focusedButton = 0
	    }
	    bp.Focus(delegate)
	case tcell.KeyBacktab:
	    bp.focusedButton--
	    if bp.focusedButton < 0 {
		bp.focusedButton = len(bp.buttons) - 1
	    }
	    bp.Focus(delegate)
	}
    }
    b := bp.buttons[bp.focusedButton]
    b.SetBlurFunc(handler)
    delegate(b)
}

func (bp *ButtonPanel)HasFocus() bool {
    if bp.hasFocus {
	return true
    }
    return bp.focusedButton >= 0
}

func (bp *ButtonPanel)InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
    return bp.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	for _, b := range bp.buttons {
	    if b.GetFocusable().HasFocus() {
		if handler := b.InputHandler(); handler != nil {
		    handler(event, setFocus)
		    return
		}
	    }
	}
    })
}

func (bp *ButtonPanel)GetFocusable() tview.Focusable {
    return bp
}
