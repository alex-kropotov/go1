package main

import (
	"fmt"
	"go1/config"
)

func main() {
	c := config.GetConfigFromDotEnv()
	fmt.Println("Конфиг из .anv")
	fmt.Println(c)

	fmt.Println("")
	fmt.Println("Конфиг из yaml")
	c2 := config.GetConfigFromYaml()
	fmt.Println(c2)
}