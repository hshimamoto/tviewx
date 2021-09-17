// tviewx.TabPanel
// MIT License Copyright(c) 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "fmt"
    "strings"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type TabPanelItem struct {
    Item tview.Primitive
    Name string
}

type TabPanel struct {
    *tview.Box
    items []TabPanelItem
    cur int
    tabFocused bool
    hasFocus bool
    focusBackword bool
    blurFunc func(tcell.Key)
    tabOnTop bool
}

func NewTabPanel() *TabPanel {
    tp := &TabPanel{
	Box: tview.NewBox(),
	tabFocused: true,
    }
    return tp
}

func (tp *TabPanel)SetTabLocationTop(top bool) *TabPanel {
    tp.tabOnTop = top
    return tp
}

func (tp *TabPanel)AddItem(name string, p tview.Primitive) *TabPanel {
    item := TabPanelItem{
	Item: p,
	Name: name,
    }
    tp.items = append(tp.items, item)
    return tp
}

func (tp *TabPanel)Draw(scr tcell.Screen) {
    tp.Box.Draw(scr)
    x, y, w, h := tp.GetInnerRect()
    //Dbg.Printf("TabPanel<%p>.Draw: %d %d %d %d", tp, x, y, w, h)
    if h < 2 {
	// no enough space to locate tab
	if h <= 0 {
	    Dbg.Printf("NO SPACE")
	    return
	}
	PrintL(scr, "No space", x, y, 8)
	return
    }
    if len(tp.items) == 0 {
	PrintL(scr, "No items", x, y, 8)
	return
    }
    item_y := y
    tab_y := y + h - 1
    if tp.tabOnTop {
	item_y = y + 1
	tab_y = y
    }
    // draw items and collect tabs
    tabs := []string{}
    for i, it := range tp.items {
	tab := it.Name
	if i == tp.cur {
	    if tp.hasFocus {
		if tp.tabFocused {
		    tab = fmt.Sprintf("\u25c0%s\u25b6", it.Name)
		} else {
		    if tp.tabOnTop {
			tab = fmt.Sprintf("\u25bd%s\u25bd", it.Name)
		    } else {
			tab = fmt.Sprintf("\u25b3%s\u25b3", it.Name)
		    }
		}
	    } else {
		tab = fmt.Sprintf("\u25c1%s\u25b7", it.Name)
	    }
	    // draw item
	    it.Item.SetRect(x, item_y, w, h-1)
	    it.Item.Draw(scr)
	}
	tabs = append(tabs, tab)
    }
    tabstr := strings.Join(tabs, " ")
    PrintL(scr, tabstr, x+1, tab_y, len(tabstr))
}

func (tp *TabPanel)Focus(delegate func(tview.Primitive)) {
    tp.hasFocus = false
    if len(tp.items) == 0 {
	// unable to have focus
	return
    }
    tp.hasFocus = true
}

func (tp *TabPanel)HasFocus() bool {
    return tp.hasFocus
}

func (tp *TabPanel)tabForward() {
    if tp.cur < (len(tp.items) - 1) {
	tp.cur++
    }
}

func (tp *TabPanel)tabBackward() {
    if tp.cur > 0 {
	tp.cur--
    }
}

func (tp *TabPanel)tabEnter(setFocus func(tview.Primitive)) {
    //Dbg.Printf("TabPanel<%p>.tabEnter: cur=%d", tp, tp.cur)
    tp.tabFocused = false
    p := tp.items[tp.cur].Item
    if xp, ok := p.(BlurFunc); ok {
	//Dbg.Printf("TabPanel<%p>.tabEnter: setBlur", tp)
	xp.SetBlurFunc(func(tcell.Key) {
	    //Dbg.Printf("TabPanel<%p>.tabEnter: invoke BlurFunc", tp)
	    tp.tabFocused = true
	    setFocus(tp)
	})
    }
    setFocus(p)
}

func (tp *TabPanel)InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
    return tp.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	if tp.tabFocused {
	    // left and right: change tab
	    // tab, space and enter: focus into item
	    switch key := event.Key(); key {
	    case tcell.KeyLeft: tp.tabBackward()
	    case tcell.KeyRight: tp.tabForward()
	    case tcell.KeyEnter: tp.tabEnter(setFocus)
	    case tcell.KeyEscape, tcell.KeyTab, tcell.KeyBacktab:
		tp.hasFocus = false
		tp.blurFunc(key)
	    }
	    switch event.Rune() {
	    case 'h': tp.tabBackward()
	    case 'l': tp.tabForward()
	    case ' ': tp.tabEnter(setFocus)
	    }
	    return
	}
	// item focused
	if tp.cur == -1 {
	    Dbg.Printf("TabPanel<%p>.InputHandler: Unknown condition", tp)
	    return
	}
	if handler := tp.items[tp.cur].Item.InputHandler(); handler != nil {
	    handler(event, setFocus)
	    return
	}
    })
}

func (tp *TabPanel)GetBlurFunc() func(tcell.Key) {
    return tp.blurFunc
}

func (tp *TabPanel)SetBlurFunc(handler func(tcell.Key)) {
    tp.blurFunc = handler
}
