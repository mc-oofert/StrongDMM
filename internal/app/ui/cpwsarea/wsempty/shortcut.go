package wsempty

import (
	"sdmm/internal/app/ui/shortcut"
	"sdmm/internal/platform"

	"github.com/hajimehoshi/ebiten/v2"
)

func (ws *WsEmpty) addShortcuts() {
	ws.shortcuts.Add(shortcut.Shortcut{
		Name:        "wsempty#loadSelectedMaps",
		FirstKey:    ebiten.KeyEnter,
		FirstKeyAlt: ebiten.KeyKPEnter,
		Action:      ws.loadSelectedMaps,
	})

	ws.shortcuts.Add(shortcut.Shortcut{
		Name:     "wsempty#dropSelectedMaps",
		FirstKey: ebiten.KeyEscape,
		Action:   ws.dropSelectedMaps,
	})
	ws.shortcuts.Add(shortcut.Shortcut{
		Name:        "wsempty#dropSelectedMaps",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyD,
		Action:      ws.dropSelectedMaps,
	})

	ws.shortcuts.Add(shortcut.Shortcut{
		Name:        "wsempty#selectAllMaps",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyA,
		Action:      ws.selectAllMaps,
	})
}
