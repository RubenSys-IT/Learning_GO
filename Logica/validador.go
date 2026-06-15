package main

import (
	"fmt"
	"strings"
)

//Creamos funcion personalizada

func contra(pass string) {
	caracteres := []string{"!", "@", "#", "$"}

	// Recorremos cada elemento del slice
	for _, char := range caracteres {
		if strings.Contains(pass, char) {
			fmt.Println("Contiene uno de los caracteres")
			return // Salimos de la función al encontrar el primero
		}
	}
}

func main() {
	contra("@34gd456")
}
