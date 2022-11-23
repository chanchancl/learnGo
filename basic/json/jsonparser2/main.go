package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)

type FailedDetails struct {
	ErrorCode int
	ErrorMsg  string
}

type LoginData struct {
	Token string
	/// other fileds
}

type LoginResult struct {
	Success      bool
	Data         LoginData
	FailedReason FailedDetails
}

var succBytes = []byte(`
{
	"data": {
		"token": "testToken"
	}
}
`)

var errBytes = []byte(`
{
	"data": ""
}
`)

func GetToken(input []byte) (string, bool) {
	token, err := jsonparser.GetString(input, "data", "token")
	if err != nil {
		return "", false
	}
	return token, true
}

func Login(input []byte) {
	token, succ := GetToken(input)
	if !succ {
		fmt.Println("登陆失败")
		return
	}
	fmt.Printf("登录成功, 切token为 : %v\n", token)
}

func main() {
	Login(succBytes)
	Login(errBytes)
}
