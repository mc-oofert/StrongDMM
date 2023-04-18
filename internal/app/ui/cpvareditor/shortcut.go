package cpvareditor

import (
	"sdmm/internal/app/ui/shortcut"
	"sdmm/internal/platform"

	"github.com/hajimehoshi/ebiten/v2"
)

func (v *VarEditor) addShortcuts() {
	v.shortcuts.Add(shortcut.Shortcut{
		Name:         "cpvareditor#doToggleShowModified",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.Key1,
		SecondKeyAlt: ebiten.KeyKP1,
		Action:       v.doToggleShowModified,
	})
	v.shortcuts.Add(shortcut.Shortcut{
		Name:         "cpvareditor#doToggleShowByType",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.Key2,
		SecondKeyAlt: ebiten.KeyKP2,
		Action:       v.doToggleShowByType,
	})
	v.shortcuts.Add(shortcut.Shortcut{
		Name:         "cpvareditor#doToggleShowPins",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.Key3,
		SecondKeyAlt: ebiten.KeyKP3,
		Action:       v.doToggleShowPins,
	})
	v.shortcuts.Add(shortcut.Shortcut{
		Name:         "cpvareditor#doToggleShowTmp",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.Key4,
		SecondKeyAlt: ebiten.KeyKP4,
		Action:       v.doToggleShowTmp,
	})
}
