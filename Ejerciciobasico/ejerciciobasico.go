// Ejercicio de varias cosas simples de Go
package main

import "fmt"

func main() {
	nombre := "Federico"
	apellidos := "Gonzalez"
	edad := int(24)
	altura := 1.982

	//fmt.Printf("El nombre es %s", nombre, "La edad es %f", edad, "La altura es %.1\n", altura)
	fmt.Printf("El nombre es %s\n", nombre)
	fmt.Printf("El apellido es %s\n", apellidos)
	fmt.Printf("La edad es %d\n", edad)
	fmt.Printf("La altura es %.2f\n", altura)

	nombre = "Hector"
	fmt.Printf("El nombre es %s\n", nombre)

	nuevo := nombre + " " + apellidos
	fmt.Printf("El nombre completo %s\n", nuevo)
}
