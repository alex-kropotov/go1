package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите небольшое целое число: ")
	fmt.Scanln(&number)
	//
	// простой код
	//
	fmt.Println()
	fmt.Println("Простой код")
	for i, first, second := 1, 0, 1; i <= number; i++ {
		first, second = second, first + second
		fmt.Print(first, " ")
	}
	fmt.Println()
	//
	// с замыканием
	//
	fmt.Println()
	fmt.Println("С замыканием")
	f := fib()
	for i := 1; i <= number; i++ {
		f()
	}
	fmt.Println()
}

func fib() func() {
	first, second := 0, 1
	return func() {
		first, second = second, first + second
		fmt.Print(first, " ")
	}
}