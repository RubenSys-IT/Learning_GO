//Ejercicio apuntadores

package main

import "fmt"

func limpiarinput(valor *string) {
	*valor = "Hola"
}

func main() {
	//La real
	variable := "Hector"
	fmt.Println("Valor de la variable: ", variable)
	//La modificada

	limpiarinput(&variable)
	fmt.Println("Valor de la variable: ", variable)

}
