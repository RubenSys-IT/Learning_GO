package main

import "fmt"

// 1. Modificado: Ahora recibe un entero común y un PUNTERO (*int)
func calcularAmenaza(tpaq int, acumuladorAmenaza *int) {
	switch {
	case tpaq > 800:
		*acumuladorAmenaza += 10 // Modifica directamente la variable de main
	case tpaq >= 500 && tpaq <= 800: // Corregido el >= por si es exactamente 500
		*acumuladorAmenaza += 3
	}
}

func main() {
	amenazas2 := 0 // Esta es la variable que queremos que la función modifique
	historialPaquetes := []int{}
	var valores int

	// --- BUCLE 1: CAPTURA DE DATOS (Tu código impecable) ---
	for {
		fmt.Print("Introduce el tamaño de los paquetes (0 para salir): ")
		fmt.Scanln(&valores)

		if valores == 0 {
			break
		}

		if valores >= 500 {
			historialPaquetes = append(historialPaquetes, valores)
			fmt.Println("¡Paquete sospechoso registrado!")
		} else {
			fmt.Println("No es necesario añadir")
		}
	}

	// --- BUCLE 2: RECORRER EL SLICE (Esto faltaba) ---
	// Usamos un bucle clásico para pasar cada paquete guardado a la función
	for i := 0; i < len(historialPaquetes); i++ {
		// Pasamos el valor actual del slice y la DIRECCIÓN (&) de amenazas2
		calcularAmenaza(historialPaquetes[i], &amenazas2)
	}

	// --- REPORTE FINAL ---
	fmt.Printf("\n--- REPORTE FINAL ---\n")
	fmt.Printf("El contenido de historial paquetes: %v\n", historialPaquetes)
	fmt.Printf("Cantidad de paquetes sospechosos: %d\n", len(historialPaquetes))
	fmt.Printf("Amenaza total acumulada: %d\n", amenazas2) // Añadido para ver el total

	if amenazas2 >= 20 {
		fmt.Println("¡ESTADO: ALERTA ROJA EN LA RED!")
	} else {
		fmt.Println("¡ESTADO: RED ESTABLE!")
	}
}
