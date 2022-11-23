package main

import (
	"encoding/json"
	"fmt"
)

type TTableData struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Data    TData  `json:"data,omitempty"`
}

type TData struct {
	Total    int         `json:"total,omitempty"`
	PageNo   int         `json:"pageNo,omitempty"`
	PageSize int         `json:"pageSize,omitempty"`
	List     []TListItem `json:"list,omitempty"`
}

type TListItem struct {
	IndexCode string `json:"indexCode,omitempty"`
	NodeType  int    `json:"nodeType,omitempty"`
	Name      string `json:"name,omitempty"`
	// ...
}

func ParseJsonObject(input []byte) *TTableData {
	data := &TTableData{}
	if err := json.Unmarshal(input, data); err != nil {
		fmt.Printf("Error when unmarshaling json, %v", err.Error())
		return nil
	}
	return data
}

func main() {

}
