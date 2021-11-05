package basic

import (
	"math"
	"math/rand"
)

type Sampler interface {
	Sample(x, y float64) float64
}

type TraceFunc func(x, y, dx, dy float64) float64

type basicSampler struct {
	sampleCount int
	trace       TraceFunc
}

func NewRandomSampler(sampleCount int, trace TraceFunc) Sampler {
	return &basicSampler{
		sampleCount: sampleCount,
		trace:       trace,
	}
}

func (c *basicSampler) Sample(x, y float64) float64 {
	var sum float64
	N := c.sampleCount
	for i := 0; i < N; i++ {
		// 随机采样 N 次， [0, 2pi]
		// a := 2.0 * math.Pi * rand.Float64()
		// 均匀采样 N 次，将[0, 2pi] 分为 N 等份
		// a := 2 * math.Pi * float64(i) / N

		// 均匀加随机， 分成等份的同时，在前后有[0,1]弧度的浮动
		a := 2.0 * math.Pi * (float64(i) + rand.Float64()) / float64(N)
		sum += c.trace(x, y, math.Cos(float64(a)), math.Sin(float64(a)))
	}
	return sum / float64(N)
}
