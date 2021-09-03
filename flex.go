// tviewx.Flex
// MIT License Copyright(c) 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type flexItem struct {
    Item Primitive
    FixedSize, Proportion int
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
func (f *Flex)AddItem(item Primitive, fixed, prop int, focus bool) *Flex {
    i := &flexItem{
	Item: item,
	FixedSize: fixed,
	Proportion: prop,
	Focus: focus,
	BlurFunc: item.GetBlurFunc(),
    }
    f.items = append(f.items, i)

    // call super
    f.Flex.AddItem(item, fixed, prop, focus)
    return f
}

func (f *Flex)RemoveItem(p Primitive) *Flex {
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
	item.Item.SetBlurFunc(item.BlurFunc)
    }
    f.items = nil
    f.focus = nil
    // call super
    f.Flex.Clear()
    return f
}

func (f *Flex)ResizeItem(p Primitive, fixed, prop int) *Flex {
    for _, item := range f.items {
	if item.Item == p {
	    item.FixedSize = fixed
	    item.Proportion = prop
	}
    }
    // call super
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
	switch key {
	case tcell.KeyTab, tcell.KeyEnter: // forward
	    for i := f.focused + 1; i < sz; i++ {
		if f.items[i].Focus {
		    focus(i)
		    return
		}
	    }
	    // not found
	    if f.blurFunc != nil {
		// call blur and end
		f.focus = nil
		f.blurFunc(key)
		return
	    }
	    // lookup from head
	    for i := 0; i < sz; i++ {
		if f.items[i].Focus {
		    focus(i)
		    return
		}
	    }
	case tcell.KeyBacktab: // backword
	    for i := f.focused - 1; i >= 0; i-- {
		if f.items[i].Focus {
		    focus(i)
		    return
		}
	    }
	    // not found
	    if f.blurFunc != nil {
		// call blur and end
		f.focus = nil
		f.blurFunc(key)
		return
	    }
	    // lookup from head
	    for i := sz - 1; i > f.focused; i-- {
		if f.items[i].Focus {
		    focus(i)
		    return
		}
	    }
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
    f.focus.Item.SetBlurFunc(onBlur)
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
