package main

import "fmt"

func main() {
	num := 0
	if num > 0 {
		fmt.Println("El número es positivo", num)
	} else if num < 0 {
		fmt.Println("El número es negativo", num)
	} else if num == 0 {
		fmt.Println("El numero es cero", num)
	}
}
