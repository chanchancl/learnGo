package basic

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
)

type Render interface {
	SetSampler(value Sampler)

	BeginRender()

	Write(file io.Writer)
}

type BasicRender struct {
	sampler Sampler
	Img     *image.RGBA
	config  *RenderConfig
}

func (c *BasicRender) SetSampler(sampler Sampler) {
	c.sampler = sampler
}

func (c *BasicRender) BeginRender() {
	if c.sampler == nil {
		log.Printf("Sampler is nil!")
		return
	}
	for y := 0; y < c.config.RenderHeight; y++ {
		for x := 0; x < c.config.RenderWidth; x++ {
			xx := float64(x) / float64(c.config.RenderWidth)
			yy := float64(y) / float64(c.config.RenderHeight)
			cl := color.Gray{uint8(Min(c.sampler.Sample(xx, yy)*255.0, 255.0))}
			//fmt.Printf("%v\n", cl)
			c.Img.Set(x, y, cl)
		}
	}
}

func (c *BasicRender) Write(file io.Writer) {
	if c.Img == nil {
		return
	}
	png.Encode(file, c.Img)
}

func NewRender(config *RenderConfig) BasicRender {
	render := BasicRender{
		Img:    image.NewRGBA(image.Rect(0, 0, config.RenderWidth, config.RenderHeight)),
		config: config,
	}

	return render
}
