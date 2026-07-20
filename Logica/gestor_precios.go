package main

import "fmt"

func aplicardecuento(precios []int) {
	for i, v := range precios {
		precios[i] = v - 10
	}
}

func main() {
	precios := []int{10, 25, 60, 80, 15}
	respaldo := make([]int, len(precios))

	// copy nos devuelve la cantidad de elementos copiados (5), pero la copia real está en 'respaldo'
	copy(respaldo, precios)

	caro := precios[2:4] // Vista con [60, 80]

	// ⚠️ TE FALTABA ESTO: ¡Llamar a la función!
	aplicardecuento(caro)

	fmt.Println("Original (cambiado por la vista):", precios)  // [10 25 50 70 15]
	fmt.Println("Vista 'caro':                    ", caro)     // [50 70]
	fmt.Println("Respaldo (copia intacta):        ", respaldo) // [10 25 60 80 15]
}
