package txcache

import (
	"github.com/SpaiR/imgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

type textureCache struct {
	// fontAtlasID used by ImGui to store a texture with fonts.
	fontAtlasID imgui.TextureID
	// idCount used to create a unique ID for every texture.
	idCount        uint
	fontAtlasImage *ebiten.Image
	cache          map[imgui.TextureID]*ebiten.Image
}

var txCache = &textureCache{
	fontAtlasID:    1,
	idCount:        999, // first 999 is for a technical usage
	cache:          make(map[imgui.TextureID]*ebiten.Image),
	fontAtlasImage: nil,
}

func (tx *textureCache) getFontAtlas() *ebiten.Image {
	if tx.fontAtlasImage == nil {
		tx.fontAtlasImage = getTexture(imgui.CurrentIO().Fonts().TextureDataRGBA32())
	}
	return tx.fontAtlasImage
}

func SetFontAtlasTextureID(id imgui.TextureID) {
	if txCache.fontAtlasImage != nil {
		txCache.fontAtlasImage.Dispose()
		txCache.fontAtlasImage = nil
	}
	txCache.fontAtlasID = id
}

func CreateTexture(img *ebiten.Image) imgui.TextureID {
	txCache.idCount++
	id := imgui.TextureID(txCache.idCount)
	SetTexture(id, img)
	return id
}

func GetTexture(id imgui.TextureID) *ebiten.Image {
	if id != txCache.fontAtlasID {
		if im, ok := txCache.cache[id]; ok {
			return im
		}
	}
	return txCache.getFontAtlas()
}

func SetTexture(id imgui.TextureID, img *ebiten.Image) {
	txCache.cache[id] = img
}

func RemoveTexture(id imgui.TextureID) {
	if tx, ok := txCache.cache[id]; ok {
		tx.Dispose()
		delete(txCache.cache, id)
	}
}

func getTexture(tex *imgui.RGBA32Image) *ebiten.Image {
	n := tex.Width * tex.Height
	srcPix := (*[1 << 28]uint8)(tex.Pixels)[: n*4 : n*4]
	pix := make([]uint8, n*4)
	// Note: Ebiten expects colors in premultiplied-alpha form.
	// However, the imgui library exports pixmaps in straight-alpha form.
	// Also, not doing this modification in-place,
	// as srcPix points right into an imgui-owned data structure.
	for i := 0; i < n; i++ {
		alpha := uint16(srcPix[4*i+3])
		pix[4*i] = uint8((uint16(srcPix[4*i])*alpha + 127) / 255)
		pix[4*i+1] = uint8((uint16(srcPix[4*i+1])*alpha + 127) / 255)
		pix[4*i+2] = uint8((uint16(srcPix[4*i+2])*alpha + 127) / 255)
		pix[4*i+3] = uint8(alpha)
	}
	img := ebiten.NewImage(tex.Width, tex.Height)
	img.WritePixels(pix)
	return img
}
