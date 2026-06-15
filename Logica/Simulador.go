package main

import "fmt"

func main() {
	intentosFallidos := 0
	passCore := "Hol9034fd"
	var pass string

	// Cambiamos la condición para que funcione mientras tenga intentos disponibles
	for intentosFallidos < 3 {
		fmt.Println("Introduzca una clave: ")
		fmt.Scanln(&pass)

		if pass == passCore {
			fmt.Println("Contraseña correcta")
			return // Salimos del programa si es correcta
		} else {
			intentosFallidos++
			fmt.Printf("Contraseña incorrecta. Intento %d de 3\n", intentosFallidos)
		}
	}

	// Si sale del bucle, es porque llegó a 3 intentos
	fmt.Println("Sistema bloqueado: demasiados intentos fallidos.")
}
