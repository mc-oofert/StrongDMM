package window

import (
	"image"
	"sdmm/internal/platform"
	"sdmm/internal/platform/renderer"
	"sdmm/internal/platform/renderer/txcache"
	"sdmm/internal/rsc"

	"github.com/SpaiR/imgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

var (
	// AppLogoTexture is a texture pointer to the application logo.
	AppLogoTexture uint32
)

type application interface {
	Process()
	PostProcess()
	CloseCheck()
	IsClosed() bool
	LayoutIniPath() string
}

var (
	pointSize float32 = 1
)

type Window struct {
	application application

	g *G

	mouseChangeCallbackId int
	mouseChangeCallbacks  map[int]func(uint, uint)
}

var wnd *Window

func New(application application) *Window {
	log.Print("creating native window")

	w := Window{application: application}
	wnd = &w
	w.mouseChangeCallbacks = make(map[int]func(uint, uint))

	mgr := renderer.New()

	fw, fh := ebiten.ScreenSizeInFullscreen()
	ebiten.SetWindowSize(fw, fh)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowIcon([]image.Image{rsc.EditorIcon().RGBA()})

	w.g = &G{
		mgr:        mgr,
		preUpdate:  w.beforeFrame,
		update:     w.runFrame,
		postUpdate: w.beforeFrame,
		retina:     platform.IsOsDarwin(),
	}

	log.Print("setting up Dear ImGui")
	w.setupImGui()

	txcache.SetTexture(imgui.TextureID(2), ebiten.NewImageFromImage(rsc.EditorIcon().RGBA()))
	AppLogoTexture = 2

	return &w
}

func (w *Window) Process() {
	ebiten.RunGame(w.g)
}

func (w *Window) mouseChangeCallback(x, y uint) {
	for _, cb := range w.mouseChangeCallbacks {
		cb(x, y)
	}
}

func (w *Window) Dispose() {
	w.disposeImGui()
}

func PointSize() float32 {
	return pointSize
}

func PointSizePtr() *float32 {
	return &pointSize
}

func SetPointSize(ps float32) {
	pointSize = ps
	configureFonts()
	wnd.g.mgr.UpdateFonts()
}

func SetFps(value int) {
	log.Print("set fps:", value)
	ebiten.SetTPS(value)
}

func (w *Window) setupImGui() {
	imgui.CreateContext(nil)

	io := imgui.CurrentIO()
	io.SetIniFilename(w.application.LayoutIniPath())
	io.SetConfigFlags(imgui.ConfigFlagsDockingEnable)

	w.setDefaultTheme()
}

func (*Window) disposeImGui() {
	if c, err := imgui.CurrentContext(); err == nil {
		c.Destroy()
	}
}

func (w *Window) beforeFrame() {
	runLaterJobs()
	runRepeatJobs()
}

func (w *Window) runFrame() {
	w.application.Process()
}

func (w *Window) afterFrame() {
	w.application.PostProcess()
}

func runLaterJobs() {
	for _, job := range laterJobs {
		job()
	}
	laterJobs = nil
}

func runRepeatJobs() {
	for _, job := range repeatJobs {
		job()
	}
}
