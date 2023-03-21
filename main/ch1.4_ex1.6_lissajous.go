package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete = []color.Color{
	color.Black,
	color.RGBA{0xf2, 0x00, 0x3c, 0xff},
	color.RGBA{0xfd, 0x6a, 0x02, 0xff},
	color.RGBA{0x4c, 0xbb, 0x17, 0xff},
	color.RGBA{0x04, 0x7c, 0xae, 0xff},
	color.RGBA{0x31, 0x14, 0x65, 0xff},
}

const (
	minColorIdx = 1
	maxColorIdx = 5
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
		colors  = 10
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		imageSize := 2*size + 1
		rect := image.Rect(0, 0, imageSize, imageSize)
		img := image.NewPaletted(rect, pallete)
		colorSectorSize := float64(imageSize) / float64(colors)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			posY := size + int(y*size+0.5)
			colorSectorId := uint8(math.Ceil(float64(posY) / colorSectorSize))
			if colorSectorId > maxColorIdx {
				colorSectorId = uint8(maxColorIdx) - (colorSectorId - uint8(maxColorIdx))
			}
			if colorSectorId < uint8(minColorIdx) {
				colorSectorId = uint8(minColorIdx)
			}

			img.SetColorIndex(size+int(x*size+0.5), posY, colorSectorId)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
