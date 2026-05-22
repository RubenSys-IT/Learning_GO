package main

import "fmt"

func main() {
	presion := 50

	if presion > 90 {
		fmt.Printf("Alerta, Presión demasiado alta")
	} else if presion < 50 {
		fmt.Printf("Alerta, Presión demasiado baja")
	} else {
		fmt.Printf("Presion normal")
	}

}
