package menu

import (
	"sdmm/internal/app/ui/shortcut"
	"sdmm/internal/platform"

	"github.com/hajimehoshi/ebiten/v2"
)

func (m *Menu) addShortcuts() {
	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoNewWorkspace",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyN,
		Action:      m.app.DoNewWorkspace,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoOpen",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyO,
		Action:      m.app.DoOpen,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoSave",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyS,
		Action:      m.app.DoSave,
		IsEnabled:   m.app.HasActiveMap,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:         "menu#DoSaveAll",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.KeyShiftLeft,
		SecondKeyAlt: ebiten.KeyShiftRight,
		ThirdKey:     ebiten.KeyS,
		Action:       m.app.DoSaveAll,
		IsEnabled:    m.app.HasActiveMap,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoClose",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyW,
		Action:      m.app.DoClose,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:         "menu#DoCloseAll",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.KeyShiftLeft,
		SecondKeyAlt: ebiten.KeyShiftRight,
		ThirdKey:     ebiten.KeyW,
		Action:       m.app.DoCloseAll,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoExit",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyQ,
		Action:      m.app.DoExit,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoUndo",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyZ,
		Action:      m.app.DoUndo,
		IsEnabled:   m.app.CommandStorage().HasUndo,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:         "menu#DoRedo",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.KeyShiftLeft,
		SecondKeyAlt: ebiten.KeyShiftRight,
		ThirdKey:     ebiten.KeyZ,
		Action:       m.app.DoRedo,
		IsEnabled:    m.app.CommandStorage().HasRedo,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoRedo",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyY,
		Action:      m.app.DoRedo,
		IsEnabled:   m.app.CommandStorage().HasRedo,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoCopy",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyC,
		Action:      m.app.DoCopy,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoPaste",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyV,
		Action:      m.app.DoPaste,
		IsEnabled:   m.app.Clipboard().HasData,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoCut",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyX,
		Action:      m.app.DoCut,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:     "menu#DoDelete",
		FirstKey: ebiten.KeyDelete,
		Action:   m.app.DoDelete,
	})
	m.shortcuts.Add(shortcut.Shortcut{
		Name:        "menu#DoSearch",
		FirstKey:    platform.KeyModLeft(),
		FirstKeyAlt: platform.KeyModRight(),
		SecondKey:   ebiten.KeyF,
		Action:      m.app.DoSearch,
		IsEnabled:   m.app.HasActiveMap,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:         "menu#DoMultiZRendering",
		FirstKey:     platform.KeyModLeft(),
		FirstKeyAlt:  platform.KeyModRight(),
		SecondKey:    ebiten.Key0,
		SecondKeyAlt: ebiten.KeyKP0,
		Action:       m.app.DoMultiZRendering,
	})

	m.shortcuts.Add(shortcut.Shortcut{
		Name:     "menu#DoResetLayout",
		FirstKey: ebiten.KeyF5,
		Action:   m.app.DoResetLayout,
	})

	m.shortcuts.SetVisible(true)
}
