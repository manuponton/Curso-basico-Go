package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}
	for _, valor := range slice {
		fmt.Println(valor)
	}
	isPalindromo("Amor a roma")
}
