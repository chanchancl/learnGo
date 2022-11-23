package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input []byte
	}{
		{
			[]byte(`{"code":"-1", "msg":"failed"}`),
		},
		{
			[]byte(`{"code":"0", "msg":"success", "data":{
				"total": 166,
				"pageNo": 1,
				"pageSize": 1000,
				"list": [
					{
						"indexCode": "root00000",
						"nodeType": 5,
						"name": "xxx"
					}
				]
			}}`),
		},
	}

	for _, tcase := range tests {
		t.Run("test", func(t *testing.T) {
			fmt.Println(*ParseJsonObject(tcase.input))
		})
	}
}
