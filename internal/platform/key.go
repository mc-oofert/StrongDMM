package platform

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func KeyModName() string {
	if IsOsDarwin() {
		return "Cmd"
	}
	return "Ctrl"
}

func KeyModLeft() ebiten.Key {
	if IsOsDarwin() {
		return ebiten.KeyMetaLeft
	}
	return ebiten.KeyControlLeft
}

func KeyModRight() ebiten.Key {
	if IsOsDarwin() {
		return ebiten.KeyMetaRight
	}
	return ebiten.KeyControlRight
}
