package main

import "fmt"

func main() {
	puertoObjetivo := 80
	ptrPuerto := &puertoObjetivo
	fmt.Println("EL puerto es: ", puertoObjetivo)
	fmt.Println("El valor de memoria es: ", *ptrPuerto)

}
