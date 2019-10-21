package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	var keybuf, msgbuf bytes.Buffer

	key := "CCL"
	message := "你好"
	keybuf.WriteString(key)
	msgbuf.WriteString(message)
	mac := hmac.New(sha256.New, keybuf.Bytes())

	mac.Write(msgbuf.Bytes())

	out := mac.Sum(nil)

	fmt.Printf("%x %v\n", out, len(out))

	//other

	empty := make([]byte, 16)
	for i := 0; i < 16; i++ {
		empty[i] = byte(i)
	}

	emptyStr := fmt.Sprintf("%x", empty)
	fmt.Printf("string byte : %v %d\n", emptyStr, len(emptyStr))
	fmt.Printf("[] byte     : %v %d\n", empty, len(empty))
}
