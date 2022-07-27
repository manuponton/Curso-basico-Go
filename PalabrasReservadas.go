package main

import "fmt"

func main() {
	/* Defer busca que se ejecute una funcion al final de la ejecucion de la principal
	 */
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es 2")
			/*Salta la ejecución y la continua*/
			continue
		}
		if i == 8 {
			fmt.Println("Break")
			/*Salta la ejecución y detiene el ciclo*/
			break
		}
	}
}
