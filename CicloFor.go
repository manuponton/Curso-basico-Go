package main

import "fmt"

func main() {
	//For condicional
	const limite = 10
	for i := 0; i <= limite; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
