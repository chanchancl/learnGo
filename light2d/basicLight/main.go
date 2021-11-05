package main

import (
	"fmt"
	"learnGo/light2d/basic"
	light "learnGo/light2d/basic"
	"math"
	"os"
	"time"
)

const (
	N            = 128
	MAX_STEP     = 64
	MAX_DISTANCE = 3.0
	EPSILON      = 1e-6
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

func main() {
	file, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error to create output file.")
		return
	}
	defer file.Close()

	H := 512
	W := 512

	start := time.Now()
	render := light.NewRender(basic.NewRenderConfig(W, H))

	render.SetSampler(light.NewRandomSampler(N, trace))
	render.BeginRender()

	fmt.Printf("%vx%v use time %vs\n", W, H, time.Since(start))

	// for x := 0; x < 512; x++ {
	// 	cl := color.Gray{uint8(float32(x) / float32(512) * 255)}
	// 	for y := 0; y < 100; y++ {
	// 		render.Img.Set(x, y, cl)
	// 	}
	// }

	render.Write(file)
}
