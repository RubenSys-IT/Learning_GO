package main

import "fmt"

func main() {
	tokenOriginal := "XYZ123"
	ptr1 := &tokenOriginal
	ptr2 := &tokenOriginal

	fmt.Println("La dirección 1: ", ptr1)
	fmt.Println("La dirección 2: ", ptr2)

	//Modificacion del token

	*ptr1 = "ABC789"
	fmt.Println("La dirección modificada 1: ", *ptr2)
}
