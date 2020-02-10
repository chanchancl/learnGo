package main

import (
	"bytes"
	"fmt"
)

func main() {
	{
		src := bytes.NewBufferString("abcdefg")
		dst := bytes.NewBuffer(nil)

		// Growing dst's buffer as needed
		n, err := dst.ReadFrom(src)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(n)
		fmt.Println(dst.String())
	}
	{
		src := bytes.NewBufferString("abcdefg")
		src.ReadByte()
		err := src.UnreadByte()
		fmt.Println(err)
		err = src.UnreadByte()
		fmt.Println(err)
	}

}
