package tilemenu

import (
	"sdmm/internal/app/ui/shortcut"

	"github.com/hajimehoshi/ebiten/v2"
)

func (t *TileMenu) addShortcuts() {
	t.shortcuts.Add(shortcut.Shortcut{
		Name:      "tileMenu#close",
		FirstKey:  ebiten.KeyEscape,
		Action:    t.close,
		IsEnabled: func() bool { return t.opened },
	})
}
