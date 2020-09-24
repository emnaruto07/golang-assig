package main

import (
	"fmt"

	"github.com/Shazeb01/golang-assig/pkg/config"
)

func main() {
	keys_array := [5]string{"port", "host", "connectionHost", "connectionDB", "connectionUser"}
	for _, keys := range keys_array {
		fmt.Println(keys + " " + config.Get(keys))
	}

	var set_value = map[string]string{
		"port":           "8001",
		"host":           "127.0.0.1",
		"connectionHost": "db.mysecretapp.prod",
		"connectionDB":   "testapp",
		"connectionUser": "user",
	}

	for keys, values := range set_value {
		config.Set(keys, values)
		fmt.Println(keys + " " + config.Get(keys))
	}
}
