package main

import "fmt"

func main() {
	logsSospechosos := []int{}
	var valores int

	for {
		fmt.Print("Introduce el código de error (0 para salir): ")
		fmt.Scanln(&valores)

		// Condición de salida directa
		if valores == 0 {
			break
		}

		// Lógica de filtrado
		if valores >= 400 {
			logsSospechosos = append(logsSospechosos, valores)
			fmt.Println("¡Error crítico registrado!")
		}
	}

	// Reporte final
	fmt.Printf("\n--- REPORTE FINAL ---\n")
	fmt.Printf("Logs detectados: %v\n", logsSospechosos)
	fmt.Printf("Total de elementos: %d\n", len(logsSospechosos))
}
