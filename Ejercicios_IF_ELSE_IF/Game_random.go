package main

import (
	"fmt"
	rand "math/rand/v2"
)

func main() {
	matriz := rand.IntN(3) + 1
	if matriz == 1 {
		fmt.Println("Has encontrado un cofre con oro")
	} else if matriz == 2 {
		fmt.Println("Un monstruo ha aparecido")
	} else if matriz == 3 {
		fmt.Println("El camino está tranquilo")
	}
}
