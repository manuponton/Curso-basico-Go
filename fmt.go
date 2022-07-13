package main

import "fmt"

func main() {
	//Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "word"
	//Println
	fmt.Println(helloMessage, worldMessage) //Println se añade un salto de linea
	fmt.Println(helloMessage, worldMessage)

	//PrintF
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos) // Es buena practica especificar el tipo de dato
	fmt.Printf("%v tiene más de %v cursos \n", nombre, cursos)

	//Sprintf
	//Se concatena texto con format y se asigna a una variable
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)
	//Rescatar el tipo de dato de una variable
	fmt.Printf("hellowMessage : %T \n", helloMessage)
	fmt.Printf("cursos : %T \n", cursos)
}
