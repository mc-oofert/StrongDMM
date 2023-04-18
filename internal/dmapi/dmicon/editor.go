package dmicon

import (
	"sdmm/internal/platform/renderer/txcache"
	"sdmm/internal/rsc"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	spritePlaceholder *Sprite
)

func initEditorSprites() {
	atlas := rsc.EditorTextureAtlas()
	img := atlas.RGBA()

	dmi := &Dmi{
		IconWidth:     32,
		IconHeight:    32,
		TextureWidth:  atlas.Width,
		TextureHeight: atlas.Height,
		Cols:          1,
		Rows:          1,
		TextureID:     txcache.CreateTexture(ebiten.NewImageFromImage(img)),
	}

	spritePlaceholder = newDmiSprite(dmi, 0)
}

func SpritePlaceholder() *Sprite {
	if spritePlaceholder == nil {
		initEditorSprites()
	}
	return spritePlaceholder
}
