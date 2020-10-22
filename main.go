package main

import (
	"encoding/json"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/micro/v2/cmd"
)

func main() {
	errors.SetErrorFormatter(func(e *errors.Error) string {
		data := make(map[string]interface{})
		data["status"] = e.Code
		data["message"] = e.Id + ", " + e.Detail
		data["data"] = nil
		j, _ := json.Marshal(data)
		return string(j)
	})
	cmd.Init()
}
