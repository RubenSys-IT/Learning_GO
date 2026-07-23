package main

import "fmt"

func main() {
	nombre := "Golang"

	// Imprime: "Golang" mide 6
	fmt.Printf("%q mide %v\n", nombre, len(nombre))

	nuevo := []byte(nombre)

	// Imprime: El slice creado: [71 111 108 97 110 103]
	fmt.Println("El slice creado:", nuevo)
}
