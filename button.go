// tviewx.ButtonPanel
// MIT License Copyright(c) 2020 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/rivo/tview"
)

type Button struct {
    *tview.Button
}

func NewButton(label string) *Button {
    return &Button{ Button: tview.NewButton(label) }
}
