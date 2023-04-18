package renderer

import (
	"fmt"
	"image"
	"sdmm/internal/platform/renderer/native"
	"sdmm/internal/platform/renderer/txcache"
	"unsafe"

	"github.com/SpaiR/imgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

type cVec2x32 struct {
	X float32
	Y float32
}

type cImDrawVertx32 struct {
	Pos cVec2x32
	UV  cVec2x32
	Col uint32
}

type cVec2x64 struct {
	X float64
	Y float64
}

type cImDrawVertx64 struct {
	Pos cVec2x64
	UV  cVec2x64
	Col uint32
}

func getVertices(vBuf unsafe.Pointer, vbLen, vSize, offPos, offUV, offCol int) []ebiten.Vertex {
	if native.SzFloat() == 4 {
		return getVerticesX32(vBuf, vbLen, vSize, offPos, offUV, offCol)
	}
	if native.SzFloat() == 8 {
		return getVerticesX64(vBuf, vbLen, vSize, offPos, offUV, offCol)
	}
	panic("invalid char size")
}

func getVerticesX32(vBuf unsafe.Pointer, vbLen, vSize, offPos, offUV, offCol int) []ebiten.Vertex {
	if offPos != 0 || offUV != 8 || offCol != 16 {
		panic("TODO: invalid vertex layout")
	}

	n := vbLen / vSize
	vertices := make([]ebiten.Vertex, 0, vbLen/vSize)
	rawVertices := (*[1 << 28]cImDrawVertx32)(vBuf)[:n:n]

	for i := 0; i < n; i++ {
		vertices = append(vertices, ebiten.Vertex{
			SrcX:   rawVertices[i].UV.X,
			SrcY:   rawVertices[i].UV.Y,
			DstX:   rawVertices[i].Pos.X,
			DstY:   rawVertices[i].Pos.Y,
			ColorR: float32(rawVertices[i].Col&0xFF) / 255,
			ColorG: float32(rawVertices[i].Col>>8&0xFF) / 255,
			ColorB: float32(rawVertices[i].Col>>16&0xFF) / 255,
			ColorA: float32(rawVertices[i].Col>>24&0xFF) / 255,
		})
	}

	return vertices
}

func getVerticesX64(vBuf unsafe.Pointer, vblLen, vSize, offPos, offUV, offCol int) []ebiten.Vertex {
	if offPos != 0 || offUV != 8 || offCol != 16 {
		panic("TODO: invalid vertex layout (64)")
	}

	n := vblLen / vSize
	vertices := make([]ebiten.Vertex, 0, vblLen/vSize)
	rawVertices := (*[1 << 28]cImDrawVertx64)(vBuf)[:n:n]

	for i := 0; i < n; i++ {
		vertices = append(vertices, ebiten.Vertex{
			SrcX:   float32(rawVertices[i].UV.X),
			SrcY:   float32(rawVertices[i].UV.Y),
			DstX:   float32(rawVertices[i].Pos.X),
			DstY:   float32(rawVertices[i].Pos.Y),
			ColorR: float32(rawVertices[i].Col&0xFF) / 255,
			ColorG: float32(rawVertices[i].Col>>8&0xFF) / 255,
			ColorB: float32(rawVertices[i].Col>>16&0xFF) / 255,
			ColorA: float32(rawVertices[i].Col>>24&0xFF) / 255,
		})
	}

	return vertices
}

func getIndices(ibuf unsafe.Pointer, iblen, isize int) []uint16 {
	n := iblen / isize
	switch isize {
	case 2:
		// direct conversion (without a data copy)
		//TODO: document the size limit (?) this fits 268435456 bytes
		// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
		return (*[1 << 28]uint16)(ibuf)[:n:n]
	case 4:
		slc := make([]uint16, n)
		for i := 0; i < n; i++ {
			slc[i] = uint16(*(*uint32)(unsafe.Pointer(uintptr(ibuf) + uintptr(i*isize))))
		}
		return slc
	case 8:
		slc := make([]uint16, n)
		for i := 0; i < n; i++ {
			slc[i] = uint16(*(*uint64)(unsafe.Pointer(uintptr(ibuf) + uintptr(i*isize))))
		}
		return slc
	default:
		panic(fmt.Sprint("byte size", isize, "not supported"))
	}
	return nil
}

// Render the ImGui drawData into the target *ebiten.Image
func Render(target *ebiten.Image, drawData imgui.DrawData, dfilter ebiten.Filter) {
	render(target, nil, drawData, dfilter)
}

// RenderMasked renders the ImGui drawData into the target *ebiten.Image with ebiten.CompositeModeCopy for masking
func RenderMasked(target *ebiten.Image, mask *ebiten.Image, drawData imgui.DrawData, dfilter ebiten.Filter) {
	render(target, mask, drawData, dfilter)
}

func render(target *ebiten.Image, mask *ebiten.Image, drawData imgui.DrawData, dfilter ebiten.Filter) {
	if !drawData.Valid() {
		return
	}

	//targetw, targeth := target.Bounds().Dx(), target.Bounds().Dy()

	vertexSize, vertexOffsetPos, vertexOffsetUv, vertexOffsetCol := imgui.VertexBufferLayout()
	indexSize := imgui.IndexBufferLayout()

	opt := &ebiten.DrawTrianglesOptions{
		Filter: dfilter,
	}
	//var opt2 *ebiten.DrawImageOptions
	if mask != nil {
		//opt2 = &ebiten.DrawImageOptions{
		//	Blend: ebiten.BlendSourceOver,
		//}
	}

	for _, clist := range drawData.CommandLists() {
		vertexBuffer, vertexLen := clist.VertexBuffer()
		indexBuffer, indexLen := clist.IndexBuffer()
		vertices := getVertices(vertexBuffer, vertexLen, vertexSize, vertexOffsetPos, vertexOffsetUv, vertexOffsetCol)
		vbuf := vcopy(vertices)
		indices := getIndices(indexBuffer, indexLen, indexSize)

		var indexBufferOffset int

		for _, cmd := range clist.Commands() {
			eCount := cmd.ElementCount()

			if cmd.HasUserCallback() {
				cmd.CallUserCallback(clist)
			} else {
				//clipRect := cmd.ClipRect()
				texID := cmd.TextureID()
				tx := txcache.GetTexture(texID)

				vmultiply(vertices, vbuf, tx.Bounds().Min, tx.Bounds().Max)

				target.DrawTriangles(vbuf, indices[indexBufferOffset:indexBufferOffset+eCount], tx, opt)
			}
			indexBufferOffset += eCount
		}
	}
}

func lerp(a, b int, t float32) float32 {
	return float32(a)*(1-t) + float32(b)*t
}

func vcopy(v []ebiten.Vertex) []ebiten.Vertex {
	cl := make([]ebiten.Vertex, len(v))
	copy(cl, v)
	return cl
}

func vmultiply(v, vbuf []ebiten.Vertex, bmin, bmax image.Point) {
	for i := range vbuf {
		vbuf[i].SrcX = lerp(bmin.X, bmax.X, v[i].SrcX)
		vbuf[i].SrcY = lerp(bmin.Y, bmax.Y, v[i].SrcY)
	}
}
