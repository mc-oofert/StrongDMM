package window

import (
	"sdmm/internal/platform/renderer"

	"github.com/hajimehoshi/ebiten/v2"
)

type G struct {
	mgr        *renderer.Manager
	preUpdate  func()
	update     func()
	postUpdate func()
	retina     bool
	w, h       int
}

func (g *G) Draw(screen *ebiten.Image) {
	g.mgr.Draw(screen)
}

func (g *G) Update() error {
	g.mgr.Update(float32(1.0 / ebiten.ActualTPS()))

	g.mgr.BeginFrame()
	g.update()
	g.mgr.EndFrame()

	return nil
}

func (g *G) Layout(outsideWidth, outsideHeight int) (int, int) {
	if g.retina {
		m := ebiten.DeviceScaleFactor()
		g.w = int(float64(outsideWidth) * m)
		g.h = int(float64(outsideHeight) * m)
	} else {
		g.w = outsideWidth
		g.h = outsideHeight
	}
	g.mgr.SetDisplaySize(float32(g.w), float32(g.h))
	return g.w, g.h
}
