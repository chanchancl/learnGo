package main

import (
	"fmt"
	light "learnGo/light2d/basic"
	"math"
	"math/rand"
	"os"
)

const (
	N            = 64
	MAX_STEP     = 10
	MAX_DISTANCE = 3.0
	EPSILON      = 1e-7
)

type MySampler struct{}

func circleSDF(x, y, cx, cy, r float64) float64 {
	ux, uy := x-cx, y-cy
	return math.Sqrt(ux*ux+uy*uy) - r
}

func trace(ox, oy, dx, dy float64) float64 {
	var t float64
	t = 0
	for i := 0; i < MAX_STEP && t < MAX_DISTANCE; i++ {
		sd := circleSDF(ox+dx*t, oy+dy*t, 0.5, 0.5, 0.1)
		if sd < EPSILON {
			return 2.0
		}
		t += sd
	}
	return 0.0
}

func (this MySampler) Sample(x, y float64) float64 {
	var sum float64
	for i := 0; i < N; i++ {
		// 随机采样 N 次， [0, 2pi]
		// a := 2.0 * math.Pi * rand.Float64()
		// 均匀采样 N 次，将[0, 2pi] 分为 N 等份
		// a := 2 * math.Pi * float64(i) / N

		// 均匀加随机， 分成等份的同时，在前后有[0,1]弧度的浮动
		a := 2.0 * math.Pi * (float64(i) + rand.Float64()) / N
		sum += trace(x, y, math.Cos(float64(a)), math.Sin(float64(a)))
	}
	return sum / N
}

func main() {
	file, err := os.Create("output av rand 233.png")
	if err != nil {
		fmt.Println("Error to create output file.")
		return
	}
	defer file.Close()

	render := light.NewRender(light.RenderConfig{512, 512})

	render.SetSampler(MySampler{})
	render.BeginRender()

	// for x := 0; x < 512; x++ {
	// 	cl := color.Gray{uint8(float32(x) / float32(512) * 255)}
	// 	for y := 0; y < 100; y++ {
	// 		render.Img.Set(x, y, cl)
	// 	}
	// }

	render.Write(file)
}
