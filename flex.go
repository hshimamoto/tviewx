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
    focusBackword bool
    blurFunc func(tcell.Key)
}

func NewFlex() *Flex {
    f := &Flex{
	Flex: tview.NewFlex(),
	focused: -1,
    }
    return f
}

func NewFlexRow() *Flex {
    f := NewFlex()
    f.SetDirection(tview.FlexRow)
    return f
}

// override
func (f *Flex)AddItem(item tview.Primitive, fixed, prop int, focus bool) *Flex {
    i := &flexItem{
	Item: item,
	Focus: focus,
	BlurFunc: nil,
    }
    if xp, ok := item.(BlurFunc); ok {
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
	    if i == f.focused {
		// lose focus
		f.focused = -1
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
	if xp, ok := item.Item.(BlurFunc); ok {
	    xp.SetBlurFunc(item.BlurFunc)
	}
    }
    f.items = nil
    f.focused = -1
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
    if f.focused == -1 {
	if f.focusBackword {
	    for i := len(f.items) - 1; i >= 0; i-- {
		item := f.items[i]
		if item.Focus {
		    f.focused = i
		    break
		}
	    }
	} else {
	    for i := 0; i < len(f.items); i++ {
		item := f.items[i]
		if item.Focus {
		    f.focused = i
		    break
		}
	    }
	}
    }
    // still nil?
    if f.focused == -1 {
	return
    }
    onBlur := func(key tcell.Key) {
	focus := func(i int, b bool) {
	    f.focused = i
	    item := f.items[i].Item
	    // backward?
	    if fb, ok := item.(FocusBackward); ok {
		fb.SetFocusBackward(b)
	    }
	    f.Focus(delegate)
	}
	sz := len(f.items)
	move := func(i int, key tcell.Key) {
	    cur := f.focused
	    n := (cur + i + sz) % sz
	    for n != cur {
		if (n == 0 && i == 1) || (n == 0 && i == -1) {
		    if f.tryBlur(key) {
			return
		    }
		}
		if f.items[n].Focus {
		    focus(n, (i == -1))
		    return
		}
		// next
		n = (n + i + sz) % sz
	    }
	    // no other items to get focus
	    focus(cur, false)
	}
	switch key {
	case tcell.KeyTab, tcell.KeyEnter: // forward
	    move(1, key)
	case tcell.KeyBacktab: // backword
	    move(-1, key)
	case tcell.KeyEscape: // just lost focus
	    if f.tryBlur(key) {
		return
	    }
	    // back to current one
	    focus(f.focused, false)
	}
    }
    if f.focused == -1 {
	// nothing to do
	return
    }
    item := f.items[f.focused].Item
    if xp, ok := item.(BlurFunc); ok {
	xp.SetBlurFunc(onBlur)
    }
    delegate(item)
}

func (f *Flex)HasFocus() bool {
    if f.focused == -1 {
	return false
    }
    return f.items[f.focused].Item.HasFocus()
}

func (f *Flex)GetBlurFunc() func(tcell.Key) {
    return f.blurFunc
}

func (f *Flex)SetBlurFunc(handler func(tcell.Key)) {
    f.blurFunc = handler
}

func (f *Flex)tryBlur(key tcell.Key) bool {
    if f.blurFunc == nil {
	return false
    }
    f.focused = -1
    f.blurFunc(key)
    return true
}

func (f *Flex)SetFocusBackward(b bool) {
    f.focusBackword = b
}
