package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i >= 1 && i <= 4 {
			fmt.Println("Intento fallido número", i)
		} else {
			fmt.Println("Cuenta bloqueada por seguridad")
		}
	}
}
