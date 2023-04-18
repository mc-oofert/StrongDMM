package renderer

import (
	"runtime"
	"sdmm/internal/platform/renderer/txcache"

	"github.com/SpaiR/imgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

type GetCursorFn func() (x, y float32)

type Manager struct {
	Filter             ebiten.Filter
	GetCursor          GetCursorFn
	SyncInputsFn       func()
	SyncCursor         bool
	SyncInputs         bool
	ControlCursorShape bool
	ClipMask           bool

	ctx     *imgui.Context
	cliptxt string

	lmask *ebiten.Image

	width        float32
	height       float32
	screenWidth  int
	screenHeight int

	inputChars []rune
}

// Text implements imgui clipboard
func (m *Manager) Text() (string, error) {
	return m.cliptxt, nil
}

// SetText implements imgui clipboard
func (m *Manager) SetText(text string) {
	m.cliptxt = text
}

func (m *Manager) SetDisplaySize(width, height float32) {
	m.width = width
	m.height = height
}

func (m *Manager) Update(delta float32) {
	io := imgui.CurrentIO()

	if m.width > 0 || m.height > 0 {
		io.SetDisplaySize(imgui.Vec2{X: m.width, Y: m.height})
	} else if m.screenWidth > 0 || m.screenHeight > 0 {
		io.SetDisplaySize(imgui.Vec2{X: float32(m.screenWidth), Y: float32(m.screenHeight)})
	}

	io.SetDeltaTime(delta)

	if m.SyncCursor {
		if m.GetCursor != nil {
			x, y := m.GetCursor()
			io.SetMousePosition(imgui.Vec2{X: x, Y: y})
		} else {
			mx, my := ebiten.CursorPosition()
			io.SetMousePosition(imgui.Vec2{X: float32(mx), Y: float32(my)})
		}
		io.SetMouseButtonDown(0, ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft))
		io.SetMouseButtonDown(1, ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight))
		io.SetMouseButtonDown(2, ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle))
		xOff, yOff := ebiten.Wheel()
		io.AddMouseWheelDelta(float32(xOff), float32(yOff))
		m.controlCursorShape()
	}

	if m.SyncInputs {
		if m.SyncInputsFn != nil {
			m.SyncInputsFn()
		} else {
			m.inputChars = sendInput(&io, m.inputChars)
		}
	}
}

func (m *Manager) BeginFrame() {
	imgui.NewFrame()
}

func (m *Manager) EndFrame() {
	imgui.EndFrame()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.screenWidth, m.screenHeight = screen.Bounds().Dx(), screen.Bounds().Dy()

	imgui.Render()
	Render(screen, imgui.RenderedDrawData(), m.Filter)
}

func New() *Manager {
	imctx := imgui.CreateContext(nil)

	m := &Manager{
		ctx:                imctx,
		SyncCursor:         true,
		SyncInputs:         true,
		ClipMask:           true,
		ControlCursorShape: true,
		inputChars:         make([]rune, 0, 256),
	}

	runtime.SetFinalizer(m, (*Manager).onFinalize)

	m.UpdateFonts()
	m.setKeyMapping()

	return m
}

func (m *Manager) UpdateFonts() {
	io := imgui.CurrentIO()
	_ = io.Fonts().TextureDataRGBA32()
	io.Fonts().SetTextureID(1)
	txcache.SetFontAtlasTextureID(1)
}

func (m *Manager) controlCursorShape() {
	if !m.ControlCursorShape {
		return
	}
	switch imgui.MouseCursor() {
	case imgui.MouseCursorNone:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	case imgui.MouseCursorArrow:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	case imgui.MouseCursorTextInput:
		ebiten.SetCursorShape(ebiten.CursorShapeText)
	case imgui.MouseCursorResizeAll:
		ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	case imgui.MouseCursorResizeEW:
		ebiten.SetCursorShape(ebiten.CursorShapeEWResize)
	case imgui.MouseCursorResizeNS:
		ebiten.SetCursorShape(ebiten.CursorShapeNSResize)
	case imgui.MouseCursorHand:
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	default:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	}
}

func (m *Manager) onFinalize() {
	runtime.SetFinalizer(m, nil)
	m.ctx.Destroy()
}

func (m *Manager) setKeyMapping() {
	// Keyboard mapping. ImGui will use those indices to peek into the io.KeysDown[] array.
	io := imgui.CurrentIO()
	for imguiKey, nativeKey := range keys {
		io.KeyMap(imguiKey, nativeKey)
	}
}
