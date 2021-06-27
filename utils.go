// tviewx.*
// MIT License Copyright(c) 2020, 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

func PrintL(scr tcell.Screen, str string, x, y, w int) (int, int) {
    return tview.Print(scr, str, x, y, w, tview.AlignLeft, tview.Styles.PrimaryTextColor)
}

func PrintR(scr tcell.Screen, str string, x, y, w int) (int, int) {
    return tview.Print(scr, str, x, y, w, tview.AlignRight, tview.Styles.PrimaryTextColor)
}

func PrintC(scr tcell.Screen, str string, x, y, w int) (int, int) {
    return tview.Print(scr, str, x, y, w, tview.AlignCenter, tview.Styles.PrimaryTextColor)
}
