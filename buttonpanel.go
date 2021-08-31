// tviewx.ButtonPanel
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type ButtonPanel struct {
    *tview.Box
    buttons []*Button
    hasFocus bool
    focusedButton int
    selected func(*Button)
    lostFocusInLoop bool
    blurFunc func(tcell.Key)
}

func NewButtonPanel() *ButtonPanel {
    bp := &ButtonPanel{}
    bp.Box = tview.NewBox()
    bp.hasFocus = false
    bp.focusedButton = -1
    bp.selected = nil
    bp.lostFocusInLoop = false
    bp.blurFunc = func(tcell.Key){}
    return bp
}

func (bp *ButtonPanel)AddButton(label string) *ButtonPanel {
    b := NewButton(label)
    selected := func() {
	if bp.selected != nil {
	    bp.selected(b)
	}
    }
    b.SetSelectedFunc(selected)
    bp.buttons = append(bp.buttons, b)
    return bp
}

func (bp *ButtonPanel)SetSelectedFunc(selected func(b *Button)) *ButtonPanel {
    bp.selected = selected
    return bp
}

func (bp *ButtonPanel)GetCurrentIndex() int {
    return bp.focusedButton
}

func (bp *ButtonPanel)SetCurrentIndex(index int) {
    if index < 0 || index >= len(bp.buttons) {
	index = -1
    }
    bp.focusedButton = index
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
    bp.hasFocus = false
    if len(bp.buttons) == 0 {
	return
    }
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
		if bp.lostFocusInLoop {
		    // lose focus here
		    bp.hasFocus = false
		    bp.blurFunc(key)
		    return
		}
	    }
	    bp.Focus(delegate)
	case tcell.KeyBacktab:
	    bp.focusedButton--
	    if bp.focusedButton < 0 {
		bp.focusedButton = len(bp.buttons) - 1
		if bp.lostFocusInLoop {
		    // lose focus here
		    bp.hasFocus = false
		    bp.blurFunc(key)
		    return
		}
	    }
	    bp.Focus(delegate)
	case tcell.KeyEscape:
	    if bp.lostFocusInLoop {
		// lose focus here
		bp.hasFocus = false
		bp.blurFunc(key)
		return
	    }
	}
    }
    bp.hasFocus = true
    b := bp.buttons[bp.focusedButton]
    b.SetBlurFunc(handler)
    delegate(b)
}

func (bp *ButtonPanel)HasFocus() bool {
    return bp.hasFocus
}

func (bp *ButtonPanel)SetLostFocusInLoop(enable bool) {
    bp.lostFocusInLoop = enable
}

func (bp *ButtonPanel)SetBlurFunc(handler func(tcell.Key)) {
    bp.blurFunc = handler
}

func (bp *ButtonPanel)GetBlurFunc() func(tcell.Key) {
    return bp.blurFunc
}

func (bp *ButtonPanel)InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
    return bp.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	for _, b := range bp.buttons {
	    if b.HasFocus() {
		if handler := b.InputHandler(); handler != nil {
		    handler(event, setFocus)
		    return
		}
	    }
	}
    })
}
