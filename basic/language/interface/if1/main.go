package main

type infA interface {
	READ()
}

type structA struct {
	infA // is a
}

type structB struct {
	B infA // has-a
}

//func (b b) READ() {}

func main() {
	a := structA{}
	b := structB{}

	// Both of them will panic
	a.infA.READ()
	b.B.READ()
}
