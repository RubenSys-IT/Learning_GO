//Este es el estado de un servidor

package main

import "fmt"

func main() {
	estado := "UP"
	switch estado {
	case "UP":
		fmt.Println("Servidor operativo")
	case "DOWN":
		fmt.Println("Servidor Caido")
	case "MAINTENANCE":
		fmt.Println("Servidor en mantenimiento")
	default:
		fmt.Println("Status desconocido.")
	}
}
