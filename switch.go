package main

import "fmt"

func main() {

	switch modulo := 13 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	}
}
