package main

import "fmt"

func main() {
	var keyBuf []byte
	kaut := "0001020304050607080910ff"
	fmt.Sscanf(kaut, "%x", &keyBuf)
	fmt.Printf("%#v\n", keyBuf)
}
