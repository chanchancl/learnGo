package basic

type RenderConfig struct {
	RenderWidth  int
	RenderHeight int
}

func NewRenderConfig(W, H int) *RenderConfig {
	return &RenderConfig{
		RenderWidth:  W,
		RenderHeight: H,
	}
}
