package main

import (
	"fmt"
	light "learnGo/light2d/basic"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	N            = 256
	MAX_STEP     = 64
	MAX_DISTANCE = 3.0
	EPSILON      = 1e-6
	OUTPUT_NAME  = "output sence 2.png"
)

type MySampler struct{}

func circleSDF(x, y, cx, cy, r float64) float64 {
	ux, uy := x-cx, y-cy
	return math.Sqrt(ux*ux+uy*uy) - r
}

type SceneResult struct {
	sd, emissive float64
}

func unionOp(a, b SceneResult) SceneResult {
	if a.sd < b.sd {
		return a
	}
	return b
}

func intersectOp(a, b SceneResult) SceneResult {
	r := a
	if a.sd > b.sd {
		r = b
	}
	r.sd = math.Max(a.sd, b.sd)
	return r
}

func subtractOp(a, b SceneResult) SceneResult {
	r := a
	r.sd = math.Max(a.sd, -b.sd)
	if a.sd > -b.sd {
		r.emissive = a.emissive
	} else {
		r.emissive = b.emissive
	}
	return r
}

// 0.29/2+0.5
// 0.145+0.5
// 0.645

// 0.5+-0.29*sqrt(3)
func scene(x, y float64) SceneResult {
	r1 := SceneResult{circleSDF(x, y, 0.5, 0.21, 0.1), 2}
	r2 := SceneResult{circleSDF(x, y, 0.2489, 0.645, 0.1), 0.3}
	r3 := SceneResult{circleSDF(x, y, 0.7511, 0.645, 0.1), 0.3}

	// return unionOp(subtractOp(r1, r2), r3)
	return unionOp(unionOp(r1, r2), r3)
}

func trace(ox, oy, dx, dy float64) float64 {
	var t float64
	t = 0
	for i := 0; i < MAX_STEP && t < MAX_DISTANCE; i++ {
		r := scene(ox+dx*t, oy+dy*t)
		if r.sd < EPSILON {
			return r.emissive
		}
		t += r.sd
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
	file, err := os.Create(OUTPUT_NAME)
	if err != nil {
		fmt.Println("Error to create output file.")
		return
	}
	defer file.Close()

	H := 512
	W := 512

	start := time.Now()
	render := light.NewRender(
		&light.RenderConfig{
			RenderWidth:  W,
			RenderHeight: H,
		})

	render.SetSampler(MySampler{})
	render.BeginRender()

	fmt.Printf("%vx%v use time %vs\n", W, H, time.Since(start))

	render.Write(file)
}
