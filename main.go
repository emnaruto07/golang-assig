package main

import (
	"fmt"

	"github.com/Shazeb01/golang-assig/pkg/config"
)

func main() {
	port := config.Get(port)
	fmt.Println(port)
}
