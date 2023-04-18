package cpenvironment

import (
	"sdmm/internal/app/ui/shortcut"

	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Environment) addShortcuts() {
	e.shortcuts.Add(shortcut.Shortcut{
		Name:     "cpenvironment#doToggleTypesFilter",
		FirstKey: ebiten.KeyF,
		Action:   e.doToggleTypesFilter,
	})
}
