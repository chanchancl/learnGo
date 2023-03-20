package main

import (
	"fmt"
	"reflect"

	"go.uber.org/fx"
)

type T1 int
type T2 int
type T3 int

func T1Construct() *T1 {
	fmt.Println("T1 is construct")
	return nil
}

func InvokeNeedT1(t *T1, tt *T2) {
	fmt.Println("Invoke need T1 is call")
}

func T2Construct(str string, t *T1) *T2 {
	fmt.Println("T2 is construct")
	return nil
}

func T3NeedString(str string, s int) *T3 {
	return nil
}

func main() {
	opts := []fx.Option{}
	opts = append(opts,
		fx.Invoke(InvokeNeedT1),
		fx.Provide(T1Construct),
		fx.Provide(func() string {
			return ""
		}),
		fx.Provide(T2Construct),
		fx.Invoke(T3NeedString),
	)

	// fx.Provide
	// fx.Invoke
	//

	app := fx.New(opts...)

	app.Run()

	Provide(T2Construct)
}

func Provide(fn interface{}) {
	tp := reflect.TypeOf(fn)

	fmt.Println(tp, tp.Kind())

	fmt.Println("Input:")
	for i := 0; i < tp.NumIn(); i++ {
		fmt.Println("\t" + tp.In(i).String())
	}

	fmt.Println("Output:")
	for i := 0; i < tp.NumOut(); i++ {
		fmt.Println("\t" + tp.Out(i).String())
	}
}
