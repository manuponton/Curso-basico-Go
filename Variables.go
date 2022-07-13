package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaracion de variables enteras

	/*
		Los dos puntos indican que esta creando la variable en ese punto
		Si no se le coloca se asume que ha sido creada anteriormente y se le quiere cambiar el valor
	*/
	base := 12          // Se crea y se asigna el valor sin especificar el tipo de dato
	var altura int = 14 // Se asigna tipo de dato y valor
	var area int        // Se crea y no se le asigna un valor

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)

	//Zero values o valores por defecto
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//Calcular area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("area cuadrado: ", areaCuadrado)

}
