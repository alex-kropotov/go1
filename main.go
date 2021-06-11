package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, square, rad float64
	var threeDigit int
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scan(&x)
	fmt.Print("В теперь ширину: ")
	fmt.Scan(&y)
	fmt.Println("Площадь такого прямоугольника ", x * y)
	fmt.Println("************************************")
	fmt.Print("Введите площадь круга: ")
	fmt.Scan(&square)
	rad = math.Sqrt(square/math.Pi)
	fmt.Println("Диаметр такой окружности равен ", rad * 2,
		" a а длина окружности равна ", rad * math.Pi * 2)
	fmt.Println("************************************")
	fmt.Print("Ну и наконец введите трехзначное целое положительное число: ")
	fmt.Scan(&threeDigit)
	if (threeDigit > 99) && (threeDigit < 1000) {

		fmt.Println("Сотен ", threeDigit/100,
			", десятков ", (threeDigit % 100) / 10,
			", единиц", threeDigit % 10,
			)
	} else {
		fmt.Println("Вы ввели что-то не то")
	}
}