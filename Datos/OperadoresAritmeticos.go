package main

import "fmt"

func main() {
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("areaCuadrado", areaCuadrado)
	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	//Division residuo
	result = y % x
	fmt.Println("Division:", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decremental", x)
}
