package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//Con AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	//Con OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Si imprime")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("57")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
