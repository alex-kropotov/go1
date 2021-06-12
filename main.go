package main

import (
	"fmt"
	"os"
)
func main() {
	var action int8
	fmt.Print("Что будем делать? 1 - калькулятор, 2 - искать простые числа: ")
	fmt.Scanln(&action)
	switch action {
	case 1:
		calc()
	case 2:
		simple()
	}
}

func calc() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя")
		} else {
			res = a / b
		}
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
}

func simple() {
	var limitValue int
	var result []int
	fmt.Print("Введите положительное целое число (больше 1 и меньше 1000): ")
	fmt.Scanln(&limitValue)
	if (limitValue > 1) && (limitValue < 1000) {
		for i := 2; i <= limitValue; i++ {
			if i < 4 {
				result = append(result, i)
			} else {
				ok := true
				for _, val := range result {
					ok = ok && !(i % val == 0)
				}
				if ok {
					result = append(result, i)
				}
			}
		}
	} else {
		fmt.Println("Не очень удачное число вы ввели")
	}
	first := true
	for _, val := range result {
		if !first {
			fmt.Print(", ", val)
		} else {
			first = false
			fmt.Print(val)
		}
	}
}

