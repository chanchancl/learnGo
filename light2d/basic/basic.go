package basic

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
)

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type Sampler interface {
	Sample(x, y float64) float64
}

type basicSampler struct{}

func (this *basicSampler) sample(x, y float64) float64 {
	return 0
}

type RenderConfig struct {
	RenderWidth  int
	RenderHeight int
}

type Render interface {
	SetSampler(value Sampler)

	BeginRender()

	Write(file io.Writer)
}

type BasicRender struct {
	sampler Sampler
	Img     *image.RGBA
	config  RenderConfig
}

func (this *BasicRender) SetSampler(sampler Sampler) {
	this.sampler = sampler
}

func (this *BasicRender) BeginRender() {
	if this.sampler == nil {
		log.Printf("Sampler is nil!")
		return
	}
	for y := 0; y < this.config.RenderHeight; y++ {
		for x := 0; x < this.config.RenderWidth; x++ {
			xx := float64(x) / float64(this.config.RenderWidth)
			yy := float64(y) / float64(this.config.RenderHeight)
			cl := color.Gray{uint8(min(this.sampler.Sample(xx, yy)*255.0, 255.0))}
			//fmt.Printf("%v\n", cl)
			this.Img.Set(x, y, cl)
		}
	}
}

func (this *BasicRender) Write(file io.Writer) {
	if this.Img == nil {
		return
	}
	png.Encode(file, this.Img)
}

func NewRender(config RenderConfig) BasicRender {
	render := BasicRender{
		Img:    image.NewRGBA(image.Rect(0, 0, config.RenderWidth, config.RenderHeight)),
		config: config,
	}

	return render
}
