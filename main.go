package main

import (
	"fmt"
	"os"

	"github.com/Shazeb01/golang-assig/pkg/config"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	conf := config.Get(port)
	fmt.Println(conf)
}
