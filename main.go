package main

import (
	"config"
	"fmt"
)

func main() {
	port := config.Get(port)
	fmt.Println(port)
}
