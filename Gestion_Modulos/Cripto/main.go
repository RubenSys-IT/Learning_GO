package main

import (
	"cifrador/cifrador"
	"fmt"
)

func main() {
	logPayload := "Select * from usuarios"

	// 1. Imprimimos el texto antes de modificarlo (usando fmt.Printf)
	fmt.Printf("El texto original: %s\n", logPayload)

	// 2. Pasamos la referencia (&) Y la clave de cifrado (ej. 3)
	cifrador.OfuscarPayload(&logPayload, 3)

	// 3. Imprimimos el texto ya ofuscado
	fmt.Printf("El texto ofuscado: %s\n", logPayload)
}
