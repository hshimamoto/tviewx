// tviewx.Flex
// MIT License Copyright(c) 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type flexItem struct {
    Item tview.Primitive
    Focus bool
    BlurFunc func(tcell.Key)
}

type Flex struct {
    *tview.Flex
    items []*flexItem
    focused int
    focus *flexItem
    blurFunc func(tcell.Key)
}

func NewFlex() *Flex {
    f := &Flex{
	Flex: tview.NewFlex(),
    }
    return f
}

// override
func (f *Flex)AddItem(item tview.Primitive, fixed, prop int, focus bool) *Flex {
    i := &flexItem{
	Item: item,
	Focus: focus,
	BlurFunc: nil,
    }
    if xp, ok := item.(Primitive); ok {
	i.BlurFunc = xp.GetBlurFunc()
    }
    f.items = append(f.items, i)

    // call super
    f.Flex.AddItem(item, fixed, prop, focus)
    return f
}

func (f *Flex)RemoveItem(p tview.Primitive) *Flex {
    for i := len(f.items) - 1; i >= 0; i-- {
	if f.items[i].Item == p {
	    if f.items[i] == f.focus {
		// lose focus
		f.focus = nil
	    }
	    f.items = append(f.items[:i], f.items[i+1:]...)
	}
    }

    // call super
    f.Flex.RemoveItem(p)
    return f
}

func (f *Flex)Clear() *Flex {
    // reset BlurFunc
    for _, item := range f.items {
	if xp, ok := item.Item.(Primitive); ok {
	    xp.SetBlurFunc(item.BlurFunc)
	}
    }
    f.items = nil
    f.focus = nil
    // call super
    f.Flex.Clear()
    return f
}

func (f *Flex)ResizeItem(p tview.Primitive, fixed, prop int) *Flex {
    f.Flex.ResizeItem(p, fixed, prop)
    return f
}

func (f *Flex)SetDirection(d int) *Flex {
    f.Flex.SetDirection(d)
    return f
}

func (f *Flex)SetFullScreen(b bool) *Flex {
    f.Flex.SetFullScreen(b)
    return f
}

func (f *Flex)Focus(delegate func(p tview.Primitive)) {
    // there is no focus
    if f.focus == nil {
	for i := 0; i < len(f.items); i++ {
	    item := f.items[i]
	    if item.Focus {
		f.focus = item
		f.focused = i
		break
	    }
	}
    }
    // still nil?
    if f.focus == nil {
	return
    }
    onBlur := func(key tcell.Key) {
	focus := func(i int) {
	    f.focus = f.items[i]
	    f.focused = i
	    f.Focus(delegate)
	}
	sz := len(f.items)
	move := func(i int, key tcell.Key) {
	    cur := f.focused
	    n := (cur + i + sz) % sz
	    for n != cur {
		if (n == 0 && i == 1) || (n == (sz - 1) && i == -1) {
		    if f.blurFunc != nil {
			f.focus = nil
			f.blurFunc(key)
			return
		    }
		}
		if f.items[n].Focus {
		    focus(n)
		    return
		}
		// next
		n = (n + i + sz) % sz
	    }
	}
	switch key {
	case tcell.KeyTab, tcell.KeyEnter: // forward
	    move(1, key)
	case tcell.KeyBacktab: // backword
	    move(-1, key)
	case tcell.KeyEscape: // just lost focus
	    if f.blurFunc != nil {
		f.focus = nil
		f.blurFunc(key)
		return
	    }
	    // back to current one
	    focus(f.focused)
	}
    }
    if f.focus == nil {
	// nothing to do
	return
    }
    if xp, ok := f.focus.Item.(Primitive); ok {
	xp.SetBlurFunc(onBlur)
    }
    delegate(f.focus.Item)
}

func (f *Flex)HasFocus() bool {
    return f.focus != nil && f.focus.Item.HasFocus()
}

func (f *Flex)GetBlurFunc() func(tcell.Key) {
    return f.blurFunc
}

func (f *Flex)SetBlurFunc(handler func(tcell.Key)) {
    f.blurFunc = handler
}
