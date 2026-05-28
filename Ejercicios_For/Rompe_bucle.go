package main

import "fmt"

func main() {
	logsProcesados := 1

	for {
		fmt.Println("Analizando log número ")
		if logsProcesados == 3 {
			fmt.Println("¡Amenaza detectada! Deteniendo análisis.")
			break
		}
		logsProcesados++
	}
}
