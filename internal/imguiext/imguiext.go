package imguiext

import (
	"sdmm/internal/platform"

	"github.com/SpaiR/imgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

func SetItemHoveredTooltip(text string) {
	if imgui.IsItemHovered() {
		imgui.SetTooltip(text)
	}
}

func InputIntClamp(
	label string,
	v *int32,
	min, max, step, stepFast int,
) bool {
	if imgui.InputIntV(label, v, step, stepFast, imgui.InputTextFlagsNone) {
		if int(*v) > max {
			*v = int32(max)
		} else if int(*v) < min {
			*v = int32(min)
		}
		return true
	}
	return false
}

func IsAltDown() bool {
	return imgui.IsKeyDown(int(ebiten.KeyAltLeft)) || imgui.IsKeyDown(int(ebiten.KeyAltRight))
}

func IsShiftDown() bool {
	return imgui.IsKeyDown(int(ebiten.KeyShiftLeft)) || imgui.IsKeyDown(int(ebiten.KeyShiftRight))
}

func IsCtrlDown() bool {
	return imgui.IsKeyDown(int(ebiten.KeyControlLeft)) || imgui.IsKeyDown(int(ebiten.KeyControlRight))
}

func IsModDown() bool {
	return imgui.IsKeyDown(int(platform.KeyModLeft())) || imgui.IsKeyDown(int(platform.KeyModRight()))
}
