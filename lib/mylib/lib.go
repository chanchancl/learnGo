package mylib

var (
	// PI the number of pi
	PI float32 = 3.1415926
)

// ExportFunction compute the square of x
func ExportFunction(x int) int {
	return x * x
}

func unexportFunction(x int) int {
	return x * x * x
}
