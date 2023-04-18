package renderer

import (
	"github.com/SpaiR/imgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var keys = map[int]int{
	imgui.KeyTab:        int(ebiten.KeyTab),
	imgui.KeyLeftArrow:  int(ebiten.KeyLeft),
	imgui.KeyRightArrow: int(ebiten.KeyRight),
	imgui.KeyUpArrow:    int(ebiten.KeyUp),
	imgui.KeyDownArrow:  int(ebiten.KeyDown),
	imgui.KeyPageUp:     int(ebiten.KeyPageUp),
	imgui.KeyPageDown:   int(ebiten.KeyPageDown),
	imgui.KeyHome:       int(ebiten.KeyHome),
	imgui.KeyEnd:        int(ebiten.KeyEnd),
	imgui.KeyInsert:     int(ebiten.KeyInsert),
	imgui.KeyDelete:     int(ebiten.KeyDelete),
	imgui.KeyBackspace:  int(ebiten.KeyBackspace),
	imgui.KeySpace:      int(ebiten.KeySpace),
	imgui.KeyEnter:      int(ebiten.KeyEnter),
	imgui.KeyEscape:     int(ebiten.KeyEscape),
	imgui.KeyA:          int(ebiten.KeyA),
	imgui.KeyC:          int(ebiten.KeyC),
	imgui.KeyV:          int(ebiten.KeyV),
	imgui.KeyX:          int(ebiten.KeyX),
	imgui.KeyY:          int(ebiten.KeyY),
	imgui.KeyZ:          int(ebiten.KeyZ),
}

func sendInput(io *imgui.IO, inputChars []rune) []rune {
	io.KeyAlt(bool2int(ebiten.IsKeyPressed(ebiten.KeyAlt)), 0)
	io.KeyShift(bool2int(ebiten.IsKeyPressed(ebiten.KeyShift)), 0)
	io.KeyCtrl(bool2int(ebiten.IsKeyPressed(ebiten.KeyControl)), 0)
	io.KeySuper(bool2int(ebiten.IsKeyPressed(ebiten.KeyMeta)), 0)

	inputChars = ebiten.AppendInputChars(inputChars)

	if len(inputChars) > 0 {
		io.AddInputCharacters(string(inputChars))
		inputChars = inputChars[:0]
	}

	for _, iv := range keys {
		if inpututil.IsKeyJustPressed(ebiten.Key(iv)) {
			io.KeyPress(iv)
		}
		if inpututil.IsKeyJustReleased(ebiten.Key(iv)) {
			io.KeyRelease(iv)
		}
	}

	return inputChars
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
