package main

import "fmt"

func hellowWord(message string) {
	fmt.Println(message)
}
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
func main() {
	hellowWord("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value 1: ", value1, " Value 2: ", value2)

	value3, _ := doubleReturn(2) //Descartar los valores si no los necesitas
	fmt.Println("Value 3: ", value3)

}
