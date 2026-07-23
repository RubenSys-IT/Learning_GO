package main

import "fmt"

func main() {
	saludo := "hola"
	usuario := "Programador"
	mensaje := saludo + " " + usuario
	mensajeFormateado := fmt.Sprintf("%s, %s!", saludo, usuario)

	fmt.Println("mensaje", mensaje)
	fmt.Println("mensaje", mensajeFormateado)
}
