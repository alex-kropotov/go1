package main

import (
	"fmt"
	"go1/config"
)

func main() {
	c := *config.GetConfig()
	fmt.Println(c)
}