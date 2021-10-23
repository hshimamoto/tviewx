// tviewx/examples/tabpanel
// MIT License Copyright(c) 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package main

import (
    "fmt"

    "github.com/hshimamoto/tviewx"
)

func main() {
    tviewx.Dbg.Enable()

    app := tviewx.NewApplication()

    w, _ := tviewx.GetCurrentScreenSize()
    header := tviewx.NewTextLine()
    header.AddText("Header", w)

    tabpanel := tviewx.NewTabPanel()
    flex := tviewx.NewFlexRow()
    flex.SetBorder(true)
    flex.AddItem(header, 1, 0, false)
    flex.AddItem(tabpanel, 0, 1, true)

    tabpanel.AddItem("Tab1", tviewx.NewButton("Tab1"))
    tabpanel.AddItem("Tab2", tviewx.NewButton("Tab2"))
    tabpanel.AddItem("Tab3", tviewx.NewButton("Tab3"))
    tabpanel.AddItem("Tab4", tviewx.NewButton("Tab4"))

    in_tabpanel := tviewx.NewTabPanel().SetTabLocationTop(true)
    in_tabpanel.AddItem("TAB1", tviewx.NewButton("Button1"))
    in_tabpanel.AddItem("TAB2", tviewx.NewButton("Button2"))
    in_tabpanel.AddItem("TAB3", tviewx.NewButton("Button3"))

    tabpanel.AddItem("Tab5", in_tabpanel)

    inin_flex := tviewx.NewFlex()

    inin_tabpanel1 := tviewx.NewTabPanel()
    inin_tabpanel1.AddItem("tab1", tviewx.NewButton("X"))
    inin_tabpanel1.AddItem("tab2", tviewx.NewButton("Y"))
    inin_tabpanel1.AddItem("tab3", tviewx.NewButton("Z"))

    inin_tabpanel2 := tviewx.NewTabPanel()
    inin_tabpanel2.AddItem("tabA", tviewx.NewButton("x"))
    inin_tabpanel2.AddItem("tabB", tviewx.NewButton("y"))
    inin_tabpanel2.AddItem("tabC", tviewx.NewButton("z"))

    inin_flex.AddItem(inin_tabpanel1, 0, 1, true)
    inin_flex.AddItem(inin_tabpanel2, 0, 1, true)

    in_tabpanel.AddItem("TAB4", inin_flex)

    in_tabpanel.AddItem("TAB5", tviewx.NewButton("Button5"))
    in_tabpanel.AddItem("TAB6", tviewx.NewButton("Button6"))

    // count
    // there are 5 tabs "Tab1" to "Tab5"
    counts := [5]int{0, 0, 0, 0, 0}
    tabpanel.SetTabChangedFunc(func(idx int, name string, item tviewx.Primitive) {
	counts[idx]++
    })

    // dynamic header change
    header.SetDynamic(func(idx int, orig string) string {
	if idx != 0 {
	    return orig
	}
	name, _ := tabpanel.GetCurrentItem()
	return fmt.Sprintf("Current:%s %d %d %d %d %d",
		name,
		counts[0], counts[1], counts[2], counts[3], counts[4])
    })

    app.SetRoot(flex, true)

    app.Run()

    // show debugging info
    for _, line := range tviewx.Dbg.Lines() {
	fmt.Println(line)
    }
}
