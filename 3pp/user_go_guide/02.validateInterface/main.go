package main

type HttpHandler interface {
	OnRequest(method, body string)
}

type MyHandler struct {
}

// 约束 Handler 必须实现 HttpHandler 接口
var _ HttpHandler = (*MyHandler)(nil)

// TODO:
// 可以尝试注释这一行
func (c *MyHandler) OnRequest(method, body string) {}

func main() {

}
