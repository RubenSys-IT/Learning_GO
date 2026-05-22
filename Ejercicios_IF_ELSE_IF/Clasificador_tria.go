package main

import "fmt"

func main() {
	ladoA := 80
	ladoB := 80
	ladoC := 80

	if ladoA == ladoA && ladoB == ladoC {
		fmt.Printf("Equilatero")
	} else if ladoA == ladoB {
		fmt.Printf("Isosceles")
	} else {
		fmt.Printf("Escaleno")
	}
}
