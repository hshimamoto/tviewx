// tviewx.Application
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type Application struct {
    *tview.Application
    // optional capture for user
    inputCapture func(event *tcell.EventKey) *tcell.EventKey
    // for tviewx custom capture
    numCtrlC int
    ctrlCFunc func()
}

func NewApplication() *Application {
    a := &Application{
	Application: tview.NewApplication(),
	inputCapture: nil,
	numCtrlC: 0,
    }
    // set custom capture
    a.Application.SetInputCapture(a.xInputCapture)
    return a
}

func (a *Application)GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
    return a.inputCapture
}

func (a *Application)SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Application {
    a.inputCapture = capture
    return a
}

func (a *Application)GetCtrlCFunc() func() {
    return a.ctrlCFunc
}

func (a *Application)SetCtrlCFunc(fn func()) {
    a.ctrlCFunc = fn
}

// Application inputCapture
func (a *Application)xInputCapture(event *tcell.EventKey) *tcell.EventKey {
    if a.inputCapture != nil {
	event = a.inputCapture(event)
	if event == nil {
	    return nil
	}
    }
    // Block Ctrl-C
    if event.Key() == tcell.KeyCtrlC {
	if a.ctrlCFunc != nil {
	    a.ctrlCFunc()
	}
	a.numCtrlC++
	if a.numCtrlC < 3 {
	    return nil
	}
    } else {
	a.numCtrlC = 0
    }
    return event
}
