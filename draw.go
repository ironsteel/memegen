package main

import (
	"image"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// Width returns width of a string when drawn using the font.
func width(f *truetype.Font, s string, size float64) float64 {
	c := freetype.NewContext()
	scale := size / float64(f.FUnitsPerEm())

	width := 0
	prev, hasPrev := truetype.Index(0), false

	for _, rune := range s {
		index := f.Index(rune)
		if hasPrev {
			width += int(f.Kern(c.PointToFixed(float64(f.FUnitsPerEm()>>6)), prev, index))
		}

		width += int(f.HMetric(c.PointToFixed(float64(f.FUnitsPerEm()>>6)), index).AdvanceWidth)
		prev, hasPrev = index, true
	}
	return float64(width) * scale
}

func drawText(f *truetype.Font, topText, botText string, meme image.Image, dpi, size float64) (*image.RGBA, error) {

	rgba := image.NewRGBA(meme.Bounds())
	draw.Draw(rgba, rgba.Bounds(), meme, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.White)
	center := meme.Bounds().Size().X / 2

	w := int(width(f, topText, size) / 2)
	pt := freetype.Pt(center-w, 10+int(c.PointToFixed(size)>>6))
	_, err := c.DrawString(topText, pt)
	if err != nil {
		return nil, err
	}

	w = int(width(f, botText, size) / 2)
	pt = freetype.Pt(center-w, meme.Bounds().Size().Y-int(20))
	_, err = c.DrawString(botText, pt)
	if err != nil {
		return nil, err
	}

	return rgba, nil
}
