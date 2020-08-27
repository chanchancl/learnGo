package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qri-io/jsonschema"
)

func main() {
	buf := []byte(`{
    "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Person",
    "type": "object",
    "properties": {
        "firstName": {
            "type": "string"
        },
        "lastName": {
            "type": "string"
        },
        "age": {
            "description": "Age in years",
            "type": "integer",
            "minimum": 0
        },
        "friends": {
          "type" : "array",
          "items" : { "title" : "REFERENCE", "$ref" : "#" }
        }
    },
    "required": ["firstName", "lastName"]
  }`)
	schemaData := &jsonschema.Schema{}
	err := json.Unmarshal(buf, schemaData)
	if err != nil {
		fmt.Println(err.Error())
	}
	var valid = []byte(`{
		"foo": "a1.b"
	}`)

	for i := 0; i < 100; i++ {
		go func() {
			for {
				errs, err := schemaData.ValidateBytes(context.Background(), valid)

				if err != nil {
					fmt.Println(err.Error())
				}

				if len(errs) > 0 {
					for i := range errs {
						fmt.Println(errs[i])
					}
				}
			}
		}()
	}

	select {}
}
