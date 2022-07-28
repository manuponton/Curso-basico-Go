package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["jose"] = 14
	m["pepito"] = 20
	fmt.Println(m)
	//Forma diferente de declararlo
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temperature)

	//Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["jose"] //Se usa ok para validar si existe en el diccionario
	fmt.Println(value, ok)
}
