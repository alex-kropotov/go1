package main

import (
	"fmt"
	"github.com/alex-kropotov/go1/fib"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fib.GetFibWithCache(i))
	}
}

