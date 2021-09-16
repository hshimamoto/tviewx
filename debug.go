// tviewx.Primitive
// MIT License Copyright(c) 2021 Hiroshi Shimamoto
// vim: set sw=4 sts=4:
package tviewx

import (
    "fmt"
)

// global variable
var Dbg *Debug = &Debug{
    enabled: false,
    lines: []string{},
}

type Debug struct {
    enabled bool
    lines []string
}

func (dbg *Debug)Enable() {
    dbg.enabled = true
}

func (dbg *Debug)Disable() {
    dbg.enabled = false
}

func (dbg *Debug)Printf(f string, a ...interface{}) {
    if ! dbg.enabled {
	return
    }
    dbg.lines = append(dbg.lines, fmt.Sprintf(f, a...))
}

func (dbg *Debug)Lines() []string {
    return dbg.lines
}
