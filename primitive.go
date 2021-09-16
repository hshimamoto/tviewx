// tviewx.Primitive
// MIT License Copyright(c) 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
)

type BlurFunc interface {
    GetBlurFunc() func(tcell.Key)
    SetBlurFunc(func(tcell.Key))
}

type FocusBackward interface {
    SetFocusBackward(bool)
}
