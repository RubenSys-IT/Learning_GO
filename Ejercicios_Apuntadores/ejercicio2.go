package main

import "fmt"

func main() {
	intentosFallidos := 4
	ptrIntentos := &intentosFallidos
	*ptrIntentos = 0
	fmt.Println("El valor es de: ", intentosFallidos)
}
