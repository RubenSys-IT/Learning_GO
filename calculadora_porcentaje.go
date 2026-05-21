package main

import "fmt"

func main() {
	precioOriginal := 1000

	if precioOriginal > 100 {
		descuento := float64(precioOriginal) * 0.9
		fmt.Println("El precio se queda en: ", descuento)
	} else {
		fmt.Println("El precio se queda igual", precioOriginal)
	}
}
