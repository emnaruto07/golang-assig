package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	conf := config.set("port", port)
	fmt.Println(conf)
}
