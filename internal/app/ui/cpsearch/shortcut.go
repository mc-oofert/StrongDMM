package cpsearch

import (
	"sdmm/internal/app/ui/shortcut"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *Search) addShortcuts() {
	s.shortcuts.Add(shortcut.Shortcut{
		Name:        "cpsearch#jumpToUp",
		FirstKey:    ebiten.KeyShiftLeft,
		FirstKeyAlt: ebiten.KeyShiftRight,
		SecondKey:   ebiten.KeyF3,
		Action:      s.jumpToUp,
	})
	s.shortcuts.Add(shortcut.Shortcut{
		Name:     "cpsearch#jumpToDown",
		FirstKey: ebiten.KeyF3,
		Action:   s.jumpToDown,
	})
	s.shortcuts.Add(shortcut.Shortcut{
		Name:     "cpseaarch#doToggleFilter",
		FirstKey: ebiten.KeyF,
		Action:   s.doToggleFilter,
	})
}
